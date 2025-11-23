package cmd_new

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/actions"
	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive operations",
	Long:  `Perform various operations on archive files.`,
}

var archiveCreateCmd = &cobra.Command{
	Use:   "create [archive-path]",
	Short: "Create an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		source, _ := cmd.Flags().GetString("source")

		action := &actions.CreateAction{
			Force: force,
		}

		options := make(map[string]interface{})
		if source != "" {
			options["source"] = source
		}

		result := action.Execute(fsobj.TypeArchive, args[0], options)
		if result.Success {
			fmt.Println(result.Message)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

var archiveListCmd = &cobra.Command{
	Use:   "list [path]",
	Short: "List archive contents",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateArchivePath(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		action := &actions.ListAction{}

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

var archiveCopyCmd = &cobra.Command{
	Use:   "copy [source] [destination]",
	Short: "Copy an archive",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		action := &actions.CopyAction{
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

var archiveMoveCmd = &cobra.Command{
	Use:   "move [source] [destination]",
	Short: "Move an archive",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
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

var archiveDeleteCmd = &cobra.Command{
	Use:   "delete [archive-path]",
	Short: "Delete an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		action := &actions.DeleteAction{
			Force: force,
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

var archiveInfoCmd = &cobra.Command{
	Use:   "info [path]",
	Short: "Get archive information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateArchivePath(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		showAll, _ := cmd.Flags().GetBool("all")
		action := &actions.InfoAction{
			ShowAll: showAll,
		}

		result := action.Execute(args[0])
		if result.Success {
			if info, ok := result.Data.(map[string]interface{}); ok {
				for key, value := range info {
					fmt.Printf("%s: %v\n", key, value)
				}
			}
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

var archiveCompareCmd = &cobra.Command{
	Use:   "compare [path1] [path2]",
	Short: "Compare two archives",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateArchivePath(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if err := validateArchivePath(args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		deep, _ := cmd.Flags().GetBool("deep")
		action := &actions.CompareAction{
			Deep: deep,
		}

		result := action.Execute(args[0], args[1])
		if result.Success {
			fmt.Println(result.Message)
			if comparison, ok := result.Data.(map[string]interface{}); ok {
				for key, value := range comparison {
					fmt.Printf("%s: %v\n", key, value)
				}
			}
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

var archiveBackupCmd = &cobra.Command{
	Use:   "backup [source] [backup-directory]",
	Short: "Backup an archive",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateArchivePath(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		timestamp, _ := cmd.Flags().GetBool("timestamp")
		compress, _ := cmd.Flags().GetBool("compress")
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

// validateArchivePath checks if the given path is actually an archive
func validateArchivePath(path string) error {
	obj, err := fsobj.NewFSObject(path)
	if err != nil {
		return err
	}

	if obj.Exists() && obj.Type() != fsobj.TypeArchive {
		return fmt.Errorf("path '%s' exists but is not an archive (it's a %s)", path, obj.Type())
	}

	return nil
}

func init() {
	// Add persistent flags
	archiveCmd.PersistentFlags().BoolP("force", "f", false, "Force operation")

	// Add subcommand-specific flags
	archiveCreateCmd.Flags().StringP("source", "s", "", "Source directory to archive")
	archiveInfoCmd.Flags().BoolP("all", "a", false, "Show all available information")
	archiveCompareCmd.Flags().BoolP("deep", "d", false, "Perform deep comparison")
	archiveBackupCmd.Flags().BoolP("timestamp", "t", true, "Add timestamp to backup name")
	archiveBackupCmd.Flags().BoolP("compress", "c", false, "Compress backup")

	// Add all subcommands to archive command
	archiveCmd.AddCommand(archiveCreateCmd)
	archiveCmd.AddCommand(archiveListCmd)
	archiveCmd.AddCommand(archiveCopyCmd)
	archiveCmd.AddCommand(archiveMoveCmd)
	archiveCmd.AddCommand(archiveDeleteCmd)
	archiveCmd.AddCommand(archiveInfoCmd)
	archiveCmd.AddCommand(archiveCompareCmd)
	archiveCmd.AddCommand(archiveBackupCmd)
}
