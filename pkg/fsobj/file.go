package fsobj

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// File represents a file in the filesystem
type File struct {
	path string
	info os.FileInfo
}

// NewFile creates a new File instance
func NewFile(path string) *File {
	info, _ := os.Stat(path)
	return &File{
		path: path,
		info: info,
	}
}

func (f *File) Type() ObjectType { return TypeFile }
func (f *File) Path() string     { return f.path }
func (f *File) Name() string     { return filepath.Base(f.path) }

func (f *File) Size() int64 {
	if f.info == nil {
		return 0
	}
	return f.info.Size()
}

func (f *File) ModTime() time.Time {
	if f.info == nil {
		return time.Time{}
	}
	return f.info.ModTime()
}

func (f *File) Exists() bool {
	_, err := os.Stat(f.path)
	return !os.IsNotExist(err)
}

func (f *File) IsReadable() bool {
	info, err := os.Stat(f.path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0400 != 0
}

func (f *File) IsWritable() bool {
	info, err := os.Stat(f.path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0200 != 0
}

func (f *File) IsExecutable() bool {
	info, err := os.Stat(f.path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0100 != 0
}

func (f *File) Create() error {
	file, err := os.Create(f.path)
	if err != nil {
		return err
	}
	return file.Close()
}

func (f *File) List() ([]FSObject, error) {
	// For files, list returns the content as lines or bytes
	// This is a simplified implementation
	return []FSObject{f}, nil
}

func (f *File) Copy(destination string) error {
	sourceFile, err := os.Open(f.path)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// Copy file permissions
	sourceInfo, err := os.Stat(f.path)
	if err != nil {
		return err
	}
	return os.Chmod(destination, sourceInfo.Mode())
}

func (f *File) Move(destination string) error {
	return os.Rename(f.path, destination)
}

func (f *File) Delete() error {
	return os.Remove(f.path)
}

func (f *File) Compare(other FSObject) (bool, error) {
	if other.Type() != TypeFile {
		return false, fmt.Errorf("cannot compare file with %s", other.Type())
	}

	// Compare file sizes first
	if f.Size() != other.Size() {
		return false, nil
	}

	// Compare modification times
	return f.ModTime().Equal(other.ModTime()), nil
}

func (f *File) Backup(destination string) error {
	backupPath := filepath.Join(destination, fmt.Sprintf("%s.backup.%d", f.Name(), time.Now().Unix()))
	return f.Copy(backupPath)
}
