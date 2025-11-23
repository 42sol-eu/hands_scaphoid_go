package fsobj

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewDirectory(t *testing.T) {
	tempDir := t.TempDir()
	dir := NewDirectory(tempDir)

	if dir.Type() != TypeDirectory {
		t.Errorf("Expected type %s, got %s", TypeDirectory, dir.Type())
	}

	if dir.Path() != tempDir {
		t.Errorf("Expected path %s, got %s", tempDir, dir.Path())
	}

	if !dir.Exists() {
		t.Error("Directory should exist")
	}
}

func TestDirectoryCreate(t *testing.T) {
	tempDir := t.TempDir()
	testPath := filepath.Join(tempDir, "testdir")

	dir := NewDirectory(testPath)

	if dir.Exists() {
		t.Error("Directory should not exist before creation")
	}

	err := dir.Create()
	if err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	if !dir.Exists() {
		t.Error("Directory should exist after creation")
	}

	// Test creating with parent directories
	nestedPath := filepath.Join(tempDir, "parent", "child", "grandchild")
	nestedDir := NewDirectory(nestedPath)

	err = nestedDir.Create()
	if err != nil {
		t.Fatalf("Failed to create nested directory: %v", err)
	}

	if !nestedDir.Exists() {
		t.Error("Nested directory should exist after creation")
	}
}

func TestDirectoryList(t *testing.T) {
	tempDir := t.TempDir()

	// Create some files and subdirectories
	subDir := filepath.Join(tempDir, "subdir")
	os.Mkdir(subDir, 0755)

	testFile := filepath.Join(tempDir, "testfile.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	dir := NewDirectory(tempDir)
	contents, err := dir.List()
	if err != nil {
		t.Fatalf("Failed to list directory: %v", err)
	}

	if len(contents) != 2 {
		t.Errorf("Expected 2 items, got %d", len(contents))
	}

	// Check that we have both a directory and a file
	foundDir := false
	foundFile := false
	for _, item := range contents {
		if item.Type() == TypeDirectory {
			foundDir = true
		}
		if item.Type() == TypeFile {
			foundFile = true
		}
	}

	if !foundDir {
		t.Error("Should have found a directory")
	}
	if !foundFile {
		t.Error("Should have found a file")
	}
}

func TestDirectoryCopy(t *testing.T) {
	tempDir := t.TempDir()

	// Create source directory with content
	sourceDir := filepath.Join(tempDir, "source")
	os.Mkdir(sourceDir, 0755)

	testFile := filepath.Join(sourceDir, "testfile.txt")
	os.WriteFile(testFile, []byte("test content"), 0644)

	subDir := filepath.Join(sourceDir, "subdir")
	os.Mkdir(subDir, 0755)

	// Copy directory
	destDir := filepath.Join(tempDir, "destination")
	source := NewDirectory(sourceDir)

	err := source.Copy(destDir)
	if err != nil {
		t.Fatalf("Failed to copy directory: %v", err)
	}

	// Verify copy
	dest := NewDirectory(destDir)
	if !dest.Exists() {
		t.Error("Destination directory should exist")
	}

	// Check that files were copied
	copiedFile := filepath.Join(destDir, "testfile.txt")
	if _, err := os.Stat(copiedFile); os.IsNotExist(err) {
		t.Error("File should have been copied")
	}

	copiedSubDir := filepath.Join(destDir, "subdir")
	if _, err := os.Stat(copiedSubDir); os.IsNotExist(err) {
		t.Error("Subdirectory should have been copied")
	}
}

func TestDirectoryMove(t *testing.T) {
	tempDir := t.TempDir()

	// Create source directory
	sourceDir := filepath.Join(tempDir, "source")
	os.Mkdir(sourceDir, 0755)

	testFile := filepath.Join(sourceDir, "testfile.txt")
	os.WriteFile(testFile, []byte("test content"), 0644)

	// Move directory
	destDir := filepath.Join(tempDir, "destination")
	source := NewDirectory(sourceDir)

	err := source.Move(destDir)
	if err != nil {
		t.Fatalf("Failed to move directory: %v", err)
	}

	// Verify move
	if _, err := os.Stat(sourceDir); !os.IsNotExist(err) {
		t.Error("Source directory should not exist after move")
	}

	dest := NewDirectory(destDir)
	if !dest.Exists() {
		t.Error("Destination directory should exist")
	}

	// Check that file was moved
	movedFile := filepath.Join(destDir, "testfile.txt")
	if _, err := os.Stat(movedFile); os.IsNotExist(err) {
		t.Error("File should have been moved")
	}
}

func TestDirectoryDelete(t *testing.T) {
	tempDir := t.TempDir()

	// Create directory with content
	testDir := filepath.Join(tempDir, "testdir")
	os.Mkdir(testDir, 0755)

	testFile := filepath.Join(testDir, "testfile.txt")
	os.WriteFile(testFile, []byte("test content"), 0644)

	dir := NewDirectory(testDir)

	err := dir.Delete()
	if err != nil {
		t.Fatalf("Failed to delete directory: %v", err)
	}

	if dir.Exists() {
		t.Error("Directory should not exist after deletion")
	}
}

func TestDirectoryPermissions(t *testing.T) {
	tempDir := t.TempDir()
	dir := NewDirectory(tempDir)

	// Test permissions (these tests might be platform-specific)
	if !dir.IsReadable() {
		t.Error("Directory should be readable")
	}

	if !dir.IsWritable() {
		t.Error("Directory should be writable")
	}

	if !dir.IsExecutable() {
		t.Error("Directory should be executable")
	}
}

func TestDirectoryCompare(t *testing.T) {
	tempDir := t.TempDir()

	// Create two directories
	dir1Path := filepath.Join(tempDir, "dir1")
	dir2Path := filepath.Join(tempDir, "dir2")

	os.Mkdir(dir1Path, 0755)
	os.Mkdir(dir2Path, 0755)

	dir1 := NewDirectory(dir1Path)
	dir2 := NewDirectory(dir2Path)

	// Empty directories should be equal
	isEqual, err := dir1.Compare(dir2)
	if err != nil {
		t.Fatalf("Failed to compare directories: %v", err)
	}

	if !isEqual {
		t.Error("Empty directories should be equal")
	}

	// Add file to one directory
	testFile := filepath.Join(dir1Path, "testfile.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	// Now they should be different
	isEqual, err = dir1.Compare(dir2)
	if err != nil {
		t.Fatalf("Failed to compare directories: %v", err)
	}

	if isEqual {
		t.Error("Directories with different contents should not be equal")
	}
}

func TestDirectoryBackup(t *testing.T) {
	tempDir := t.TempDir()

	// Create source directory
	sourceDir := filepath.Join(tempDir, "source")
	os.Mkdir(sourceDir, 0755)

	testFile := filepath.Join(sourceDir, "testfile.txt")
	os.WriteFile(testFile, []byte("test content"), 0644)

	// Create backup directory
	backupDir := filepath.Join(tempDir, "backups")
	os.Mkdir(backupDir, 0755)

	source := NewDirectory(sourceDir)

	err := source.Backup(backupDir)
	if err != nil {
		t.Fatalf("Failed to backup directory: %v", err)
	}

	// Check that backup was created (it should have a timestamped name)
	backupEntries, err := os.ReadDir(backupDir)
	if err != nil {
		t.Fatalf("Failed to read backup directory: %v", err)
	}

	if len(backupEntries) != 1 {
		t.Errorf("Expected 1 backup, got %d", len(backupEntries))
	}

	// The backup should contain our test file
	backupPath := filepath.Join(backupDir, backupEntries[0].Name())
	backupTestFile := filepath.Join(backupPath, "testfile.txt")
	if _, err := os.Stat(backupTestFile); os.IsNotExist(err) {
		t.Error("Backup should contain the test file")
	}
}
