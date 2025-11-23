package naming

import (
	"testing"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
)

func TestLogFileNamingRule(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name        string
		filename    string
		expectWarn  bool
		description string
	}{
		// Valid log file names
		{"ISO date prefix", "2025-11-23-app.log", false, "Log file with ISO date prefix"},
		{"ISO with details", "2025-11-23-server-error.log", false, "Log with ISO date and details"},
		{"ISO compact in middle", "app-2025-11-23.log", true, "ISO date not at start"},
		
		// Invalid log file names
		{"No date prefix", "application.log", true, "Log file without date prefix"},
		{"US date format", "11-23-2025-app.log", true, "Log with US date format"},
		{"Month name", "Nov-23-2025.log", true, "Log with month name"},
		
		// Non-log files (should not be checked)
		{"Text file no date", "application.txt", false, "Not a log file"},
		{"CSV no date", "data.csv", false, "Not a log file"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violations := validator.Validate(tt.filename, fsobj.TypeFile)
			
			// Check specifically for log file rule violations
			hasLogWarning := false
			for _, v := range violations {
				if v.Rule == "log-files-iso-date-prefix" && v.IsWarning() {
					hasLogWarning = true
					break
				}
			}

			if hasLogWarning != tt.expectWarn {
				t.Errorf("%s: expected log warning=%v, got=%v for '%s'. Violations: %+v", 
					tt.description, tt.expectWarn, hasLogWarning, tt.filename, violations)
			}
		})
	}
}

func TestReadmeNamingRule(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name        string
		filename    string
		expectWarn  bool
		description string
	}{
		// Valid (uppercase)
		{"Uppercase README", "README", false, "Properly uppercase"},
		{"Uppercase with ext", "README.md", false, "Uppercase with extension"},
		{"Uppercase txt", "README.txt", false, "Uppercase with .txt"},
		
		// Invalid (lowercase or mixed)
		{"Lowercase", "readme", true, "Should be uppercase"},
		{"Lowercase with ext", "readme.md", true, "Should be README.md"},
		{"Mixed case", "Readme.md", true, "Should be all caps"},
		{"Title case", "ReadMe.md", true, "Should be all caps"},
		
		// Non-readme files
		{"Other file", "myfile.md", false, "Not a readme file"},
		{"Contains readme", "my-readme.txt", false, "Contains but doesn't start with readme"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violations := validator.Validate(tt.filename, fsobj.TypeFile)
			
			// Check specifically for readme rule violations
			hasReadmeWarning := false
			for _, v := range violations {
				if v.Rule == "readme-uppercase" && v.IsWarning() {
					hasReadmeWarning = true
					break
				}
			}

			if hasReadmeWarning != tt.expectWarn {
				t.Errorf("%s: expected readme warning=%v, got=%v for '%s'. Violations: %+v", 
					tt.description, tt.expectWarn, hasReadmeWarning, tt.filename, violations)
			}
		})
	}
}

func TestCommonFilesUppercaseRule(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name        string
		filename    string
		expectWarn  bool
		description string
	}{
		// Valid (uppercase)
		{"LICENSE uppercase", "LICENSE", false, "Properly uppercase"},
		{"LICENSE with ext", "LICENSE.txt", false, "Uppercase with extension"},
		{"CHANGELOG", "CHANGELOG.md", false, "Changelog uppercase"},
		{"CONTRIBUTING", "CONTRIBUTING.md", false, "Contributing uppercase"},
		
		// Invalid (lowercase)
		{"license lowercase", "license", true, "Should be uppercase"},
		{"changelog lowercase", "changelog.md", true, "Should be CHANGELOG.md"},
		{"contributing lowercase", "contributing.md", true, "Should be CONTRIBUTING.md"},
		{"makefile lowercase", "makefile", true, "Should be MAKEFILE"},
		
		// Non-common files
		{"Regular file", "myfile.txt", false, "Not a common file"},
		{"Code file", "main.go", false, "Not a common file"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violations := validator.Validate(tt.filename, fsobj.TypeFile)
			
			// Check specifically for common files rule violations
			hasCommonWarning := false
			for _, v := range violations {
				if v.Rule == "common-files-uppercase" && v.IsWarning() {
					hasCommonWarning = true
					break
				}
			}

			if hasCommonWarning != tt.expectWarn {
				t.Errorf("%s: expected common file warning=%v, got=%v for '%s'. Violations: %+v", 
					tt.description, tt.expectWarn, hasCommonWarning, tt.filename, violations)
			}
		})
	}
}

func TestArchivePreferenceRule(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name        string
		filename    string
		expectWarn  bool
		description string
	}{
		// Should suggest 7z
		{"ZIP archive", "backup.zip", true, "Should suggest 7z"},
		{"ZIP uppercase", "BACKUP.ZIP", true, "Should suggest 7z"},
		
		// Should not warn
		{"7z archive", "backup.7z", false, "Already using 7z"},
		{"tar.gz", "backup.tar.gz", false, "tar.gz is acceptable"},
		{"tar", "backup.tar", false, "tar is acceptable"},
		{"Not an archive", "backup.txt", false, "Not an archive"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violations := validator.Validate(tt.filename, fsobj.TypeArchive)
			
			// Check specifically for archive preference rule violations
			hasArchiveWarning := false
			for _, v := range violations {
				if v.Rule == "prefer-7z-archives" && v.IsWarning() {
					hasArchiveWarning = true
					break
				}
			}

			if hasArchiveWarning != tt.expectWarn {
				t.Errorf("%s: expected archive warning=%v, got=%v for '%s'. Violations: %+v", 
					tt.description, tt.expectWarn, hasArchiveWarning, tt.filename, violations)
			}
		})
	}
}
