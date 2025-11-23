package main

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/actions"
	"github.com/spf13/cobra"
)

var (
	recursive    bool
	force        bool
	preserveMeta bool
)

var rootCmd = &cobra.Command{
	Use:   "copy [source] [destination]",
	Short: "Copy filesystem objects",
	Long:  `Copy files, directories, links, and other filesystem objects.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		action := &actions.CopyAction{
			Recursive:    recursive,
			Force:        force,
			PreserveMeta: preserveMeta,
		}

		result := action.Execute(args[0], args[1])
		if result.Success {
			fmt.Println(result.Message)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Copy directories recursively")
	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "Force copy, overwrite existing")
	rootCmd.Flags().BoolVarP(&preserveMeta, "preserve", "p", false, "Preserve metadata (permissions, timestamps)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
