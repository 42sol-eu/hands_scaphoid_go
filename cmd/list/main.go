package main

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/actions"
	"github.com/spf13/cobra"
)

var (
	showHidden  bool
	longFormat  bool
	recursive   bool
	sortBy      string
	reverseSort bool
)

var rootCmd = &cobra.Command{
	Use:   "list [path]",
	Short: "List filesystem objects",
	Long:  `List the contents of directories, files, archives, and other filesystem objects.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := &actions.ListAction{
			ShowHidden:  showHidden,
			LongFormat:  longFormat,
			Recursive:   recursive,
			SortBy:      sortBy,
			ReverseSort: reverseSort,
		}

		result := action.Execute(args[0])
		if result.Success {
			if output, ok := result.Data.([]string); ok {
				for _, line := range output {
					fmt.Println(line)
				}
			}
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&showHidden, "all", "a", false, "Show hidden files")
	rootCmd.Flags().BoolVarP(&longFormat, "long", "l", false, "Use long listing format")
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "List subdirectories recursively")
	rootCmd.Flags().StringVar(&sortBy, "sort", "", "Sort by: name, size, time")
	rootCmd.Flags().BoolVar(&reverseSort, "reverse", false, "Reverse sort order")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
