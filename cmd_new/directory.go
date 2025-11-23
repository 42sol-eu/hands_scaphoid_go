package cmd_new

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/actions"
	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
	"github.com/spf13/cobra"
)

var directoryCmd = &cobra.Command{
	Use:   "directory",
	Short: "Directory operations",
	Long:  `Perform various operations on directories.`,
}

var directoryCreateCmd = &cobra.Command{
	Use:   "create [path]",
	Short: "Create a directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		recursive, _ := cmd.Flags().GetBool("recursive")
		
		// Validate naming rules
		if err := validateNamingRules(args[0], fsobj.TypeDirectory); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", formatErrorMessage(err.Error()))
			os.Exit(1)
		}
		
		action := &actions.CreateAction{
			Force:     force,
			Recursive: recursive,
		}

		result := action.Execute(fsobj.TypeDirectory, args[0], nil)
		if result.Success {
			fmt.Println(wrapSuccessMessage(result.Message))
		} else {
			fmt.Fprintln(os.Stderr, formatErrorMessage(fmt.Sprintf("Error: %v", result.Error)))
			os.Exit(1)
		}
	},
}

var directoryListCmd = &cobra.Command{
	Use:   "list [path]",
	Short: "List directory contents",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateDirectoryPath(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		showHidden, _ := cmd.Flags().GetBool("all")
		longFormat, _ := cmd.Flags().GetBool("long")
		recursive, _ := cmd.Flags().GetBool("recursive")
		action := &actions.ListAction{
			ShowHidden: showHidden,
			LongFormat: longFormat,
			Recursive:  recursive,
		}

		result := action.Execute(args[0])
		if result.Success {
			if output, ok := result.Data.([]string); ok {
				for _, line := range output {
					fmt.Println(line)
				}
			}
		} else {
			fmt.Fprintln(os.Stderr, formatErrorMessage(fmt.Sprintf("Error: %v", result.Error)))
			os.Exit(1)
		}
	},
}

var directoryCopyCmd = &cobra.Command{
	Use:   "copy [source] [destination]",
	Short: "Copy a directory",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		recursive, _ := cmd.Flags().GetBool("recursive")
		action := &actions.CopyAction{
			Recursive: recursive,
			Force:     force,
		}

		result := action.Execute(args[0], args[1])
		if result.Success {
			fmt.Println(wrapSuccessMessage(result.Message))
		} else {
			fmt.Fprintln(os.Stderr, formatErrorMessage(fmt.Sprintf("Error: %v", result.Error)))
			os.Exit(1)
		}
	},
}

var directoryMoveCmd = &cobra.Command{
	Use:   "move [source] [destination]",
	Short: "Move a directory",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		action := &actions.MoveAction{
			Force: force,
		}

		result := action.Execute(args[0], args[1])
		if result.Success {
			fmt.Println(wrapSuccessMessage(result.Message))
		} else {
			fmt.Fprintln(os.Stderr, formatErrorMessage(fmt.Sprintf("Error: %v", result.Error)))
			os.Exit(1)
		}
	},
}

var directoryDeleteCmd = &cobra.Command{
	Use:   "delete [path]",
	Short: "Delete a directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		recursive, _ := cmd.Flags().GetBool("recursive")
		action := &actions.DeleteAction{
			Force:     force,
			Recursive: recursive,
		}

		result := action.Execute(args[0])
		if result.Success {
			fmt.Println(wrapSuccessMessage(result.Message))
		} else {
			fmt.Fprintln(os.Stderr, formatErrorMessage(fmt.Sprintf("Error: %v", result.Error)))
			os.Exit(1)
		}
	},
}

var directoryInfoCmd = &cobra.Command{
	Use:   "info [path]",
	Short: "Get directory information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateDirectoryPath(args[0]); err != nil {
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
			fmt.Fprintln(os.Stderr, formatErrorMessage(fmt.Sprintf("Error: %v", result.Error)))
			os.Exit(1)
		}
	},
}

var directoryCompareCmd = &cobra.Command{
	Use:   "compare [path1] [path2]",
	Short: "Compare two directories",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateDirectoryPath(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if err := validateDirectoryPath(args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		deep, _ := cmd.Flags().GetBool("deep")
		action := &actions.CompareAction{
			Deep: deep,
		}

		result := action.Execute(args[0], args[1])
		if result.Success {
			fmt.Println(wrapSuccessMessage(result.Message))
			if comparison, ok := result.Data.(map[string]interface{}); ok {
				for key, value := range comparison {
					fmt.Printf("%s: %v\n", key, value)
				}
			}
		} else {
			fmt.Fprintln(os.Stderr, formatErrorMessage(fmt.Sprintf("Error: %v", result.Error)))
			os.Exit(1)
		}
	},
}

var directoryBackupCmd = &cobra.Command{
	Use:   "backup [source] [backup-directory]",
	Short: "Backup a directory",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateDirectoryPath(args[0]); err != nil {
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
			fmt.Println(wrapSuccessMessage(result.Message))
		} else {
			fmt.Fprintln(os.Stderr, formatErrorMessage(fmt.Sprintf("Error: %v", result.Error)))
			os.Exit(1)
		}
	},
}

// validateDirectoryPath checks if the given path is actually a directory
func validateDirectoryPath(path string) error {
	obj, err := fsobj.NewFSObject(path)
	if err != nil {
		return err
	}

	if obj.Exists() && obj.Type() != fsobj.TypeDirectory {
		return fmt.Errorf("path '%s' exists but is not a directory (it's a %s)", path, obj.Type())
	}

	return nil
}

func init() {
	// Add persistent flags
	directoryCmd.PersistentFlags().BoolP("force", "f", false, "Force operation")
	directoryCmd.PersistentFlags().BoolP("recursive", "r", false, "Recursive operation")

	// Add subcommand-specific flags
	directoryListCmd.Flags().BoolP("all", "a", false, "Show hidden files")
	directoryListCmd.Flags().BoolP("long", "l", false, "Long format")

	directoryInfoCmd.Flags().BoolP("all", "a", false, "Show all available information")
	directoryCompareCmd.Flags().BoolP("deep", "d", false, "Perform deep comparison")
	directoryBackupCmd.Flags().BoolP("timestamp", "t", true, "Add timestamp to backup name")
	directoryBackupCmd.Flags().BoolP("compress", "c", false, "Compress backup")

	// Add all subcommands to directory command
	directoryCmd.AddCommand(directoryCreateCmd)
	directoryCmd.AddCommand(directoryListCmd)
	directoryCmd.AddCommand(directoryCopyCmd)
	directoryCmd.AddCommand(directoryMoveCmd)
	directoryCmd.AddCommand(directoryDeleteCmd)
	directoryCmd.AddCommand(directoryInfoCmd)
	directoryCmd.AddCommand(directoryCompareCmd)
	directoryCmd.AddCommand(directoryBackupCmd)
}
