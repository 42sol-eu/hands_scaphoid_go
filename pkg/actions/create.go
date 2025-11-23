package actions

import (
	"fmt"
	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
)

// ActionResult represents the result of an action
type ActionResult struct {
	Success bool
	Message string
	Data    interface{}
	Error   error
}

// CreateAction handles creation of filesystem objects
type CreateAction struct {
	Force     bool
	Recursive bool
	Mode      string
}

// Execute creates the specified filesystem object
func (a *CreateAction) Execute(objType fsobj.ObjectType, path string, options map[string]interface{}) *ActionResult {
	obj, err := fsobj.CreateFSObject(objType, path)
	if err != nil {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("failed to create object: %w", err),
		}
	}

	// Check if object already exists
	if obj.Exists() && !a.Force {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("object already exists at %s (use --force to overwrite)", path),
		}
	}

	// Handle special cases for different object types
	switch objType {
	case fsobj.TypeLink:
		link := obj.(*fsobj.Link)
		if target, ok := options["target"]; ok {
			if targetStr, ok := target.(string); ok {
				err = link.CreateWithTarget(targetStr)
			} else {
				return &ActionResult{
					Success: false,
					Error:   fmt.Errorf("invalid target type for link creation"),
				}
			}
		} else {
			return &ActionResult{
				Success: false,
				Error:   fmt.Errorf("target is required for link creation"),
			}
		}
	case fsobj.TypeArchive:
		archive := obj.(*fsobj.Archive)
		if source, ok := options["source"]; ok {
			if sourceStr, ok := source.(string); ok {
				err = archive.CreateFromPath(sourceStr)
			} else {
				return &ActionResult{
					Success: false,
					Error:   fmt.Errorf("invalid source type for archive creation"),
				}
			}
		} else {
			err = archive.Create()
		}
	default:
		// Standard creation for files and directories
		if creator, ok := obj.(fsobj.Creatable); ok {
			err = creator.Create()
		} else {
			return &ActionResult{
				Success: false,
				Error:   fmt.Errorf("object type %s does not support creation", objType),
			}
		}
	}

	if err != nil {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("failed to create %s at %s: %w", objType, path, err),
		}
	}

	return &ActionResult{
		Success: true,
		Message: fmt.Sprintf("Successfully created %s at %s", objType, path),
		Data:    obj,
	}
}
