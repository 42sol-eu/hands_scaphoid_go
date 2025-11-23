package cmd_new

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/actions"
	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link operations",
	Long:  `Perform various operations on symbolic links.`,
}

var linkCreateCmd = &cobra.Command{
	Use:   "create [path]",
	Short: "Create a symbolic link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		if target == "" {
			fmt.Fprintf(os.Stderr, "Error: --target is required for link creation\n")
			os.Exit(1)
		}

		force, _ := cmd.Flags().GetBool("force")
		action := &actions.CreateAction{
			Force: force,
		}

		options := map[string]interface{}{
			"target": target,
		}

		result := action.Execute(fsobj.TypeLink, args[0], options)
		if result.Success {
			fmt.Println(result.Message)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

var linkListCmd = &cobra.Command{
	Use:   "list [path]",
	Short: "List link target contents",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateLinkPath(args[0]); err != nil {
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

var linkCopyCmd = &cobra.Command{
	Use:   "copy [source] [destination]",
	Short: "Copy a link",
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

var linkMoveCmd = &cobra.Command{
	Use:   "move [source] [destination]",
	Short: "Move a link",
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

var linkDeleteCmd = &cobra.Command{
	Use:   "delete [path]",
	Short: "Delete a link",
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

var linkInfoCmd = &cobra.Command{
	Use:   "info [path]",
	Short: "Get link information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateLinkPath(args[0]); err != nil {
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

var linkCompareCmd = &cobra.Command{
	Use:   "compare [path1] [path2]",
	Short: "Compare two links",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateLinkPath(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if err := validateLinkPath(args[1]); err != nil {
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

var linkBackupCmd = &cobra.Command{
	Use:   "backup [source] [backup-directory]",
	Short: "Backup a link",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateLinkPath(args[0]); err != nil {
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

// validateLinkPath checks if the given path is actually a symbolic link
func validateLinkPath(path string) error {
	obj, err := fsobj.NewFSObject(path)
	if err != nil {
		return err
	}

	if obj.Exists() && obj.Type() != fsobj.TypeLink {
		return fmt.Errorf("path '%s' exists but is not a symbolic link (it's a %s)", path, obj.Type())
	}

	return nil
}

func init() {
	// Add persistent flags
	linkCmd.PersistentFlags().BoolP("force", "f", false, "Force operation")

	// Add subcommand-specific flags
	linkCreateCmd.Flags().StringP("target", "t", "", "Target path for the symbolic link")
	linkInfoCmd.Flags().BoolP("all", "a", false, "Show all available information")
	linkCompareCmd.Flags().BoolP("deep", "d", false, "Perform deep comparison")
	linkBackupCmd.Flags().BoolP("timestamp", "t", true, "Add timestamp to backup name")
	linkBackupCmd.Flags().BoolP("compress", "c", false, "Compress backup")

	// Add all subcommands to link command
	linkCmd.AddCommand(linkCreateCmd)
	linkCmd.AddCommand(linkListCmd)
	linkCmd.AddCommand(linkCopyCmd)
	linkCmd.AddCommand(linkMoveCmd)
	linkCmd.AddCommand(linkDeleteCmd)
	linkCmd.AddCommand(linkInfoCmd)
	linkCmd.AddCommand(linkCompareCmd)
	linkCmd.AddCommand(linkBackupCmd)
}
