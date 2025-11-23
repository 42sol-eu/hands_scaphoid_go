package cmd_new

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
	"github.com/spf13/cobra"
)

var existsCmd = &cobra.Command{
	Use:   "exists [path]",
	Short: "Check if filesystem objects exist",
	Long:  `Check if files, directories, links, or other filesystem objects exist.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if IsVerbose() {
			fmt.Fprintf(os.Stderr, "[VERBOSE] Checking existence of: %s\n", colorizePath(args[0]))
		}

		obj, err := fsobj.NewFSObject(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, colorizeError(fmt.Sprintf("Error: %v", err)))
			os.Exit(1)
		}

		if obj.Exists() {
			fmt.Println(formatExistsMessage(args[0], true, string(obj.Type())))
			os.Exit(0)
		} else {
			fmt.Println(formatExistsMessage(args[0], false, ""))
			os.Exit(1)
		}
	},
}

var fileExistsCmd = &cobra.Command{
	Use:   "exists [path]",
	Short: "Check if a file exists",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if IsVerbose() {
			fmt.Fprintf(os.Stderr, "[VERBOSE] Checking if file exists: %s\n", colorizePath(args[0]))
		}

		obj, err := fsobj.NewFSObject(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, colorizeError(fmt.Sprintf("Error: %v", err)))
			os.Exit(1)
		}

		if obj.Exists() && obj.Type() == fsobj.TypeFile {
			fmt.Println(formatTypeCheckMessage(args[0], true, "file", "file"))
			os.Exit(0)
		} else if obj.Exists() {
			fmt.Println(formatTypeCheckMessage(args[0], true, "file", string(obj.Type())))
			os.Exit(1)
		} else {
			fmt.Println(formatExistsMessage(args[0], false, ""))
			os.Exit(1)
		}
	},
}

var directoryExistsCmd = &cobra.Command{
	Use:   "exists [path]",
	Short: "Check if a directory exists",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if IsVerbose() {
			fmt.Fprintf(os.Stderr, "[VERBOSE] Checking if directory exists: %s\n", colorizePath(args[0]))
		}

		obj, err := fsobj.NewFSObject(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, colorizeError(fmt.Sprintf("Error: %v", err)))
			os.Exit(1)
		}

		if obj.Exists() && obj.Type() == fsobj.TypeDirectory {
			fmt.Println(formatTypeCheckMessage(args[0], true, "directory", "directory"))
			os.Exit(0)
		} else if obj.Exists() {
			fmt.Println(formatTypeCheckMessage(args[0], true, "directory", string(obj.Type())))
			os.Exit(1)
		} else {
			fmt.Println(formatExistsMessage(args[0], false, ""))
			os.Exit(1)
		}
	},
}

var linkExistsCmd = &cobra.Command{
	Use:   "exists [path]",
	Short: "Check if a symbolic link exists",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if IsVerbose() {
			fmt.Fprintf(os.Stderr, "[VERBOSE] Checking if link exists: %s\n", colorizePath(args[0]))
		}

		obj, err := fsobj.NewFSObject(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, colorizeError(fmt.Sprintf("Error: %v", err)))
			os.Exit(1)
		}

		if obj.Exists() && obj.Type() == fsobj.TypeLink {
			fmt.Println(formatTypeCheckMessage(args[0], true, "link", "link"))
			os.Exit(0)
		} else if obj.Exists() {
			fmt.Println(formatTypeCheckMessage(args[0], true, "link", string(obj.Type())))
			os.Exit(1)
		} else {
			fmt.Println(formatExistsMessage(args[0], false, ""))
			os.Exit(1)
		}
	},
}

var pathExistsCmd = &cobra.Command{
	Use:   "exists [path]",
	Short: "Check if a path exists (any type)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if IsVerbose() {
			fmt.Fprintf(os.Stderr, "[VERBOSE] Checking if path exists: %s\n", colorizePath(args[0]))
		}

		obj, err := fsobj.NewFSObject(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, colorizeError(fmt.Sprintf("Error: %v", err)))
			os.Exit(1)
		}

		if obj.Exists() {
			fmt.Println(formatExistsMessage(args[0], true, string(obj.Type())))
			os.Exit(0)
		} else {
			fmt.Println(formatExistsMessage(args[0], false, ""))
			os.Exit(1)
		}
	},
}

func init() {
	// Add exists subcommand to file command
	fileCmd.AddCommand(fileExistsCmd)
	
	// Add exists subcommand to directory command
	directoryCmd.AddCommand(directoryExistsCmd)
	
	// Add exists subcommand to link command
	linkCmd.AddCommand(linkExistsCmd)
	
	// Create a separate path command with exists subcommand
	var pathCmd = &cobra.Command{
		Use:   "path",
		Short: "Path operations",
		Long:  `Check existence and properties of any filesystem path.`,
	}
	pathCmd.AddCommand(pathExistsCmd)
	
	// Add path command to root (will be done in root.go init)
	// We'll need to export pathCmd
}
