package naming

import (
	"testing"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
)

func TestFileNamingRules(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name        string
		filename    string
		expectError bool
		expectWarn  bool
	}{
		{"valid file", "test.txt", false, false},
		{"file with spaces - error", "this is stupid", true, false},
		{"file without extension - error", "README", true, false},
		{"dotfile - valid", ".gitignore", false, false},
		{"valid file with multiple dots", "test.backup.txt", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violations := validator.Validate(tt.filename, fsobj.TypeFile)
			hasErrors := HasErrors(violations)
			hasWarnings := HasWarnings(violations)

			if hasErrors != tt.expectError {
				t.Errorf("expected error=%v, got=%v for '%s'. Violations: %+v", tt.expectError, hasErrors, tt.filename, violations)
			}
			if hasWarnings != tt.expectWarn {
				t.Errorf("expected warning=%v, got=%v for '%s'. Violations: %+v", tt.expectWarn, hasWarnings, tt.filename, violations)
			}
		})
	}
}

func TestDirectoryNamingRules(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name        string
		dirname     string
		expectError bool
		expectWarn  bool
	}{
		{"lowercase dir - valid", "myproject", false, false},
		{"uppercase dir - warning", "TestMe", false, true},
		{"python keyword - error", "class", true, false},
		{"python keyword False - error", "False", true, true}, // Has both error and warning
		{"mixed case - warning", "MyProject", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violations := validator.Validate(tt.dirname, fsobj.TypeDirectory)
			hasErrors := HasErrors(violations)
			hasWarnings := HasWarnings(violations)

			if hasErrors != tt.expectError {
				t.Errorf("expected error=%v, got=%v for '%s'. Violations: %+v", tt.expectError, hasErrors, tt.dirname, violations)
			}
			if hasWarnings != tt.expectWarn {
				t.Errorf("expected warning=%v, got=%v for '%s'. Violations: %+v", tt.expectWarn, hasWarnings, tt.dirname, violations)
			}
		})
	}
}

func TestArchiveNamingRules(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name        string
		archivename string
		expectError bool
		expectWarn  bool
	}{
		{"valid zip", "backup.zip", false, true}, // Now warns about .zip vs .7z
		{"valid tar.gz", "backup.tar.gz", false, false},
		{"archive with spaces - error", "my backup.zip", true, true}, // Error + warning
		{"archive without extension - error", "backup", true, false},
		{"invalid extension - error", "backup.txt", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violations := validator.Validate(tt.archivename, fsobj.TypeArchive)
			hasErrors := HasErrors(violations)
			hasWarnings := HasWarnings(violations)

			if hasErrors != tt.expectError {
				t.Errorf("expected error=%v, got=%v for '%s'. Violations: %+v", tt.expectError, hasErrors, tt.archivename, violations)
			}
			if hasWarnings != tt.expectWarn {
				t.Errorf("expected warning=%v, got=%v for '%s'. Violations: %+v", tt.expectWarn, hasWarnings, tt.archivename, violations)
			}
		})
	}
}
