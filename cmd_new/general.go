package cmd_new

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/actions"
	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
	"github.com/spf13/cobra"
)

// General-purpose commands that work on any filesystem object

var copyCmd = &cobra.Command{
	Use:   "copy [source] [destination]",
	Short: "Copy filesystem objects",
	Long:  `Copy files, directories, links, and other filesystem objects.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		recursive, _ := cmd.Flags().GetBool("recursive")
		force, _ := cmd.Flags().GetBool("force")
		preserveMeta, _ := cmd.Flags().GetBool("preserve")

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

var moveCmd = &cobra.Command{
	Use:   "move [source] [destination]",
	Short: "Move filesystem objects",
	Long:  `Move files, directories, links, and other filesystem objects.`,
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

var deleteCmd = &cobra.Command{
	Use:   "delete [path]",
	Short: "Delete filesystem objects",
	Long:  `Delete files, directories, links, and other filesystem objects.`,
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
			fmt.Println(result.Message)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

var infoCmd = &cobra.Command{
	Use:   "info [path]",
	Short: "Get information about filesystem objects",
	Long:  `Get detailed information about files, directories, links, and other filesystem objects.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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

var compareCmd = &cobra.Command{
	Use:   "compare [path1] [path2]",
	Short: "Compare two filesystem objects",
	Long:  `Compare two files, directories, links, or other filesystem objects.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
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

var backupCmd = &cobra.Command{
	Use:   "backup [source] [backup-directory]",
	Short: "Backup filesystem objects",
	Long:  `Backup files, directories, links, and other filesystem objects.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
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

var listCmd = &cobra.Command{
	Use:   "list [path]",
	Short: "List filesystem objects",
	Long:  `List contents of files, directories, archives, and other filesystem objects.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

var createCmd = &cobra.Command{
	Use:   "create [type] [path]",
	Short: "Create filesystem objects",
	Long: `Create files, directories, links, archives, and other filesystem objects.

Type can be one of: file, directory, link, archive`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		recursive, _ := cmd.Flags().GetBool("recursive")
		target, _ := cmd.Flags().GetString("target")
		source, _ := cmd.Flags().GetString("source")

		action := &actions.CreateAction{
			Force:     force,
			Recursive: recursive,
		}

		// Parse the type argument
		var fsType fsobj.ObjectType
		switch args[0] {
		case "file", "f":
			fsType = fsobj.TypeFile
		case "directory", "dir", "d":
			fsType = fsobj.TypeDirectory
		case "link", "symlink", "l":
			fsType = fsobj.TypeLink
		case "archive", "a":
			fsType = fsobj.TypeArchive
		default:
			fmt.Fprintf(os.Stderr, "Error: unknown type '%s'. Valid types: file, directory, link, archive\n", args[0])
			os.Exit(1)
		}

		// Build options map
		options := make(map[string]interface{})
		if target != "" {
			options["target"] = target
		}
		if source != "" {
			options["source"] = source
		}

		result := action.Execute(fsType, args[1], options)
		if result.Success {
			fmt.Println(result.Message)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

func init() {
	// Copy command flags
	copyCmd.Flags().BoolP("recursive", "r", false, "Copy directories recursively")
	copyCmd.Flags().BoolP("force", "f", false, "Force copy, overwrite existing")
	copyCmd.Flags().BoolP("preserve", "p", false, "Preserve metadata (permissions, timestamps)")

	// Move command flags
	moveCmd.Flags().BoolP("force", "f", false, "Force move, overwrite existing")

	// Delete command flags
	deleteCmd.Flags().BoolP("force", "f", false, "Force delete without confirmation")
	deleteCmd.Flags().BoolP("recursive", "r", false, "Delete directories recursively")

	// Info command flags
	infoCmd.Flags().BoolP("all", "a", false, "Show all available information")

	// Compare command flags
	compareCmd.Flags().BoolP("deep", "d", false, "Perform deep comparison")

	// Backup command flags
	backupCmd.Flags().BoolP("timestamp", "t", true, "Add timestamp to backup name")
	backupCmd.Flags().BoolP("compress", "c", false, "Compress backup")

	// List command flags
	listCmd.Flags().BoolP("all", "a", false, "Show hidden files")
	listCmd.Flags().BoolP("long", "l", false, "Long format")
	listCmd.Flags().BoolP("recursive", "r", false, "List recursively")

	// Create command flags
	createCmd.Flags().BoolP("force", "f", false, "Force creation, overwrite existing")
	createCmd.Flags().BoolP("recursive", "r", false, "Create parent directories as needed")
	createCmd.Flags().StringP("target", "t", "", "Target path for symbolic links")
	createCmd.Flags().StringP("source", "s", "", "Source directory for archives")
}
