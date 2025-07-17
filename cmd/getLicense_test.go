package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-github/v28/github"
)

func TestFetchLicenseInfo(t *testing.T) {
	tests := []struct {
		name        string
		repo        string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "Valid repository format",
			repo:        "facebook/react",
			expectError: false,
		},
		{
			name:        "Invalid repository format - no slash",
			repo:        "invalid-repo",
			expectError: true,
			errorMsg:    "invalid repository format",
		},
		{
			name:        "Invalid repository format - empty string",
			repo:        "",
			expectError: true,
			errorMsg:    "invalid repository format",
		},
		{
			name:        "Invalid repository format - only slash",
			repo:        "/",
			expectError: true,
			errorMsg:    "invalid repository format",
		},
		{
			name:        "Non-existent repository",
			repo:        "nonexistent/repository-that-does-not-exist-12345",
			expectError: true,
			errorMsg:    "not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := fetchLicenseInfo(tt.repo)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error for repo '%s', but got nil", tt.repo)
				}
				if tt.errorMsg != "" && !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error message to contain '%s', but got: %s", tt.errorMsg, err.Error())
				}
				if repo != nil {
					t.Errorf("Expected nil repository for error case, but got: %v", repo)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error for repo '%s', but got: %s", tt.repo, err.Error())
				}
				if repo == nil {
					t.Errorf("Expected repository object for repo '%s', but got nil", tt.repo)
				}
				if repo != nil {
					// Validate that we got expected fields
					if repo.GetFullName() == "" {
						t.Errorf("Expected repository to have a full name")
					}
					if repo.GetHTMLURL() == "" {
						t.Errorf("Expected repository to have an HTML URL")
					}
				}
			}
		})
	}
}

func TestDisplayLicenseInfo(t *testing.T) {
	tests := []struct {
		name     string
		repo     *github.Repository
		expected []string
	}{
		{
			name: "Repository with MIT license",
			repo: &github.Repository{
				FullName:        github.String("test/repo"),
				Description:     github.String("Test repository"),
				HTMLURL:         github.String("https://github.com/test/repo"),
				StargazersCount: github.Int(100),
				ForksCount:      github.Int(50),
				Language:        github.String("Go"),
				License: &github.License{
					Name:   github.String("MIT License"),
					Key:    github.String("mit"),
					SPDXID: github.String("MIT"),
					URL:    github.String("https://api.github.com/licenses/mit"),
				},
			},
			expected: []string{
				"Repository: test/repo",
				"Description: Test repository",
				"URL: https://github.com/test/repo",
				"Stars: 100",
				"Forks: 50",
				"Language: Go",
				"License: MIT License",
				"License Key: mit",
				"SPDX ID: MIT",
				"License URL: https://api.github.com/licenses/mit",
			},
		},
		{
			name: "Repository with no license",
			repo: &github.Repository{
				FullName:        github.String("test/no-license"),
				Description:     github.String("Repository without license"),
				HTMLURL:         github.String("https://github.com/test/no-license"),
				StargazersCount: github.Int(10),
				ForksCount:      github.Int(5),
				Language:        github.String("Python"),
				License:         nil,
			},
			expected: []string{
				"Repository: test/no-license",
				"Description: Repository without license",
				"URL: https://github.com/test/no-license",
				"Stars: 10",
				"Forks: 5",
				"Language: Python",
				"License: No license information available",
				"Note: This repository may not have a license file or the license is not recognized by GitHub.",
			},
		},
		{
			name: "Repository with empty description",
			repo: &github.Repository{
				FullName:        github.String("test/empty-desc"),
				Description:     github.String(""),
				HTMLURL:         github.String("https://github.com/test/empty-desc"),
				StargazersCount: github.Int(0),
				ForksCount:      github.Int(0),
				Language:        github.String("JavaScript"),
				License: &github.License{
					Name:   github.String("Apache License 2.0"),
					Key:    github.String("apache-2.0"),
					SPDXID: github.String("Apache-2.0"),
					URL:    github.String("https://api.github.com/licenses/apache-2.0"),
				},
			},
			expected: []string{
				"Repository: test/empty-desc",
				"Description: ",
				"URL: https://github.com/test/empty-desc",
				"Stars: 0",
				"Forks: 0",
				"Language: JavaScript",
				"License: Apache License 2.0",
				"License Key: apache-2.0",
				"SPDX ID: Apache-2.0",
				"License URL: https://api.github.com/licenses/apache-2.0",
			},
		},
		{
			name: "Repository with license but no URL",
			repo: &github.Repository{
				FullName:        github.String("test/license-no-url"),
				Description:     github.String("Repository with license but no URL"),
				HTMLURL:         github.String("https://github.com/test/license-no-url"),
				StargazersCount: github.Int(25),
				ForksCount:      github.Int(3),
				Language:        github.String("Rust"),
				License: &github.License{
					Name:   github.String("Custom License"),
					Key:    github.String("custom"),
					SPDXID: github.String("NOASSERTION"),
					URL:    github.String(""),
				},
			},
			expected: []string{
				"Repository: test/license-no-url",
				"Description: Repository with license but no URL",
				"URL: https://github.com/test/license-no-url",
				"Stars: 25",
				"Forks: 3",
				"Language: Rust",
				"License: Custom License",
				"License Key: custom",
				"SPDX ID: NOASSERTION",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call the function
			displayLicenseInfo(tt.repo)

			// Restore stdout
			w.Close()
			os.Stdout = old

			// Read the output
			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			// Check that all expected strings are present
			for _, expected := range tt.expected {
				if !strings.Contains(output, expected) {
					t.Errorf("Expected output to contain '%s', but it didn't.\nActual output:\n%s", expected, output)
				}
			}

			// Check that the separator line is present
			if !strings.Contains(output, "─────────────────────────────────────────") {
				t.Errorf("Expected output to contain separator line, but it didn't.\nActual output:\n%s", output)
			}
		})
	}
}

func TestDisplayLicenseInfoOutput(t *testing.T) {
	// Test to ensure output format is consistent
	repo := &github.Repository{
		FullName:        github.String("owner/repo"),
		Description:     github.String("Test description"),
		HTMLURL:         github.String("https://github.com/owner/repo"),
		StargazersCount: github.Int(42),
		ForksCount:      github.Int(7),
		Language:        github.String("Go"),
		License: &github.License{
			Name:   github.String("MIT License"),
			Key:    github.String("mit"),
			SPDXID: github.String("MIT"),
			URL:    github.String("https://api.github.com/licenses/mit"),
		},
	}

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	displayLicenseInfo(repo)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Split output into lines for more detailed testing
	lines := strings.Split(strings.TrimSpace(output), "\n")

	expectedLines := []string{
		"Repository: owner/repo",
		"Description: Test description",
		"URL: https://github.com/owner/repo",
		"Stars: 42",
		"Forks: 7",
		"Language: Go",
		"─────────────────────────────────────────",
		"License: MIT License",
		"License Key: mit",
		"SPDX ID: MIT",
		"License URL: https://api.github.com/licenses/mit",
	}

	if len(lines) != len(expectedLines) {
		t.Errorf("Expected %d lines, but got %d lines.\nActual output:\n%s", len(expectedLines), len(lines), output)
	}

	for i, expectedLine := range expectedLines {
		if i < len(lines) && lines[i] != expectedLine {
			t.Errorf("Line %d: expected '%s', but got '%s'", i+1, expectedLine, lines[i])
		}
	}
}

// Benchmark tests
func BenchmarkFetchLicenseInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := fetchLicenseInfo("facebook/react")
		if err != nil {
			b.Errorf("Error in benchmark: %v", err)
		}
	}
}

func BenchmarkDisplayLicenseInfo(b *testing.B) {
	repo := &github.Repository{
		FullName:        github.String("test/repo"),
		Description:     github.String("Test repository"),
		HTMLURL:         github.String("https://github.com/test/repo"),
		StargazersCount: github.Int(100),
		ForksCount:      github.Int(50),
		Language:        github.String("Go"),
		License: &github.License{
			Name:   github.String("MIT License"),
			Key:    github.String("mit"),
			SPDXID: github.String("MIT"),
			URL:    github.String("https://api.github.com/licenses/mit"),
		},
	}

	// Redirect stdout to discard output during benchmarking
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() { os.Stdout = old }()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		displayLicenseInfo(repo)
	}
}
