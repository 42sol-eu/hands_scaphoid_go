package actions

import (
	"fmt"
	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
)

// CopyAction handles copying of filesystem objects
type CopyAction struct {
	Recursive    bool
	Force        bool
	PreserveMeta bool
}

// Execute copies the source object to the destination
func (a *CopyAction) Execute(sourcePath, destPath string) *ActionResult {
	sourceObj, err := fsobj.NewFSObject(sourcePath)
	if err != nil {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("failed to create source object: %w", err),
		}
	}

	if !sourceObj.Exists() {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("source path does not exist: %s", sourcePath),
		}
	}

	// Check if destination exists and handle accordingly
	destObj, _ := fsobj.NewFSObject(destPath)
	if destObj != nil && destObj.Exists() && !a.Force {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("destination already exists: %s (use --force to overwrite)", destPath),
		}
	}

	// Check if source supports copying
	copier, ok := sourceObj.(fsobj.Copyable)
	if !ok {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("object type %s does not support copying", sourceObj.Type()),
		}
	}

	// Perform the copy operation
	err = copier.Copy(destPath)
	if err != nil {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("failed to copy %s to %s: %w", sourcePath, destPath, err),
		}
	}

	return &ActionResult{
		Success: true,
		Message: fmt.Sprintf("Successfully copied %s to %s", sourcePath, destPath),
		Data:    map[string]string{"source": sourcePath, "destination": destPath},
	}
}

// MoveAction handles moving/renaming of filesystem objects
type MoveAction struct {
	Force bool
}

// Execute moves the source object to the destination
func (a *MoveAction) Execute(sourcePath, destPath string) *ActionResult {
	sourceObj, err := fsobj.NewFSObject(sourcePath)
	if err != nil {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("failed to create source object: %w", err),
		}
	}

	if !sourceObj.Exists() {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("source path does not exist: %s", sourcePath),
		}
	}

	// Check if destination exists and handle accordingly
	destObj, _ := fsobj.NewFSObject(destPath)
	if destObj != nil && destObj.Exists() && !a.Force {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("destination already exists: %s (use --force to overwrite)", destPath),
		}
	}

	// Check if source supports moving
	mover, ok := sourceObj.(fsobj.Movable)
	if !ok {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("object type %s does not support moving", sourceObj.Type()),
		}
	}

	// Perform the move operation
	err = mover.Move(destPath)
	if err != nil {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("failed to move %s to %s: %w", sourcePath, destPath, err),
		}
	}

	return &ActionResult{
		Success: true,
		Message: fmt.Sprintf("Successfully moved %s to %s", sourcePath, destPath),
		Data:    map[string]string{"source": sourcePath, "destination": destPath},
	}
}
