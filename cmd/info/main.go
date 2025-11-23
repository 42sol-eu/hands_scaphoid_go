package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/actions"
	"github.com/spf13/cobra"
)

var showAll bool

var rootCmd = &cobra.Command{
	Use:   "info [path]",
	Short: "Get information about filesystem objects",
	Long:  `Display detailed information about files, directories, links, and other filesystem objects.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := &actions.InfoAction{
			ShowAll: showAll,
		}

		result := action.Execute(args[0])
		if result.Success {
			if info, ok := result.Data.(map[string]interface{}); ok {
				// Pretty print the information
				data, _ := json.MarshalIndent(info, "", "  ")
				fmt.Println(string(data))
			}
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all available information")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
