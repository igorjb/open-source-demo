/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v28/github"
	"github.com/spf13/cobra"
)

var repoFlag string

// getLicenseCmd represents the getLicense command
var getLicenseCmd = &cobra.Command{
	Use:   "getLicense",
	Short: "Retrieve the open source license for a GitHub repository",
	Long: `This command retrieves the open source license information for a specified GitHub repository.
It queries the GitHub API to fetch license details and displays them in a user-friendly format.

Example usage:
  open-source-demo getLicense --repo amenocal/gh-pin-actions
  open-source-demo getLicense --repo facebook/react`,

	Run: func(cmd *cobra.Command, args []string) {
		if repoFlag == "" {
			fmt.Println("Error: --repo flag is required")
			fmt.Println("Usage: open-source-demo getLicense --repo owner/repository")
			os.Exit(1)
		}

		license, err := fetchLicenseInfo(repoFlag)
		if err != nil {
			fmt.Printf("Error retrieving license information: %v\n", err)
			os.Exit(1)
		}

		displayLicenseInfo(license)
	},
}

func init() {
	rootCmd.AddCommand(getLicenseCmd)

	// Add the --repo flag
	getLicenseCmd.Flags().StringVarP(&repoFlag, "repo", "r", "", "GitHub repository in format owner/repository (required)")
	getLicenseCmd.MarkFlagRequired("repo")
}

// fetchLicenseInfo retrieves license information from GitHub API using go-github
func fetchLicenseInfo(repo string) (*github.Repository, error) {
	// Validate repo format
	if !strings.Contains(repo, "/") {
		return nil, fmt.Errorf("invalid repository format. Expected format: owner/repository")
	}

	// Split repo into owner and name
	parts := strings.SplitN(repo, "/", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid repository format. Expected format: owner/repository")
	}
	owner := strings.TrimSpace(parts[0])
	repoName := strings.TrimSpace(parts[1])

	// Validate that both owner and repo name are not empty
	if owner == "" || repoName == "" {
		return nil, fmt.Errorf("invalid repository format. Expected format: owner/repository")
	}

	// Create GitHub client (using older vulnerable version)
	client := github.NewClient(nil)
	ctx := context.Background()

	// Get repository information
	repository, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		// Check if it's a 404 error
		if strings.Contains(err.Error(), "404") {
			return nil, fmt.Errorf("repository '%s' not found", repo)
		}
		return nil, fmt.Errorf("failed to fetch repository information: %w", err)
	}

	return repository, nil
}

// displayLicenseInfo formats and displays the license information
func displayLicenseInfo(repo *github.Repository) {
	fmt.Printf("Repository: %s\n", repo.GetFullName())
	fmt.Printf("Description: %s\n", repo.GetDescription())
	fmt.Printf("URL: %s\n", repo.GetHTMLURL())
	fmt.Printf("Stars: %d\n", repo.GetStargazersCount())
	fmt.Printf("Forks: %d\n", repo.GetForksCount())
	fmt.Printf("Language: %s\n", repo.GetLanguage())
	fmt.Println("─────────────────────────────────────────")

	license := repo.GetLicense()
	if license == nil {
		fmt.Println("License: No license information available")
		fmt.Println("Note: This repository may not have a license file or the license is not recognized by GitHub.")
	} else {
		fmt.Printf("License: %s\n", license.GetName())
		fmt.Printf("License Key: %s\n", license.GetKey())
		fmt.Printf("SPDX ID: %s\n", license.GetSPDXID())
		if license.GetURL() != "" {
			fmt.Printf("License URL: %s\n", license.GetURL())
		}
	}
}
