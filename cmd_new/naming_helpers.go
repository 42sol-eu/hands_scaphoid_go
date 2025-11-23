package cmd_new

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
	"github.com/42sol-eu/hands_scaphoid_go/pkg/naming"
	"github.com/spf13/viper"
)

// validateNamingRules checks if a path conforms to naming rules
func validateNamingRules(path string, objType fsobj.ObjectType) error {
	// Check if validation is enabled
	if !viper.GetBool("naming_rules.enabled") {
		return nil
	}

	validator := naming.NewValidator()
	violations := validator.Validate(path, objType)

	// Separate errors and warnings
	var errors []naming.Violation
	var warnings []naming.Violation
	
	for _, v := range violations {
		if v.IsError() {
			errors = append(errors, v)
		} else {
			warnings = append(warnings, v)
		}
	}

	// Show warnings if enabled
	if viper.GetBool("naming_rules.show_warnings") && len(warnings) > 0 {
		for _, w := range warnings {
			fmt.Fprintf(os.Stderr, "⚠  %s\n", colorizeKeyword("Warning: ")+w.Message)
		}
	}

	// Check for errors
	if viper.GetBool("naming_rules.enforce_must") && len(errors) > 0 {
		// Return first error
		return fmt.Errorf("naming validation failed: %s", errors[0].Message)
	}

	return nil
}
