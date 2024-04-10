package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// getLatestProviderVersion retrieves the latest stable provider version from GitHub releases
func GetLatestProviderVersion(provider string, source string) (string, error) {
	splitSource := strings.Split(source, `/`)
	providerReleasesURL := fmt.Sprintf("https://api.github.com/repos/%s/terraform-provider-%s/releases/latest", splitSource[0], provider)
	resp, err := http.Get(providerReleasesURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to retrieve latest %s provider version. Status code: %d", provider, resp.StatusCode)
	}

	var releaseInfo map[string]interface{}
	if err := ReadJSON(resp.Body, &releaseInfo); err != nil {
		return "", err
	}

	tagName, ok := releaseInfo["tag_name"].(string)
	if !ok {
		return "", fmt.Errorf("Failed to extract latest %s provider version from the response", provider)
	}

	return strings.TrimPrefix(tagName, "v"), nil
}

// getLatestTerraformVersion retrieves the latest stable Terraform version from GitHub releases
func GetLatestTerraformVersion() (string, error) {
	const terraformReleasesURL = "https://api.github.com/repos/hashicorp/terraform/releases/latest"
	resp, err := http.Get(terraformReleasesURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to retrieve latest Terraform version. Status code: %d", resp.StatusCode)
	}

	var releaseInfo map[string]interface{}
	if err := ReadJSON(resp.Body, &releaseInfo); err != nil {
		return "", err
	}

	tagName, ok := releaseInfo["tag_name"].(string)
	if !ok {
		return "", fmt.Errorf("Failed to extract latest Terraform version from the response")
	}

	// Extract the version from the tag name (e.g., "v0.15.3" becomes "0.15.3")
	return strings.TrimPrefix(tagName, "v"), nil
}

// readJSON reads JSON from a reader and unmarshals it into the provided interface
func ReadJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
