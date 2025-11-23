package cmd_new

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
	"github.com/42sol-eu/hands_scaphoid_go/pkg/naming"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var checkCmd = &cobra.Command{
	Use:   "check [path]",
	Short: "Check naming rules for files and directories",
	Long: `Recursively check all files and directories against naming rules.
Reports violations with severity levels (MUST errors, WANT warnings).

Automatically skips blacklisted directories like .git, node_modules, etc.
See .scaphoid.yaml configuration for the default blacklist.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		recursive, _ := cmd.Flags().GetBool("recursive")
		showValid, _ := cmd.Flags().GetBool("show-valid")
		errorsOnly, _ := cmd.Flags().GetBool("errors-only")
		skipDirs, _ := cmd.Flags().GetStringSlice("skip")

		// Add custom skip directories to blacklist temporarily
		if len(skipDirs) > 0 {
			existing := viper.GetStringSlice("naming_rules.blacklist")
			viper.Set("naming_rules.blacklist", append(existing, skipDirs...))
		}

		// Check if validation is enabled
		if !viper.GetBool("naming_rules.enabled") {
			fmt.Fprintln(os.Stderr, wrapErrorMessage("Naming rules validation is disabled in configuration"))
			os.Exit(1)
		}

		validator := naming.NewValidator()
		stats := &CheckStats{}

		err := checkPath(args[0], validator, recursive, showValid, errorsOnly, stats)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", formatErrorMessage(fmt.Sprintf("Error: %v", err)))
			os.Exit(1)
		}

		// Print summary
		fmt.Println()
		fmt.Println(formatSectionHeader("Summary", ""))
		fmt.Printf("Total objects checked: %d\n", stats.Total)
		fmt.Printf("- Files: %d\n", stats.Files)
		fmt.Printf("- Directories: %d\n", stats.Directories)
		fmt.Printf("- Links: %d\n", stats.Links)
		fmt.Printf("- Archives: %d\n", stats.Archives)
		fmt.Println()

		if stats.Errors > 0 {
			fmt.Printf("%s %d MUST violations (errors)\n", colorizeError("✗"), stats.Errors)
		} else {
			fmt.Printf("%s No MUST violations\n", colorizeSuccess("✓"))
		}

		if stats.Warnings > 0 {
			fmt.Printf("%s %d WANT violations (warnings)\n", colorizeKeyword("⚠"), stats.Warnings)
		} else {
			fmt.Printf("%s No WANT violations\n", colorizeSuccess("✓"))
		}

		// Exit with error code if there are MUST violations
		if stats.Errors > 0 {
			os.Exit(1)
		}
	},
}

// CheckStats tracks statistics for the check command
type CheckStats struct {
	Total       int
	Files       int
	Directories int
	Links       int
	Archives    int
	Errors      int
	Warnings    int
}

func checkPath(path string, validator *naming.Validator, recursive bool, showValid bool, errorsOnly bool, stats *CheckStats) error {
	// Get file info
	info, err := os.Lstat(path)
	if err != nil {
		return fmt.Errorf("cannot access %s: %w", path, err)
	}

	// Determine object type
	var objType fsobj.ObjectType
	if info.Mode()&os.ModeSymlink != 0 {
		objType = fsobj.TypeLink
		stats.Links++
	} else if info.IsDir() {
		objType = fsobj.TypeDirectory
		stats.Directories++
	} else {
		// Check if it's an archive based on extension
		ext := filepath.Ext(path)
		archiveExts := map[string]bool{
			".zip": true, ".tar": true, ".gz": true, ".tgz": true,
			".bz2": true, ".tbz2": true, ".7z": true, ".rar": true,
		}
		if archiveExts[ext] || filepath.Ext(filepath.Base(path[:len(path)-len(ext)])) == ".tar" {
			objType = fsobj.TypeArchive
			stats.Archives++
		} else {
			objType = fsobj.TypeFile
			stats.Files++
		}
	}

	stats.Total++

	// Validate the name
	violations := validator.Validate(path, objType)

	// Separate errors and warnings
	var errors []naming.Violation
	var warnings []naming.Violation

	for _, v := range violations {
		if v.IsError() {
			errors = append(errors, v)
			stats.Errors++
		} else {
			warnings = append(warnings, v)
			stats.Warnings++
		}
	}

	// Print results
	if len(errors) > 0 {
		fmt.Printf("%s %s\n", colorizeError("✗ MUST"), colorizePath(path))
		for _, e := range errors {
			fmt.Printf("  %s\n", colorizeError(e.Message))
		}
	} else if len(warnings) > 0 && !errorsOnly {
		fmt.Printf("%s %s\n", colorizeKeyword("⚠ WANT"), colorizePath(path))
		for _, w := range warnings {
			fmt.Printf("  %s\n", colorizeKeyword(w.Message))
		}
	} else if showValid && len(violations) == 0 {
		fmt.Printf("%s %s\n", colorizeSuccess("✓"), colorizePath(path))
	}

	// Recurse into directories if requested
	if recursive && info.IsDir() {
		// Check if directory is blacklisted
		dirName := filepath.Base(path)
		if isBlacklisted(dirName) {
			if IsVerbose() {
				fmt.Fprintf(os.Stderr, "[VERBOSE] Skipping blacklisted directory: %s\n", path)
			}
			return nil
		}

		entries, err := os.ReadDir(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: cannot read directory %s: %v\n", path, err)
			return nil
		}

		for _, entry := range entries {
			childPath := filepath.Join(path, entry.Name())
			if err := checkPath(childPath, validator, recursive, showValid, errorsOnly, stats); err != nil {
				// Continue on error, just report it
				fmt.Fprintf(os.Stderr, "Warning: %v\n", err)
			}
		}
	}

	return nil
}

// isBlacklisted checks if a directory name is in the blacklist
func isBlacklisted(dirName string) bool {
	blacklist := viper.GetStringSlice("naming_rules.blacklist")
	for _, blocked := range blacklist {
		if dirName == blocked {
			return true
		}
	}
	return false
}

func init() {
	checkCmd.Flags().BoolP("recursive", "r", false, "Check directories recursively")
	checkCmd.Flags().Bool("show-valid", false, "Show valid names (no violations)")
	checkCmd.Flags().BoolP("errors-only", "e", false, "Show only MUST violations (errors)")
	checkCmd.Flags().StringSliceP("skip", "s", []string{}, "Additional directories to skip (comma-separated)")
}
