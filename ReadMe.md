# [WIP]Terraform Version Checker CLI

## Concept Document

### Introduction:
The Terraform Version Checker CLI is a command-line tool designed to assist developers and DevOps teams in managing and validating Terraform versions within a codebase. This lightweight application aims to streamline version checking processes for users who prefer a command-line interface.


### Key Features:

- **Version Detection:**
  - Scan the codebase to automatically identify Terraform configurations.
  - Extract and display Terraform version information specified in configuration files.

- **Dependency Analysis:**
  - Analyze Terraform configuration dependencies to ensure compatibility with the specified version.

- **Version Comparison:**
  - Compare the detected Terraform version with the latest stable release.
  - Provide warnings or notifications for outdated versions.

- **Cross-Project Consistency:**
  - Support checking and enforcing consistent Terraform versions across multiple projects.

- **Customization:**
  - Allow users to configure version checking rules through a simple CLI interface.

- **Reporting:**
  - Generate concise reports on Terraform version usage within a codebase.

### CLI Usage:

The Terraform Version Checker CLI will provide a straightforward command-line interface with options such as:
  - `terraform-version-checker scan`: Initiates a scan of the current directory for Terraform configurations.
  - `terraform-version-checker check`: Performs version checks and displays results.
  - `terraform-version-checker enforce`: Enforces version policies across projects.




---
---
# Implementation specifics
## Terraform Version Checks
For terraform code version checks we could use a regex pattern and directory walk to check for required versions. 
Checking for latest versions can be done by using the github api and creating the url from the provider name. 

Url example:
```
https://api.github.com/repos/${providerSource}/terraform-provider-${providerName}/releases/latest

```
Export the report or warnings to prometheus metrics or another datasource. 
## K8S Version checks
Different codebases can use k8s in different ways. 
You could use a cloud based solution(AKS,EKS,GKE) or a onprem solution(VMWare Tanzu, Rancher ...etc). 


## Usage Guide

### Installation:

Ensure you have [Go](https://golang.org/dl/) installed on your machine.

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/terraform-version-checker.git
2. Navigate to folder
   ```sh
   cd terraform-version-checker
3. Build the application
   ```sh
   go build -o terraform-version-checker

## Commands:

### Scan for Terraform Configurations
Initiate a scan of the current or a specified directory for Terraform configurations.

```sh
./terraform-version-checker -scan -path /path/to/yourdirectory
```
### Check for Terraform Configurations
Initiate a check of the current or a specified directory for Terraform configurations.
```sh
./terraform-version-checker -check -path /path/to/yourdirectory
```
### Enforce for Terraform Configurations
Enforce predefined policies and standards
```sh
./terraform-version-checker -enforce -path /path/to/yourdirectory
```
