package actions

import (
	"fmt"
	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
)

// DeleteAction handles deletion of filesystem objects
type DeleteAction struct {
	Force     bool
	Recursive bool
}

// Execute deletes the specified filesystem object
func (a *DeleteAction) Execute(path string) *ActionResult {
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

	// Check if object is a directory and recursive flag is needed
	if obj.Type() == fsobj.TypeDirectory && !a.Recursive {
		// Check if directory is empty
		if lister, ok := obj.(fsobj.Listable); ok {
			contents, err := lister.List()
			if err != nil {
				return &ActionResult{
					Success: false,
					Error:   fmt.Errorf("failed to check directory contents: %w", err),
				}
			}
			if len(contents) > 0 {
				return &ActionResult{
					Success: false,
					Error:   fmt.Errorf("directory is not empty: %s (use --recursive to delete)", path),
				}
			}
		}
	}

	// Check if object supports deletion
	deleter, ok := obj.(fsobj.Deletable)
	if !ok {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("object type %s does not support deletion", obj.Type()),
		}
	}

	// Perform the deletion
	err = deleter.Delete()
	if err != nil {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("failed to delete %s: %w", path, err),
		}
	}

	return &ActionResult{
		Success: true,
		Message: fmt.Sprintf("Successfully deleted %s", path),
		Data:    map[string]interface{}{"path": path, "type": obj.Type()},
	}
}

// InfoAction handles getting information about filesystem objects
type InfoAction struct {
	ShowAll bool
}

// Execute retrieves information about the specified filesystem object
func (a *InfoAction) Execute(path string) *ActionResult {
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

	info := map[string]interface{}{
		"path":       obj.Path(),
		"name":       obj.Name(),
		"type":       obj.Type(),
		"size":       obj.Size(),
		"modTime":    obj.ModTime(),
		"exists":     obj.Exists(),
		"readable":   obj.IsReadable(),
		"writable":   obj.IsWritable(),
		"executable": obj.IsExecutable(),
	}

	// Add type-specific information
	switch obj.Type() {
	case fsobj.TypeLink:
		if link, ok := obj.(*fsobj.Link); ok {
			info["target"] = link.Target()
		}
	case fsobj.TypeDirectory:
		if lister, ok := obj.(fsobj.Listable); ok {
			contents, err := lister.List()
			if err == nil {
				info["itemCount"] = len(contents)
			}
		}
	}

	return &ActionResult{
		Success: true,
		Message: fmt.Sprintf("Information for %s", path),
		Data:    info,
	}
}

// CompareAction handles comparison of filesystem objects
type CompareAction struct {
	Deep bool
}

// Execute compares two filesystem objects
func (a *CompareAction) Execute(path1, path2 string) *ActionResult {
	obj1, err := fsobj.NewFSObject(path1)
	if err != nil {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("failed to create object for path %s: %w", path1, err),
		}
	}

	obj2, err := fsobj.NewFSObject(path2)
	if err != nil {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("failed to create object for path %s: %w", path2, err),
		}
	}

	if !obj1.Exists() {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("first path does not exist: %s", path1),
		}
	}

	if !obj2.Exists() {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("second path does not exist: %s", path2),
		}
	}

	// Check if first object supports comparison
	comparer, ok := obj1.(fsobj.Comparable)
	if !ok {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("object type %s does not support comparison", obj1.Type()),
		}
	}

	// Perform the comparison
	isEqual, err := comparer.Compare(obj2)
	if err != nil {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("failed to compare %s with %s: %w", path1, path2, err),
		}
	}

	result := map[string]interface{}{
		"path1": path1,
		"path2": path2,
		"equal": isEqual,
		"type1": obj1.Type(),
		"type2": obj2.Type(),
		"size1": obj1.Size(),
		"size2": obj2.Size(),
	}

	status := "different"
	if isEqual {
		status = "equal"
	}

	return &ActionResult{
		Success: true,
		Message: fmt.Sprintf("Objects are %s", status),
		Data:    result,
	}
}

// BackupAction handles backup of filesystem objects
type BackupAction struct {
	Timestamp bool
	Compress  bool
}

// Execute creates a backup of the specified filesystem object
func (a *BackupAction) Execute(sourcePath, backupDir string) *ActionResult {
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

	// Check if object supports backup
	backupper, ok := sourceObj.(fsobj.Backupable)
	if !ok {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("object type %s does not support backup", sourceObj.Type()),
		}
	}

	// Perform the backup
	err = backupper.Backup(backupDir)
	if err != nil {
		return &ActionResult{
			Success: false,
			Error:   fmt.Errorf("failed to backup %s to %s: %w", sourcePath, backupDir, err),
		}
	}

	return &ActionResult{
		Success: true,
		Message: fmt.Sprintf("Successfully backed up %s to %s", sourcePath, backupDir),
		Data:    map[string]string{"source": sourcePath, "backupDir": backupDir},
	}
}
