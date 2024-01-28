package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"tfversion-checker/utils"
)

const terraformReleasesURL = "https://api.github.com/repos/hashicorp/terraform/releases/latest"

func main() {
	// Define CLI flags

	scanFlag := flag.Bool("scan", false, "Initiate a scan of the current directory for Terraform configurations")
	checkFlag := flag.Bool("check", false, "Perform version checks and display results")
	pathFlag := flag.String("path", ".", "Specify the path to scan for Terraform configurations")
	enforceFlag := flag.Bool("enforce", false, "Enforce version policies across projects")

	// Parse command-line flags
	flag.Parse()

	// Validate flag combinations
	flagCount := countFlags(*scanFlag, *checkFlag, *enforceFlag)
	if flagCount != 1 {
		fmt.Println("Exactly one of -scan, -check, or -enforce must be specified.")
		flag.Usage()
		os.Exit(1)
	}

	switch {
	case *scanFlag:
		scanAction(*pathFlag)
	case *checkFlag:
		checkAction(*pathFlag)
	case *enforceFlag:
		enforceAction(*pathFlag)
	}
}

// countFlags counts the number of specified flags
func countFlags(flags ...bool) int {
	count := 0
	for _, flag := range flags {
		if flag {
			count++
		}
	}
	return count
}

// scanAction initiates a scan of the current directory for Terraform configurations
func scanAction(path string) {
	fmt.Println("Scanning for Terraform configurations in the current directory...")
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".tf") {
			checkTerraformVersion(path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error scanning for Terraform configurations: %v\n", err)
	}
}

// checkTerraformVersion reads the content of a Terraform configuration file and checks for the required_version attribute
func checkTerraformVersion(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if match := regexp.MustCompile(`\brequired_version\s*=\s*"(.*?)"`).FindStringSubmatch(line); len(match) > 1 {
			version := match[1]
			fmt.Printf("Terraform version specified in %s: %s\n", filePath, version)
			// Compare with the latest stable release
			latestVersion, err := utils.GetLatestProviderVersion(terraformReleasesURL)
			if err != nil {
				fmt.Printf("Error retrieving latest Terraform version: %v\n", err)
				return
			}

			if version != latestVersion {
				fmt.Printf("Warning: The specified version is outdated. Latest version: %s\n", latestVersion)
			}
		}
		if match := regexp.MustCompile(`\brequired_providers\s*{([^}]*)}`).FindStringSubmatch(line); len(match) > 1 {
			requiredProvidersBlock := match[1]

			// Parse required_providers block and check versions
			parseRequiredProvidersBlock(requiredProvidersBlock)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
	}
}

// parseRequiredProvidersBlock parses the required_providers block and checks versions
func parseRequiredProvidersBlock(block string) {
	scanner := bufio.NewScanner(strings.NewReader(block))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if match := regexp.MustCompile(`\b"([^"]+)"\s*=\s*"(.*?)"`).FindStringSubmatch(line); len(match) > 2 {
			provider := match[1]
			version := match[2]

			// Compare provider version with the latest stable release
			latestProviderVersion, err := utils.GetLatestProviderVersion(provider)
			if err != nil {
				fmt.Printf("Error retrieving latest %s provider version: %v\n", provider, err)
				continue
			}

			if version != latestProviderVersion {
				fmt.Printf("Warning: The specified %s provider version is outdated. Latest version: %s\n", provider, latestProviderVersion)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error parsing required_providers block: %v\n", err)
	}
}

// getLatestProviderVersion retrieves the latest stable version of a Terraform provider from GitHub releases

// checkAction performs version checks and displays results
func checkAction(path string) {
	fmt.Println("Performing version checks and displaying results...")
	// Implement logic for checking Terraform versions
}

// enforceAction enforces version policies across projects
func enforceAction(path string) {
	fmt.Println("Enforcing version policies across projects...")
	// Implement logic for enforcing version policies
}
