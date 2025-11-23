package cmd_new

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "scaphoid",
	Short: "A unified file system operations CLI tool",
	Long: `Scaphoid is a comprehensive file system operations tool that provides
commands for managing files, directories, links, and archives.

It offers subcommands for creating, listing, copying, moving, deleting,
comparing, and backing up various file system objects.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Fprintf(os.Stderr, "[VERBOSE] Command: %s\n", cmd.CommandPath())
			fmt.Fprintf(os.Stderr, "[VERBOSE] Args: %v\n", args)
		}
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

// IsVerbose returns whether verbose mode is enabled
func IsVerbose() bool {
	return verbose
}

func init() {
	// Add global persistent flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	// Add all subcommands to root
	rootCmd.AddCommand(fileCmd)
	rootCmd.AddCommand(directoryCmd)
	rootCmd.AddCommand(linkCmd)
	rootCmd.AddCommand(archiveCmd)
	rootCmd.AddCommand(copyCmd)
	rootCmd.AddCommand(moveCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(compareCmd)
	rootCmd.AddCommand(backupCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(existsCmd)
	
	// Add formatting commands
	rootCmd.AddCommand(sectionCmd)
	rootCmd.AddCommand(separatorCmd)
	
	// Add interactive command
	rootCmd.AddCommand(interactiveCmd)
}
