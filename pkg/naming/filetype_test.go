package naming

import (
	"testing"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
)

func TestLogFileRule(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name        string
		filename    string
		expectWarn  bool
		description string
	}{
		{"log with ISO prefix", "2025-11-23-application.log", false, "Correct ISO prefix"},
		{"log with ISO prefix and time", "2025-11-23-app.log", false, "ISO date prefix is valid"},
		{"log without date", "application.log", true, "Should warn - no date prefix"},
		{"log with wrong format", "11-23-2025-app.log", true, "Should warn - not ISO format"},
		{"log with compact date", "20251123-app.log", true, "Should warn - no separators"},
		{"non-log file", "application.txt", false, "Rule doesn't apply to non-log files"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violations := validator.Validate(tt.filename, fsobj.TypeFile)
			
			// Filter for log-specific violations
			hasLogWarning := false
			for _, v := range violations {
				if v.Rule == "log-files-iso-date-prefix" {
					hasLogWarning = true
					break
				}
			}

			if hasLogWarning != tt.expectWarn {
				t.Errorf("%s: expected warning=%v, got=%v. Violations: %+v", 
					tt.description, tt.expectWarn, hasLogWarning, violations)
			}
		})
	}
}

func TestReadmeRule(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name        string
		filename    string
		expectWarn  bool
	}{
		{"README.md uppercase", "README.md", false},
		{"README without ext", "README", false},
		{"readme.md lowercase", "readme.md", true},
		{"Readme.md mixed case", "Readme.md", true},
		{"readme.txt lowercase", "readme.txt", true},
		{"not a readme", "myfile.md", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violations := validator.Validate(tt.filename, fsobj.TypeFile)
			
			hasReadmeWarning := false
			for _, v := range violations {
				if v.Rule == "readme-uppercase" {
					hasReadmeWarning = true
					break
				}
			}

			if hasReadmeWarning != tt.expectWarn {
				t.Errorf("expected warning=%v, got=%v for '%s'. Violations: %+v", 
					tt.expectWarn, hasReadmeWarning, tt.filename, violations)
			}
		})
	}
}

func TestCommonFilesRule(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name        string
		filename    string
		expectWarn  bool
	}{
		{"LICENSE uppercase", "LICENSE", false},
		{"LICENSE.md uppercase", "LICENSE.md", false},
		{"license lowercase", "license", true},
		{"License.txt mixed", "License.txt", true},
		{"CHANGELOG uppercase", "CHANGELOG", false},
		{"changelog.md lowercase", "changelog.md", true},
		{"CONTRIBUTING uppercase", "CONTRIBUTING.md", false},
		{"contributing lowercase", "contributing.md", true},
		{"regular file", "myfile.txt", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violations := validator.Validate(tt.filename, fsobj.TypeFile)
			
			hasCommonFileWarning := false
			for _, v := range violations {
				if v.Rule == "common-files-uppercase" {
					hasCommonFileWarning = true
					break
				}
			}

			if hasCommonFileWarning != tt.expectWarn {
				t.Errorf("expected warning=%v, got=%v for '%s'. Violations: %+v", 
					tt.expectWarn, hasCommonFileWarning, tt.filename, violations)
			}
		})
	}
}

func TestArchivePreferenceRuleFiletype(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name        string
		filename    string
		expectWarn  bool
	}{
		{"zip archive", "backup.zip", true},
		{"ZIP uppercase", "BACKUP.ZIP", true},
		{"7z archive", "backup.7z", false},
		{"tar.gz archive", "backup.tar.gz", false},
		{"regular file", "file.txt", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violations := validator.Validate(tt.filename, fsobj.TypeArchive)
			
			hasArchiveWarning := false
			for _, v := range violations {
				if v.Rule == "prefer-7z-archives" {
					hasArchiveWarning = true
					break
				}
			}

			if hasArchiveWarning != tt.expectWarn {
				t.Errorf("expected warning=%v, got=%v for '%s'. Violations: %+v", 
					tt.expectWarn, hasArchiveWarning, tt.filename, violations)
			}
		})
	}
}
