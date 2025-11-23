package naming

import (
	"testing"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
)

func TestISODateRule(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name        string
		filename    string
		expectWarn  bool
		description string
	}{
		// Valid ISO 8601 formats - no warning
		{"ISO 8601 with dashes", "report-2025-11-23.txt", false, "YYYY-MM-DD is ISO standard"},
		{"ISO 8601 compact", "backup_20251123.tar.gz", false, "YYYYMMDD is acceptable"},
		{"ISO 8601 in middle", "data-2025-11-23-final.csv", false, "ISO format in filename"},
		
		// US format - should warn
		{"US date format", "report-11-23-2025.txt", true, "MM-DD-YYYY should suggest ISO"},
		{"US date with underscores", "backup_11_23_2025.txt", true, "MM_DD_YYYY should suggest ISO"},
		{"US date compact", "log11232025.txt", true, "MMDDYYYY should suggest ISO"},
		
		// European format - should warn
		{"EU date format", "report-23-11-2025.txt", true, "DD-MM-YYYY should suggest ISO"},
		{"EU date with dots", "backup.23.11.2025.txt", true, "DD.MM.YYYY should suggest ISO"},
		
		// Month names - should warn
		{"Month name full", "report-November-23-2025.txt", true, "Month name should suggest ISO"},
		{"Month name short", "backup-Nov-23-2025.txt", true, "Short month should suggest ISO"},
		{"Reverse month name", "data-23-Nov-2025.csv", true, "Day-Month-Year should suggest ISO"},
		
		// Ambiguous but likely dates - should warn
		{"Possible date", "file-12-25-2025.txt", true, "Could be MM-DD-YYYY"},
		
		// Not dates - should NOT warn (conservative)
		{"Version numbers", "v1.2.3.txt", false, "Version numbers aren't dates"},
		{"Port numbers", "server-8080-config.txt", false, "Port numbers aren't dates"},
		{"Just numbers", "file-123456.txt", false, "Random numbers"},
		{"High numbers", "data-99-99-9999.txt", false, "Invalid date values"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violations := validator.Validate(tt.filename, fsobj.TypeFile)
			hasWarnings := HasWarnings(violations)

			if hasWarnings != tt.expectWarn {
				t.Errorf("%s: expected warning=%v, got=%v for '%s'. Violations: %+v", 
					tt.description, tt.expectWarn, hasWarnings, tt.filename, violations)
			}
		})
	}
}

func TestDatePatternDetection(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{"MM-DD-YYYY", "11-23-2025", true},
		{"MM_DD_YYYY", "11_23_2025", true},
		{"MMDDYYYY", "11232025", true},
		{"Nov-23-2025", "Nov-23-2025", true},
		{"23-Nov-2025", "23-Nov-2025", true},
		{"YYYYMMDD compact", "20251123", true},
		{"Version 1.2.3", "1.2.3", false},
		{"Random numbers", "12345", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := detectNonISODate(tt.filename)
			hasDate := result != ""
			if hasDate != tt.want {
				t.Errorf("detectNonISODate(%q) = %q, want match=%v", tt.filename, result, tt.want)
			}
		})
	}
}
