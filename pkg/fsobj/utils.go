package fsobj

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// copyRecursive recursively copies source to destination
func copyRecursive(source, destination string) error {
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// Create destination directory
	if err := os.MkdirAll(destination, sourceInfo.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(source, entry.Name())
		destPath := filepath.Join(destination, entry.Name())

		entryInfo, err := entry.Info()
		if err != nil {
			return err
		}

		if entryInfo.IsDir() {
			if err := copyRecursive(sourcePath, destPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(sourcePath, destPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// copyFile copies a single file from source to destination
func copyFile(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return err
	}

	sourceInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	return os.Chmod(destination, sourceInfo.Mode())
}

// Factory function to create appropriate FSObject based on path
func NewFSObject(path string) (FSObject, error) {
	info, err := os.Lstat(path)
	if err != nil {
		// If file doesn't exist, determine type by extension or path
		if filepath.Ext(path) != "" {
			// Has extension, likely a file
			return NewFile(path), nil
		}
		// No extension, could be directory
		return NewDirectory(path), nil
	}

	if info.Mode()&os.ModeSymlink != 0 {
		return NewLink(path), nil
	}

	if info.IsDir() {
		return NewDirectory(path), nil
	}

	// Check if it's an archive by extension
	ext := filepath.Ext(path)
	archiveExts := []string{".tar", ".gz", ".zip", ".tar.gz"}
	for _, archiveExt := range archiveExts {
		if ext == archiveExt || filepath.Base(path) == archiveExt {
			return NewArchive(path), nil
		}
	}

	return NewFile(path), nil
}

// CreateFSObject creates a filesystem object of the specified type
func CreateFSObject(objectType ObjectType, path string) (FSObject, error) {
	switch objectType {
	case TypeDirectory:
		return NewDirectory(path), nil
	case TypeFile:
		return NewFile(path), nil
	case TypeLink:
		return NewLink(path), nil
	case TypeArchive:
		return NewArchive(path), nil
	default:
		return nil, fmt.Errorf("unsupported object type: %s", objectType)
	}
}
