package main

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/actions"
	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
	"github.com/spf13/cobra"
)

var (
	force     bool
	source    string
	showAll   bool
	deep      bool
	timestamp bool
	compress  bool
)

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

var rootCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive operations",
	Long:  `Perform various operations on archive files.`,
}

var createCmd = &cobra.Command{
	Use:   "create [archive-path]",
	Short: "Create an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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

var listCmd = &cobra.Command{
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

var copyCmd = &cobra.Command{
	Use:   "copy [source] [destination]",
	Short: "Copy an archive",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
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

var moveCmd = &cobra.Command{
	Use:   "move [source] [destination]",
	Short: "Move an archive",
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

var deleteCmd = &cobra.Command{
	Use:   "delete [archive-path]",
	Short: "Delete an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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

var infoCmd = &cobra.Command{
	Use:   "info [path]",
	Short: "Get archive information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateArchivePath(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

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
	Short: "Backup an archive",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateArchivePath(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

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
	rootCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "Force operation")

	createCmd.Flags().StringVarP(&source, "source", "s", "", "Source directory to archive")
	infoCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all available information")
	compareCmd.Flags().BoolVarP(&deep, "deep", "d", false, "Perform deep comparison")
	backupCmd.Flags().BoolVarP(&timestamp, "timestamp", "t", true, "Add timestamp to backup name")
	backupCmd.Flags().BoolVarP(&compress, "compress", "c", false, "Compress backup")

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(copyCmd)
	rootCmd.AddCommand(moveCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(compareCmd)
	rootCmd.AddCommand(backupCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
