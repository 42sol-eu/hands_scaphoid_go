package fsobj

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Archive represents an archive file (tar, tar.gz, etc.)
type Archive struct {
	path        string
	archiveType string
	info        os.FileInfo
}

// NewArchive creates a new Archive instance
func NewArchive(path string) *Archive {
	info, _ := os.Stat(path)
	archiveType := detectArchiveType(path)
	return &Archive{
		path:        path,
		archiveType: archiveType,
		info:        info,
	}
}

func detectArchiveType(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".gz":
		if strings.HasSuffix(strings.ToLower(path), ".tar.gz") {
			return "tar.gz"
		}
		return "gz"
	case ".tar":
		return "tar"
	case ".zip":
		return "zip"
	default:
		return "unknown"
	}
}

func (a *Archive) Type() ObjectType { return TypeArchive }
func (a *Archive) Path() string     { return a.path }
func (a *Archive) Name() string     { return filepath.Base(a.path) }

func (a *Archive) Size() int64 {
	if a.info == nil {
		return 0
	}
	return a.info.Size()
}

func (a *Archive) ModTime() time.Time {
	if a.info == nil {
		return time.Time{}
	}
	return a.info.ModTime()
}

func (a *Archive) Exists() bool {
	_, err := os.Stat(a.path)
	return !os.IsNotExist(err)
}

func (a *Archive) IsReadable() bool {
	info, err := os.Stat(a.path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0400 != 0
}

func (a *Archive) IsWritable() bool {
	info, err := os.Stat(a.path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0200 != 0
}

func (a *Archive) IsExecutable() bool {
	info, err := os.Stat(a.path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0100 != 0
}

func (a *Archive) CreateFromPath(sourcePath string) error {
	return a.createTarGzArchive(sourcePath)
}

func (a *Archive) Create() error {
	// Create empty archive
	file, err := os.Create(a.path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create minimal tar.gz structure
	gzWriter := gzip.NewWriter(file)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	return nil
}

func (a *Archive) createTarGzArchive(sourcePath string) error {
	file, err := os.Create(a.path)
	if err != nil {
		return err
	}
	defer file.Close()

	gzWriter := gzip.NewWriter(file)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	return filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(sourcePath, path)
		if err != nil {
			return err
		}
		header.Name = relPath

		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(tarWriter, file)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (a *Archive) List() ([]FSObject, error) {
	if a.archiveType == "tar.gz" || a.archiveType == "tar" {
		return a.listTarArchive()
	}
	return nil, fmt.Errorf("archive type %s not supported for listing", a.archiveType)
}

func (a *Archive) listTarArchive() ([]FSObject, error) {
	file, err := os.Open(a.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reader io.Reader = file
	if a.archiveType == "tar.gz" {
		gzReader, err := gzip.NewReader(file)
		if err != nil {
			return nil, err
		}
		defer gzReader.Close()
		reader = gzReader
	}

	tarReader := tar.NewReader(reader)
	var objects []FSObject

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Create virtual file system objects for archive entries
		if header.FileInfo().IsDir() {
			objects = append(objects, NewDirectory(header.Name))
		} else {
			objects = append(objects, NewFile(header.Name))
		}
	}

	return objects, nil
}

func (a *Archive) Copy(destination string) error {
	sourceFile, err := os.Open(a.path)
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
	return err
}

func (a *Archive) Move(destination string) error {
	return os.Rename(a.path, destination)
}

func (a *Archive) Delete() error {
	return os.Remove(a.path)
}

func (a *Archive) Compare(other FSObject) (bool, error) {
	if other.Type() != TypeArchive {
		return false, fmt.Errorf("cannot compare archive with %s", other.Type())
	}

	return a.Size() == other.Size(), nil
}

func (a *Archive) Backup(destination string) error {
	backupPath := filepath.Join(destination, fmt.Sprintf("%s.backup.%d", a.Name(), time.Now().Unix()))
	return a.Copy(backupPath)
}
