package fsobj

import (
	"time"
)

// ObjectType represents the type of filesystem object
type ObjectType string

const (
	TypeDirectory ObjectType = "directory"
	TypeFile      ObjectType = "file"
	TypePath      ObjectType = "path"
	TypeVariable  ObjectType = "variable"
	TypeLink      ObjectType = "link"
	TypeArchive   ObjectType = "archive"
)

// FSObject represents a filesystem object with common properties
type FSObject interface {
	Type() ObjectType
	Path() string
	Name() string
	Size() int64
	ModTime() time.Time
	Exists() bool
	IsReadable() bool
	IsWritable() bool
	IsExecutable() bool
}

// Creatable objects that can be created
type Creatable interface {
	Create() error
}

// Listable objects that can list their contents
type Listable interface {
	List() ([]FSObject, error)
}

// Copyable objects that can be copied
type Copyable interface {
	Copy(destination string) error
}

// Movable objects that can be moved/renamed
type Movable interface {
	Move(destination string) error
}

// Deletable objects that can be deleted
type Deletable interface {
	Delete() error
}

// Comparable objects that can be compared with others
type Comparable interface {
	Compare(other FSObject) (bool, error)
}

// Backupable objects that can be backed up
type Backupable interface {
	Backup(destination string) error
}
