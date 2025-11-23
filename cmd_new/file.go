package cmd_new

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/actions"
	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "File operations",
	Long:  `Perform various operations on files.`,
}

var fileCreateCmd = &cobra.Command{
	Use:   "create [path]",
	Short: "Create a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		
		if IsVerbose() {
			fmt.Fprintf(os.Stderr, "[VERBOSE] Creating file: %s\n", colorizePath(args[0]))
			fmt.Fprintf(os.Stderr, "[VERBOSE] Force: %v\n", force)
		}
		
		// Validate naming rules
		if err := validateNamingRules(args[0], fsobj.TypeFile); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", formatErrorMessage(err.Error()))
			os.Exit(1)
		}
		
		action := &actions.CreateAction{
			Force: force,
		}

		result := action.Execute(fsobj.TypeFile, args[0], nil)
		if result.Success {
			fmt.Println(formatCreateMessage("file", args[0]))
		} else {
			fmt.Fprintln(os.Stderr, formatErrorMessage(fmt.Sprintf("Error: %v", result.Error)))
			os.Exit(1)
		}
	},
}

var fileListCmd = &cobra.Command{
	Use:   "list [path]",
	Short: "List file contents",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateFilePath(args[0]); err != nil {
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

var fileCopyCmd = &cobra.Command{
	Use:   "copy [source] [destination]",
	Short: "Copy a file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		
		if IsVerbose() {
			fmt.Fprintf(os.Stderr, "[VERBOSE] Copying file from %s to %s\n", colorizePath(args[0]), colorizePath(args[1]))
			fmt.Fprintf(os.Stderr, "[VERBOSE] Force: %v\n", force)
		}
		
		action := &actions.CopyAction{
			Force: force,
		}

		result := action.Execute(args[0], args[1])
		if result.Success {
			fmt.Println(formatCopyMoveMessage("copied", args[0], args[1]))
		} else {
			fmt.Fprintln(os.Stderr, formatErrorMessage(fmt.Sprintf("Error: %v", result.Error)))
			os.Exit(1)
		}
	},
}

var fileMoveCmd = &cobra.Command{
	Use:   "move [source] [destination]",
	Short: "Move a file",
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

var fileDeleteCmd = &cobra.Command{
	Use:   "delete [path]",
	Short: "Delete a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		
		if IsVerbose() {
			fmt.Fprintf(os.Stderr, "[VERBOSE] Deleting file: %s\n", colorizePath(args[0]))
			fmt.Fprintf(os.Stderr, "[VERBOSE] Force: %v\n", force)
		}
		
		// Interactive confirmation if not forced
		if !force {
			confirmed, err := Confirm(fmt.Sprintf("Delete %s?", args[0]), false)
			if err != nil {
				fmt.Fprintln(os.Stderr, formatErrorMessage(fmt.Sprintf("Error: %v", err)))
				os.Exit(1)
			}
			if !confirmed {
				fmt.Println(wrapErrorMessage("Operation cancelled by user"))
				os.Exit(1)
			}
		}
		
		action := &actions.DeleteAction{
			Force: force,
		}

		result := action.Execute(args[0])
		if result.Success {
			fmt.Println(formatSuccessMessage("deleted", args[0]))
		} else {
			fmt.Fprintln(os.Stderr, formatErrorMessage(fmt.Sprintf("Error: %v", result.Error)))
			os.Exit(1)
		}
	},
}

var fileInfoCmd = &cobra.Command{
	Use:   "info [path]",
	Short: "Get file information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateFilePath(args[0]); err != nil {
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

var fileCompareCmd = &cobra.Command{
	Use:   "compare [path1] [path2]",
	Short: "Compare two files",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateFilePath(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if err := validateFilePath(args[1]); err != nil {
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

var fileBackupCmd = &cobra.Command{
	Use:   "backup [source] [backup-directory]",
	Short: "Backup a file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateFilePath(args[0]); err != nil {
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

// validateFilePath checks if the given path is actually a file
func validateFilePath(path string) error {
	obj, err := fsobj.NewFSObject(path)
	if err != nil {
		return err
	}

	if obj.Exists() && obj.Type() != fsobj.TypeFile {
		return fmt.Errorf("path '%s' exists but is not a file (it's a %s)", path, obj.Type())
	}

	return nil
}

func init() {
	// Add persistent flags
	fileCmd.PersistentFlags().BoolP("force", "f", false, "Force operation")

	// Add subcommand-specific flags
	fileInfoCmd.Flags().BoolP("all", "a", false, "Show all available information")
	fileCompareCmd.Flags().BoolP("deep", "d", false, "Perform deep comparison")
	fileBackupCmd.Flags().BoolP("timestamp", "t", true, "Add timestamp to backup name")
	fileBackupCmd.Flags().BoolP("compress", "c", false, "Compress backup")

	// Add all subcommands to file command
	fileCmd.AddCommand(fileCreateCmd)
	fileCmd.AddCommand(fileListCmd)
	fileCmd.AddCommand(fileCopyCmd)
	fileCmd.AddCommand(fileMoveCmd)
	fileCmd.AddCommand(fileDeleteCmd)
	fileCmd.AddCommand(fileInfoCmd)
	fileCmd.AddCommand(fileCompareCmd)
	fileCmd.AddCommand(fileBackupCmd)
}
