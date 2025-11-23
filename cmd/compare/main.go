package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/actions"
	"github.com/spf13/cobra"
)

var deep bool

var rootCmd = &cobra.Command{
	Use:   "compare [path1] [path2]",
	Short: "Compare filesystem objects",
	Long:  `Compare two files, directories, or other filesystem objects.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		action := &actions.CompareAction{
			Deep: deep,
		}

		result := action.Execute(args[0], args[1])
		if result.Success {
			if comparison, ok := result.Data.(map[string]interface{}); ok {
				data, _ := json.MarshalIndent(comparison, "", "  ")
				fmt.Println(string(data))
			}
			fmt.Println(result.Message)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&deep, "deep", "d", false, "Perform deep comparison (content-based)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
