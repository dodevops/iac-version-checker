package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"tfversion-checker/utils"
)


var terraformFilePath []string
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
	println(flagCount)
	if flagCount == 0 {
		fmt.Println("One of the flags must be specified.")
		flag.Usage()
		os.Exit(1)
	}
// Check if file path is provided
	if *pathFlag == "" {
		fmt.Println("Error: Please provide the path to the Terraform configuration file using the -file flag.")
		return
	}

	if *scanFlag {
		scanAction(*pathFlag)
	}

	if *checkFlag {
		checkAction()
	}

	if *enforceFlag {
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
	fileCount := 0
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".tf") {
			//store the path of the file in a list
			
			// display number of files found 
			fileCount++
			terraformFilePath = append(terraformFilePath, path)
			
			
		}
		return nil
	})
	fmt.Printf("Found Terraform configuration file: %d\n", fileCount)
	if err != nil {
		fmt.Printf("Error scanning for Terraform configurations: %v\n", err)
	}
}





// checkAction performs version checks and displays results
func checkAction() {
	fmt.Println("Performing version checks and displaying results...")
	for _, filePath := range terraformFilePath {
		checkTerraformVersion(filePath)
	}
}

// enforceAction enforces version policies across projects
func enforceAction(path string) {
	fmt.Println("Enforcing version policies across projects...")
	// Implement logic for enforcing version policies
}


// checkTerraformVersion reads the content of a Terraform configuration file and checks for the required_version attribute
func checkTerraformVersion(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	

	   str := string(file)

		if match := regexp.MustCompile(`\brequired_version\s*=\s*"(.*?)"`).FindStringSubmatch(str); len(match) > 1 {
			version := match[1]
			fmt.Printf("Terraform version specified in %s: %s\n", filePath, version)
			// Compare with the latest stable release
			latestVersion, err := utils.GetLatestTerraformVersion()
			if err != nil {
				fmt.Printf("Error retrieving latest Terraform version: %v\n", err)
				return
			}

			if version != latestVersion {
				fmt.Printf("Warning: The specified version is outdated. Latest version: %s\n", latestVersion)
			}
		}
		// different option: required_providers\s*{(?:[^{}]+|{(?:[^{}]+|{[^{}]*})*})*}
		//pattern := `required_providers\s*{(?s:(.*?))}`
		pattern := `required_providers\s*{(?:[^{}]+|{(?:[^{}]+|{[^{}]*})*})*}`
		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(str,-1)
		fmt.Println("Match is ",matches)
		fmt.Println("Length of matches is ",len(matches))
		for _, match := range matches {
			fmt.Println("Match is ",match)
			parseRequiredProvidersBlock(match[0])
		}
			
		
		
	
}


// parseRequiredProvidersBlock parses the required_providers block and checks versions
func parseRequiredProvidersBlock(block string) {
		fmt.Println("Required Providers Block: ",block)
		// Extract provider and version from the block
		providerPattern := `\s*([a-zA-Z0-9_]+)\s*=\s*{[^{}]*\bsource\s*=\s*"([^"]+)"[^{}]*\bversion\s*=\s*"([^"]+)"}`
		providerMatches := regexp.MustCompile(providerPattern).FindAllStringSubmatch(block, -1)
		
		for _, provider := range providerMatches {
			if len(provider) > 0 {
				fmt.Printf("Required Provider: %s\n", provider[1])
				fmt.Printf("Source: %s\n", provider[2])
				fmt.Printf("Version: %s\n", provider[3])
				fmt.Println(strings.Repeat("-", 20))
			}
		}

			// Compare provider version with the latest stable release
			// latestProviderVersion, err := utils.GetLatestProviderVersion(provider)
			// if err != nil {
			// 	fmt.Printf("Error retrieving latest %s provider version: %v\n", provider, err)
			// 	return
			// }

			// if version != latestProviderVersion {
			// 	fmt.Printf("Warning: The specified %s provider version is outdated. Latest version: %s\n", provider, latestProviderVersion)
			// }
		}


