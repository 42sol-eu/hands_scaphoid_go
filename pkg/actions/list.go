package actions

import (
	"fmt"
	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
	"strings"
)

// ListAction handles listing of filesystem objects
type ListAction struct {
	ShowHidden  bool
	LongFormat  bool
	Recursive   bool
	SortBy      string
	ReverseSort bool
}

// Execute lists the contents of the specified filesystem object
func (a *ListAction) Execute(path string) *ActionResult {
	obj, err := fsobj.NewFSObject(path)
	if err != nil {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("failed to create object for path %s: %w", path, err),
		}
	}

	if !obj.Exists() {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("path does not exist: %s", path),
		}
	}

	// Check if object supports listing
	lister, ok := obj.(fsobj.Listable)
	if !ok {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("object type %s does not support listing", obj.Type()),
		}
	}

	contents, err := lister.List()
	if err != nil {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("failed to list contents of %s: %w", path, err),
		}
	}

	// Filter hidden files if not requested
	if !a.ShowHidden {
		contents = a.filterHidden(contents)
	}

	// Sort contents if requested
	if a.SortBy != "" {
		contents = a.sortContents(contents)
	}

	// Format output
	var output []string
	for _, item := range contents {
		if a.LongFormat {
			output = append(output, a.formatLong(item))
		} else {
			output = append(output, item.Name())
		}

		// Handle recursive listing for directories
		if a.Recursive && item.Type() == fsobj.TypeDirectory {
			subResult := a.Execute(item.Path())
			if subResult.Success {
				if subOutput, ok := subResult.Data.([]string); ok {
					for _, subItem := range subOutput {
						output = append(output, fmt.Sprintf("  %s", subItem))
					}
				}
			}
		}
	}

	return &ActionResult{
		Success: true,
		Message: fmt.Sprintf("Listed %d items from %s", len(contents), path),
		Data:    output,
	}
}

func (a *ListAction) filterHidden(contents []fsobj.FSObject) []fsobj.FSObject {
	var filtered []fsobj.FSObject
	for _, item := range contents {
		if !strings.HasPrefix(item.Name(), ".") {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func (a *ListAction) sortContents(contents []fsobj.FSObject) []fsobj.FSObject {
	// Simple sorting implementation
	// In a real implementation, you'd use sort.Slice with proper comparison functions
	return contents
}

func (a *ListAction) formatLong(obj fsobj.FSObject) string {
	permissions := a.getPermissionsString(obj)
	size := obj.Size()
	modTime := obj.ModTime().Format("Jan 02 15:04")
	name := obj.Name()

	typeIndicator := ""
	switch obj.Type() {
	case fsobj.TypeDirectory:
		typeIndicator = "d"
	case fsobj.TypeLink:
		typeIndicator = "l"
		if link, ok := obj.(*fsobj.Link); ok {
			name = fmt.Sprintf("%s -> %s", name, link.Target())
		}
	case fsobj.TypeArchive:
		typeIndicator = "a"
	default:
		typeIndicator = "-"
	}

	return fmt.Sprintf("%s%s %8d %s %s", typeIndicator, permissions, size, modTime, name)
}

func (a *ListAction) getPermissionsString(obj fsobj.FSObject) string {
	var perms strings.Builder

	if obj.IsReadable() {
		perms.WriteString("r")
	} else {
		perms.WriteString("-")
	}

	if obj.IsWritable() {
		perms.WriteString("w")
	} else {
		perms.WriteString("-")
	}

	if obj.IsExecutable() {
		perms.WriteString("x")
	} else {
		perms.WriteString("-")
	}

	// Simplified - in real implementation, you'd handle group and other permissions
	return perms.String() + "------"
}
