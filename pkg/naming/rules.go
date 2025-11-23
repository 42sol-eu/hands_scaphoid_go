package naming

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
)

// Severity levels for naming rules
type Severity string

const (
	SeverityMust Severity = "MUST" // Errors - blocks operation
	SeverityWant Severity = "WANT" // Warnings - allows operation
)

// Violation represents a naming rule violation
type Violation struct {
	Rule     string
	Message  string
	Severity Severity
}

// IsError returns true if violation is an error (MUST rule)
func (v *Violation) IsError() bool {
	return v.Severity == SeverityMust
}

// IsWarning returns true if violation is a warning (WANT rule)
func (v *Violation) IsWarning() bool {
	return v.Severity == SeverityWant
}

// Rule represents a naming validation rule
type Rule struct {
	Name        string
	Description string
	Severity    Severity
	Applies     fsobj.ObjectType // Which object types this rule applies to
	Validate    func(name string) (bool, string)
}

// Validator manages and executes naming rules
type Validator struct {
	rules []Rule
}

// NewValidator creates a new validator with default rules
func NewValidator() *Validator {
	v := &Validator{
		rules: []Rule{},
	}
	v.AddDefaultRules()
	return v
}

// AddRule adds a custom rule to the validator
func (v *Validator) AddRule(rule Rule) {
	v.rules = append(v.rules, rule)
}

// Validate checks a name against all applicable rules
func (v *Validator) Validate(name string, objType fsobj.ObjectType) []Violation {
	violations := []Violation{}

	// Extract just the filename/dirname, not the full path
	basename := filepath.Base(name)

	for _, rule := range v.rules {
		// Check if rule applies to this object type
		if rule.Applies != "" && rule.Applies != objType {
			continue
		}

		// Run validation
		if valid, message := rule.Validate(basename); !valid {
			violations = append(violations, Violation{
				Rule:     rule.Name,
				Message:  message,
				Severity: rule.Severity,
			})
		}
	}

	return violations
}

// HasErrors checks if any violations are errors
func HasErrors(violations []Violation) bool {
	for _, v := range violations {
		if v.IsError() {
			return true
		}
	}
	return false
}

// HasWarnings checks if any violations are warnings
func HasWarnings(violations []Violation) bool {
	for _, v := range violations {
		if v.IsWarning() {
			return true
		}
	}
	return false
}

// AddDefaultRules adds the built-in naming rules
func (v *Validator) AddDefaultRules() {
	// File rules
	v.AddRule(Rule{
		Name:        "file-no-spaces",
		Description: "File names must not contain spaces",
		Severity:    SeverityMust,
		Applies:     fsobj.TypeFile,
		Validate: func(name string) (bool, string) {
			if strings.Contains(name, " ") {
				return false, fmt.Sprintf("file name '%s' contains spaces", name)
			}
			return true, ""
		},
	})

	v.AddRule(Rule{
		Name:        "file-must-have-extension",
		Description: "File names must have an extension or start with a dot",
		Severity:    SeverityMust,
		Applies:     fsobj.TypeFile,
		Validate: func(name string) (bool, string) {
			// Allow dotfiles (e.g., .gitignore)
			if strings.HasPrefix(name, ".") && len(name) > 1 {
				return true, ""
			}
			// Check for extension
			ext := filepath.Ext(name)
			if ext == "" {
				return false, fmt.Sprintf("file name '%s' must have an extension or start with a dot", name)
			}
			return true, ""
		},
	})

	// Directory rules
	v.AddRule(Rule{
		Name:        "directory-lowercase",
		Description: "Directory names should be lowercase",
		Severity:    SeverityWant,
		Applies:     fsobj.TypeDirectory,
		Validate: func(name string) (bool, string) {
			if name != strings.ToLower(name) {
				return false, fmt.Sprintf("directory name '%s' should be lowercase (recommended: '%s')", name, strings.ToLower(name))
			}
			return true, ""
		},
	})

	v.AddRule(Rule{
		Name:        "directory-no-python-keywords",
		Description: "Directory names must not be Python keywords",
		Severity:    SeverityMust,
		Applies:     fsobj.TypeDirectory,
		Validate: func(name string) (bool, string) {
			pythonKeywords := []string{
				"False", "None", "True", "and", "as", "assert", "async", "await",
				"break", "class", "continue", "def", "del", "elif", "else", "except",
				"finally", "for", "from", "global", "if", "import", "in", "is",
				"lambda", "nonlocal", "not", "or", "pass", "raise", "return",
				"try", "while", "with", "yield",
			}
			for _, keyword := range pythonKeywords {
				if name == keyword {
					return false, fmt.Sprintf("directory name '%s' is a Python keyword and is not allowed", name)
				}
			}
			return true, ""
		},
	})

	// Archive rules (similar to files)
	v.AddRule(Rule{
		Name:        "archive-no-spaces",
		Description: "Archive names must not contain spaces",
		Severity:    SeverityMust,
		Applies:     fsobj.TypeArchive,
		Validate: func(name string) (bool, string) {
			if strings.Contains(name, " ") {
				return false, fmt.Sprintf("archive name '%s' contains spaces", name)
			}
			return true, ""
		},
	})

	v.AddRule(Rule{
		Name:        "archive-must-have-extension",
		Description: "Archive names must have an appropriate extension",
		Severity:    SeverityMust,
		Applies:     fsobj.TypeArchive,
		Validate: func(name string) (bool, string) {
			validExts := []string{".zip", ".tar", ".tar.gz", ".tgz", ".tar.bz2", ".tbz2", ".gz", ".bz2", ".7z", ".rar"}
			ext := filepath.Ext(name)
			
			// Check for compound extensions like .tar.gz
			if strings.HasSuffix(name, ".tar.gz") || strings.HasSuffix(name, ".tar.bz2") {
				return true, ""
			}
			
			for _, validExt := range validExts {
				if ext == validExt {
					return true, ""
				}
			}
			return false, fmt.Sprintf("archive name '%s' must have a valid archive extension (%s)", name, strings.Join(validExts, ", "))
		},
	})

	// Link rules (more permissive, but no spaces)
	v.AddRule(Rule{
		Name:        "link-no-spaces",
		Description: "Link names should not contain spaces",
		Severity:    SeverityWant,
		Applies:     fsobj.TypeLink,
		Validate: func(name string) (bool, string) {
			if strings.Contains(name, " ") {
				return false, fmt.Sprintf("link name '%s' contains spaces (not recommended)", name)
			}
			return true, ""
		},
	})

	// Date format rule - suggests ISO 8601 format (configurable via naming_rules.rules.prefer_iso_dates)
	v.AddRule(Rule{
		Name:        "prefer-iso-dates",
		Description: "Dates in names should use ISO 8601 format (YYYY-MM-DD)",
		Severity:    SeverityWant,
		Applies:     "", // Applies to all types
		Validate: func(name string) (bool, string) {
			// Check if rule is enabled (default true if not specified)
			// This will be checked by viper if available, otherwise skip
			// Note: This check happens at validation time, not rule creation time
			
			// Remove extension for checking
			nameWithoutExt := name
			if ext := filepath.Ext(name); ext != "" {
				nameWithoutExt = name[:len(name)-len(ext)]
			}

			suggestion := detectNonISODate(nameWithoutExt)
			if suggestion != "" {
				return false, suggestion
			}
			return true, ""
		},
	})

	// File extension specific rules
	v.addExtensionRules()
	
	// Special filename rules
	v.addSpecialFilenameRules()
}

// addExtensionRules adds rules that apply to specific file extensions
func (v *Validator) addExtensionRules() {
	// Rule: .log files should start with ISO date
	v.AddRule(Rule{
		Name:        "log-files-iso-date-prefix",
		Description: "Log files should start with ISO 8601 date (YYYY-MM-DD)",
		Severity:    SeverityWant,
		Applies:     fsobj.TypeFile,
		Validate: func(name string) (bool, string) {
			// Only apply to .log files
			if !strings.HasSuffix(strings.ToLower(name), ".log") {
				return true, ""
			}

			// Extract basename without extension
			basename := name[:len(name)-4] // Remove .log

			// Check if it starts with ISO date pattern YYYY-MM-DD
			isoDatePrefix := regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})`)
			if !isoDatePrefix.MatchString(basename) {
				return false, fmt.Sprintf("log file '%s' should start with ISO date (e.g., 2025-11-23-app.log)", name)
			}

			return true, ""
		},
	})

	// Rule: Prefer .7z over .zip for archives
	v.AddRule(Rule{
		Name:        "prefer-7z-archives",
		Description: "7z archives are preferred over zip for better compression",
		Severity:    SeverityWant,
		Applies:     fsobj.TypeArchive,
		Validate: func(name string) (bool, string) {
			// Only flag .zip files
			if strings.HasSuffix(strings.ToLower(name), ".zip") {
				basename := name[:len(name)-4]
				return false, fmt.Sprintf("'%s' uses .zip format. Consider using .7z for better compression (e.g., %s.7z)", name, basename)
			}
			return true, ""
		},
	})
}

// addSpecialFilenameRules adds rules for specific well-known filenames
func (v *Validator) addSpecialFilenameRules() {
	// Rule: README should be uppercase
	v.AddRule(Rule{
		Name:        "readme-uppercase",
		Description: "README files should use uppercase naming convention",
		Severity:    SeverityWant,
		Applies:     fsobj.TypeFile,
		Validate: func(name string) (bool, string) {
			// Check if filename (case-insensitive) is readme with any extension
			lower := strings.ToLower(name)
			
			// Match readme, readme.md, readme.txt, etc.
			if strings.HasPrefix(lower, "readme") {
				// Check if it's not already uppercase
				if !strings.HasPrefix(name, "README") {
					// Suggest uppercase version
					suggested := "README" + name[6:] // Keep extension as-is
					return false, fmt.Sprintf("'%s' should use uppercase convention (recommended: '%s')", name, suggested)
				}
			}
			
			return true, ""
		},
	})

	// Rule: Common files that should be uppercase
	commonUppercaseFiles := []string{
		"license", "licence", "changelog", "contributing", 
		"authors", "copying", "install", "makefile",
	}
	
	v.AddRule(Rule{
		Name:        "common-files-uppercase",
		Description: "Common project files should use uppercase naming",
		Severity:    SeverityWant,
		Applies:     fsobj.TypeFile,
		Validate: func(name string) (bool, string) {
			// Get basename without extension
			nameWithoutExt := name
			ext := filepath.Ext(name)
			if ext != "" {
				nameWithoutExt = name[:len(name)-len(ext)]
			}
			
			lower := strings.ToLower(nameWithoutExt)
			
			for _, commonFile := range commonUppercaseFiles {
				if lower == commonFile {
					// Check if it's not already uppercase
					if nameWithoutExt != strings.ToUpper(lower) {
						suggested := strings.ToUpper(lower) + ext
						return false, fmt.Sprintf("'%s' should use uppercase convention (recommended: '%s')", name, suggested)
					}
				}
			}
			
			return true, ""
		},
	})
}

// detectNonISODate detects common date patterns that are not ISO 8601
func detectNonISODate(name string) string {
	// Pattern 1: MM-DD-YYYY or DD-MM-YYYY (with - or _ or .)
	// Examples: 11-23-2025, 23-11-2025, 11_23_2025
	mmddyyyyPattern := regexp.MustCompile(`\b(\d{1,2})[-_\.](\d{1,2})[-_\.](\d{4})\b`)
	if matches := mmddyyyyPattern.FindStringSubmatch(name); matches != nil {
		first, second, year := matches[1], matches[2], matches[3]
		// Check if it looks like a date (both parts <= 31)
		if isLikelyDate(first, second) {
			return fmt.Sprintf("'%s' appears to contain a date in MM-DD-YYYY or DD-MM-YYYY format. Consider ISO 8601: YYYY-MM-DD (e.g., %s-%s-%s)", name, year, first, second)
		}
	}

	// Pattern 2: MMDDYYYY or DDMMYYYY (8 digits, sometimes with separators)
	// Examples: 11232025, 23112025, 11_23_2025
	// Check for 8 consecutive digits OR pattern with underscores
	mmddyyyyCompactPattern := regexp.MustCompile(`(\d{2})[_]?(\d{2})[_]?(\d{4})`)
	if matches := mmddyyyyCompactPattern.FindStringSubmatch(name); matches != nil {
		first, second, year := matches[1], matches[2], matches[3]
		// Make sure these are potential date values and year starts with 19 or 20
		if isLikelyDate(first, second) && (strings.HasPrefix(year, "20") || strings.HasPrefix(year, "19")) {
			// Don't flag if it's already ISO format (year first)
			if !strings.HasPrefix(name, year) {
				return fmt.Sprintf("'%s' appears to contain a date in MMDDYYYY or DD_MM_YYYY format. Consider ISO 8601: YYYY-MM-DD (e.g., %s-%s-%s)", name, year, first, second)
			}
		}
	}

	// Pattern 3: MM/DD/YYYY (with slashes, though unlikely in filenames)
	// This would be caught by other rules as invalid filename chars, but check anyway
	slashPattern := regexp.MustCompile(`(\d{1,2})/(\d{1,2})/(\d{4})`)
	if matches := slashPattern.FindStringSubmatch(name); matches != nil {
		first, second, year := matches[1], matches[2], matches[3]
		if isLikelyDate(first, second) {
			return fmt.Sprintf("'%s' contains a date with slashes. Use ISO 8601: YYYY-MM-DD (e.g., %s-%s-%s)", name, year, first, second)
		}
	}

	// Pattern 4: Month name formats (e.g., Nov-23-2025, 23-Nov-2025)
	// Examples: Nov-23-2025, November-23-2025, 23-Nov-2025
	monthNamePattern := regexp.MustCompile(`\b(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec|January|February|March|April|May|June|July|August|September|October|November|December)[-_\.]?(\d{1,2})[-_\.]?(\d{4})\b`)
	if matches := monthNamePattern.FindStringSubmatch(name); matches != nil {
		monthStr, day, year := matches[1], matches[2], matches[3]
		monthNum := monthNameToNumber(monthStr)
		return fmt.Sprintf("'%s' contains a date with month name '%s'. Consider ISO 8601: YYYY-MM-DD (e.g., %s-%s-%s)", name, monthStr, year, monthNum, padZero(day))
	}

	// Pattern 5: Reverse month name (day first)
	reversMonthPattern := regexp.MustCompile(`\b(\d{1,2})[-_\.]?(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec|January|February|March|April|May|June|July|August|September|October|November|December)[-_\.]?(\d{4})\b`)
	if matches := reversMonthPattern.FindStringSubmatch(name); matches != nil {
		day, monthStr, year := matches[1], matches[2], matches[3]
		monthNum := monthNameToNumber(monthStr)
		return fmt.Sprintf("'%s' contains a date with month name '%s'. Consider ISO 8601: YYYY-MM-DD (e.g., %s-%s-%s)", name, monthStr, year, monthNum, padZero(day))
	}

	// Pattern 6: YYYYMMDD without separators (this is actually acceptable, but suggest with separators)
	// Only flag if it's clearly meant as a date (like starts with 20xx)
	yyyymmddCompactPattern := regexp.MustCompile(`\b(20\d{2})(\d{2})(\d{2})\b`)
	if matches := yyyymmddCompactPattern.FindStringSubmatch(name); matches != nil {
		year, month, day := matches[1], matches[2], matches[3]
		// Only suggest if month and day are valid
		if isValidMonth(month) && isValidDay(day) {
			return fmt.Sprintf("'%s' contains date in YYYYMMDD format. Consider adding separators: YYYY-MM-DD (e.g., %s-%s-%s)", name, year, month, day)
		}
	}

	return ""
}

// isLikelyDate checks if two number strings could represent month/day
func isLikelyDate(first, second string) bool {
	f, err1 := strconv.Atoi(first)
	s, err2 := strconv.Atoi(second)
	if err1 != nil || err2 != nil {
		return false
	}
	// Both should be valid month or day values (1-31)
	return (f >= 1 && f <= 31) && (s >= 1 && s <= 31)
}

// isValidMonth checks if a string is a valid month (01-12)
func isValidMonth(month string) bool {
	m, err := strconv.Atoi(month)
	return err == nil && m >= 1 && m <= 12
}

// isValidDay checks if a string is a valid day (01-31)
func isValidDay(day string) bool {
	d, err := strconv.Atoi(day)
	return err == nil && d >= 1 && d <= 31
}

// monthNameToNumber converts month name to zero-padded number
func monthNameToNumber(month string) string {
	months := map[string]string{
		"Jan": "01", "January": "01",
		"Feb": "02", "February": "02",
		"Mar": "03", "March": "03",
		"Apr": "04", "April": "04",
		"May": "05",
		"Jun": "06", "June": "06",
		"Jul": "07", "July": "07",
		"Aug": "08", "August": "08",
		"Sep": "09", "September": "09",
		"Oct": "10", "October": "10",
		"Nov": "11", "November": "11",
		"Dec": "12", "December": "12",
	}
	if num, ok := months[month]; ok {
		return num
	}
	return "XX"
}

// padZero pads a number string with leading zero if needed
func padZero(num string) string {
	if len(num) == 1 {
		return "0" + num
	}
	return num
}

// ValidatePath is a convenience function to validate a path
func ValidatePath(path string, objType fsobj.ObjectType) []Violation {
	validator := NewValidator()
	return validator.Validate(path, objType)
}

// CustomRule allows creating rules from configuration
type CustomRule struct {
	Name        string         `yaml:"name"`
	Description string         `yaml:"description"`
	Severity    string         `yaml:"severity"` // "MUST" or "WANT"
	Applies     string         `yaml:"applies"`  // "file", "directory", "link", "archive"
	Pattern     string         `yaml:"pattern"`  // Regex pattern to match (violation if matches)
	Message     string         `yaml:"message"`  // Custom error message
}

// ToRule converts a CustomRule to a Rule
func (cr *CustomRule) ToRule() Rule {
	severity := SeverityWant
	if strings.ToUpper(cr.Severity) == "MUST" {
		severity = SeverityMust
	}

	var applies fsobj.ObjectType
	switch strings.ToLower(cr.Applies) {
	case "file":
		applies = fsobj.TypeFile
	case "directory", "dir":
		applies = fsobj.TypeDirectory
	case "link", "symlink":
		applies = fsobj.TypeLink
	case "archive":
		applies = fsobj.TypeArchive
	default:
		applies = "" // Applies to all types
	}

	pattern := regexp.MustCompile(cr.Pattern)

	return Rule{
		Name:        cr.Name,
		Description: cr.Description,
		Severity:    severity,
		Applies:     applies,
		Validate: func(name string) (bool, string) {
			if pattern.MatchString(name) {
				message := cr.Message
				if message == "" {
					message = fmt.Sprintf("'%s' matches prohibited pattern '%s'", name, cr.Pattern)
				}
				return false, message
			}
			return true, ""
		},
	}
}
