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
	recursive bool
	target    string
	source    string
)

var rootCmd = &cobra.Command{
	Use:   "create",
	Short: "Create filesystem objects",
	Long:  `Create various types of filesystem objects like files, directories, links, and archives.`,
}

var directoryCmd = &cobra.Command{
	Use:   "directory [path]",
	Short: "Create a directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := &actions.CreateAction{
			Force:     force,
			Recursive: recursive,
		}

		result := action.Execute(fsobj.TypeDirectory, args[0], nil)
		if result.Success {
			fmt.Println(result.Message)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

var fileCmd = &cobra.Command{
	Use:   "file [path]",
	Short: "Create a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := &actions.CreateAction{
			Force:     force,
			Recursive: recursive,
		}

		result := action.Execute(fsobj.TypeFile, args[0], nil)
		if result.Success {
			fmt.Println(result.Message)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	},
}

var linkCmd = &cobra.Command{
	Use:   "link [path]",
	Short: "Create a symbolic link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if target == "" {
			fmt.Fprintf(os.Stderr, "Error: --target is required for link creation\n")
			os.Exit(1)
		}

		action := &actions.CreateAction{
			Force:     force,
			Recursive: recursive,
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

var archiveCmd = &cobra.Command{
	Use:   "archive [path]",
	Short: "Create an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := &actions.CreateAction{
			Force:     force,
			Recursive: recursive,
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

func init() {
	rootCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "Force creation, overwrite existing")
	rootCmd.PersistentFlags().BoolVarP(&recursive, "recursive", "r", false, "Create parent directories if needed")

	linkCmd.Flags().StringVarP(&target, "target", "t", "", "Target path for the symbolic link")
	archiveCmd.Flags().StringVarP(&source, "source", "s", "", "Source directory to archive")

	rootCmd.AddCommand(directoryCmd)
	rootCmd.AddCommand(fileCmd)
	rootCmd.AddCommand(linkCmd)
	rootCmd.AddCommand(archiveCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
