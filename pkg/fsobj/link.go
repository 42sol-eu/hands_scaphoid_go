package fsobj

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Link represents a symbolic link in the filesystem
type Link struct {
	path   string
	target string
	info   os.FileInfo
}

// NewLink creates a new Link instance
func NewLink(path string) *Link {
	info, _ := os.Lstat(path)
	target, _ := os.Readlink(path)
	return &Link{
		path:   path,
		target: target,
		info:   info,
	}
}

func (l *Link) Type() ObjectType { return TypeLink }
func (l *Link) Path() string     { return l.path }
func (l *Link) Name() string     { return filepath.Base(l.path) }
func (l *Link) Target() string   { return l.target }

func (l *Link) Size() int64 {
	if l.info == nil {
		return 0
	}
	return l.info.Size()
}

func (l *Link) ModTime() time.Time {
	if l.info == nil {
		return time.Time{}
	}
	return l.info.ModTime()
}

func (l *Link) Exists() bool {
	_, err := os.Lstat(l.path)
	return !os.IsNotExist(err)
}

func (l *Link) IsReadable() bool {
	info, err := os.Lstat(l.path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0400 != 0
}

func (l *Link) IsWritable() bool {
	info, err := os.Lstat(l.path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0200 != 0
}

func (l *Link) IsExecutable() bool {
	info, err := os.Lstat(l.path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0100 != 0
}

func (l *Link) CreateWithTarget(target string) error {
	l.target = target
	return os.Symlink(target, l.path)
}

func (l *Link) Create() error {
	if l.target == "" {
		return fmt.Errorf("target not specified for link creation")
	}
	return os.Symlink(l.target, l.path)
}

func (l *Link) List() ([]FSObject, error) {
	// Follow the link and list the target
	if l.target == "" {
		return nil, fmt.Errorf("link target not found")
	}

	targetInfo, err := os.Stat(l.target)
	if err != nil {
		return nil, err
	}

	if targetInfo.IsDir() {
		dir := NewDirectory(l.target)
		return dir.List()
	}

	file := NewFile(l.target)
	return []FSObject{file}, nil
}

func (l *Link) Copy(destination string) error {
	return os.Symlink(l.target, destination)
}

func (l *Link) Move(destination string) error {
	err := l.Copy(destination)
	if err != nil {
		return err
	}
	return l.Delete()
}

func (l *Link) Delete() error {
	return os.Remove(l.path)
}

func (l *Link) Compare(other FSObject) (bool, error) {
	if other.Type() != TypeLink {
		return false, fmt.Errorf("cannot compare link with %s", other.Type())
	}

	otherLink := other.(*Link)
	return l.target == otherLink.target, nil
}

func (l *Link) Backup(destination string) error {
	backupPath := filepath.Join(destination, fmt.Sprintf("%s.backup.%d", l.Name(), time.Now().Unix()))
	return l.Copy(backupPath)
}
