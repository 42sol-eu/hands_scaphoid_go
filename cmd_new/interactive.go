package cmd_new

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// Interactive prompt examples and utilities

// Confirm asks a yes/no question
func Confirm(message string, defaultValue bool) (bool, error) {
	result := defaultValue
	prompt := &survey.Confirm{
		Message: message,
		Default: defaultValue,
	}
	err := survey.AskOne(prompt, &result)
	return result, err
}

// Select asks user to choose from a list of options
func Select(message string, options []string) (string, error) {
	var result string
	prompt := &survey.Select{
		Message: message,
		Options: options,
	}
	err := survey.AskOne(prompt, &result)
	return result, err
}

// MultiSelect asks user to choose multiple options
func MultiSelect(message string, options []string) ([]string, error) {
	var result []string
	prompt := &survey.MultiSelect{
		Message: message,
		Options: options,
	}
	err := survey.AskOne(prompt, &result)
	return result, err
}

// Input asks for text input
func Input(message string, defaultValue string) (string, error) {
	result := defaultValue
	prompt := &survey.Input{
		Message: message,
		Default: defaultValue,
	}
	err := survey.AskOne(prompt, &result)
	return result, err
}

// Password asks for password input (hidden)
func Password(message string) (string, error) {
	var result string
	prompt := &survey.Password{
		Message: message,
	}
	err := survey.AskOne(prompt, &result)
	return result, err
}

// Demo command showing interactive prompts
var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Interactive prompt examples",
	Long:  `Demonstrates various interactive prompt types available in Scaphoid.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(formatSectionHeader("Interactive Prompts Demo", "Testing various prompt types"))
		
		// Confirmation prompt
		if confirmed, err := Confirm("Do you want to continue?", true); err == nil {
			if confirmed {
				fmt.Println(wrapSuccessMessage("User confirmed!"))
			} else {
				fmt.Println(wrapErrorMessage("User declined."))
			}
		}
		
		fmt.Println()
		
		// Selection prompt
		if choice, err := Select("Choose an operation:", []string{
			"Create file",
			"Delete file",
			"Copy file",
			"Move file",
		}); err == nil {
			fmt.Println(wrapSuccessMessage(fmt.Sprintf("Selected: %s", choice)))
		}
		
		fmt.Println()
		
		// Multi-select prompt
		if choices, err := MultiSelect("Select file types to process:", []string{
			".txt",
			".md",
			".go",
			".yaml",
			".json",
		}); err == nil {
			fmt.Println(wrapSuccessMessage(fmt.Sprintf("Selected: %v", choices)))
		}
		
		fmt.Println()
		
		// Text input prompt
		if input, err := Input("Enter a filename:", "example.txt"); err == nil {
			fmt.Println(wrapSuccessMessage(fmt.Sprintf("Filename: %s", input)))
		}
	},
}

func init() {
	// Will be added to root in root.go
}
