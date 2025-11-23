package fsobj

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Directory represents a directory in the filesystem
type Directory struct {
	path string
	info os.FileInfo
}

// NewDirectory creates a new Directory instance
func NewDirectory(path string) *Directory {
	info, _ := os.Stat(path)
	return &Directory{
		path: path,
		info: info,
	}
}

func (d *Directory) Type() ObjectType { return TypeDirectory }
func (d *Directory) Path() string     { return d.path }
func (d *Directory) Name() string     { return filepath.Base(d.path) }

func (d *Directory) Size() int64 {
	if d.info == nil {
		return 0
	}
	return d.info.Size()
}

func (d *Directory) ModTime() time.Time {
	if d.info == nil {
		return time.Time{}
	}
	return d.info.ModTime()
}

func (d *Directory) Exists() bool {
	_, err := os.Stat(d.path)
	return !os.IsNotExist(err)
}

func (d *Directory) IsReadable() bool {
	info, err := os.Stat(d.path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0400 != 0
}

func (d *Directory) IsWritable() bool {
	info, err := os.Stat(d.path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0200 != 0
}

func (d *Directory) IsExecutable() bool {
	info, err := os.Stat(d.path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0100 != 0
}

func (d *Directory) Create() error {
	return os.MkdirAll(d.path, 0755)
}

func (d *Directory) List() ([]FSObject, error) {
	entries, err := os.ReadDir(d.path)
	if err != nil {
		return nil, err
	}

	var objects []FSObject
	for _, entry := range entries {
		fullPath := filepath.Join(d.path, entry.Name())
		if entry.IsDir() {
			objects = append(objects, NewDirectory(fullPath))
		} else {
			objects = append(objects, NewFile(fullPath))
		}
	}
	return objects, nil
}

func (d *Directory) Copy(destination string) error {
	return copyRecursive(d.path, destination)
}

func (d *Directory) Move(destination string) error {
	return os.Rename(d.path, destination)
}

func (d *Directory) Delete() error {
	return os.RemoveAll(d.path)
}

func (d *Directory) Compare(other FSObject) (bool, error) {
	if other.Type() != TypeDirectory {
		return false, fmt.Errorf("cannot compare directory with %s", other.Type())
	}

	// Compare by checking if both paths exist and have similar structure
	otherDir := other.(*Directory)
	thisEntries, err := d.List()
	if err != nil {
		return false, err
	}

	otherEntries, err := otherDir.List()
	if err != nil {
		return false, err
	}

	return len(thisEntries) == len(otherEntries), nil
}

func (d *Directory) Backup(destination string) error {
	backupPath := filepath.Join(destination, fmt.Sprintf("%s.backup.%d", d.Name(), time.Now().Unix()))
	return d.Copy(backupPath)
}
