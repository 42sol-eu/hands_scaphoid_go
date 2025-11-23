package cmd_new

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sectionCmd = &cobra.Command{
	Use:   "section [title] [details]",
	Short: "Print a markdown section header",
	Long:  `Print a markdown-style section header (## title) with optional details.`,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		details := ""
		if len(args) > 1 {
			details = args[1]
		}
		
		fmt.Print(formatSectionHeader(title, details))
	},
}

var separatorCmd = &cobra.Command{
	Use:   "separator",
	Short: "Print a horizontal rule separator",
	Long:  `Print a markdown-style horizontal rule (---) for visual separation.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---")
	},
}

func init() {
	// These will be added to rootCmd in root.go
}
