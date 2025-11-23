package main

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/actions"
	"github.com/spf13/cobra"
)

var (
	timestamp bool
	compress  bool
)

var rootCmd = &cobra.Command{
	Use:   "backup [source] [backup-directory]",
	Short: "Backup filesystem objects",
	Long:  `Create backups of files, directories, and other filesystem objects.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		action := &actions.BackupAction{
			Timestamp: timestamp,
			Compress:  compress,
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
	rootCmd.Flags().BoolVarP(&timestamp, "timestamp", "t", true, "Add timestamp to backup name")
	rootCmd.Flags().BoolVarP(&compress, "compress", "c", false, "Compress backup")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
