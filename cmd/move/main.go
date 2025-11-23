package main

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/actions"
	"github.com/spf13/cobra"
)

var force bool

var rootCmd = &cobra.Command{
	Use:   "move [source] [destination]",
	Short: "Move/rename filesystem objects",
	Long:  `Move or rename files, directories, links, and other filesystem objects.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		action := &actions.MoveAction{
			Force: force,
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
	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "Force move, overwrite existing")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
