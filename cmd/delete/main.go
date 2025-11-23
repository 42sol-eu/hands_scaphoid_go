package main

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/actions"
	"github.com/spf13/cobra"
)

var (
	force     bool
	recursive bool
)

var rootCmd = &cobra.Command{
	Use:   "delete [path]",
	Short: "Delete filesystem objects",
	Long:  `Delete files, directories, links, and other filesystem objects.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := &actions.DeleteAction{
			Force:     force,
			Recursive: recursive,
		}

		result := action.Execute(args[0])
		if result.Success {
			fmt.Println(result.Message)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "Force deletion without confirmation")
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Delete directories recursively")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
