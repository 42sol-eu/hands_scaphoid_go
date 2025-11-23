package actions

import (
	"github.com/42sol-eu/hands_scaphoid_go/pkg/fsobj"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateAction_Directory(t *testing.T) {
	tempDir := t.TempDir()
	testPath := filepath.Join(tempDir, "testdir")

	action := &CreateAction{Force: false, Recursive: true}
	result := action.Execute(fsobj.TypeDirectory, testPath, nil)

	if !result.Success {
		t.Fatalf("Expected success, got error: %v", result.Error)
	}

	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		t.Error("Directory should have been created")
	}
}

func TestCreateAction_File(t *testing.T) {
	tempDir := t.TempDir()
	testPath := filepath.Join(tempDir, "testfile.txt")

	action := &CreateAction{Force: false}
	result := action.Execute(fsobj.TypeFile, testPath, nil)

	if !result.Success {
		t.Fatalf("Expected success, got error: %v", result.Error)
	}

	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		t.Error("File should have been created")
	}
}

func TestCreateAction_Link(t *testing.T) {
	tempDir := t.TempDir()

	// Create target file
	targetPath := filepath.Join(tempDir, "target.txt")
	os.WriteFile(targetPath, []byte("target content"), 0644)

	linkPath := filepath.Join(tempDir, "testlink")

	action := &CreateAction{Force: false}
	options := map[string]interface{}{"target": targetPath}
	result := action.Execute(fsobj.TypeLink, linkPath, options)

	if !result.Success {
		t.Fatalf("Expected success, got error: %v", result.Error)
	}

	// Check if link exists and points to correct target
	linkInfo, err := os.Lstat(linkPath)
	if err != nil {
		t.Fatalf("Link should exist: %v", err)
	}

	if linkInfo.Mode()&os.ModeSymlink == 0 {
		t.Error("Should be a symbolic link")
	}

	actualTarget, err := os.Readlink(linkPath)
	if err != nil {
		t.Fatalf("Failed to read link: %v", err)
	}

	if actualTarget != targetPath {
		t.Errorf("Expected target %s, got %s", targetPath, actualTarget)
	}
}

func TestCreateAction_Archive(t *testing.T) {
	tempDir := t.TempDir()

	// Create source directory
	sourceDir := filepath.Join(tempDir, "source")
	os.Mkdir(sourceDir, 0755)
	os.WriteFile(filepath.Join(sourceDir, "file1.txt"), []byte("content1"), 0644)
	os.WriteFile(filepath.Join(sourceDir, "file2.txt"), []byte("content2"), 0644)

	archivePath := filepath.Join(tempDir, "test.tar.gz")

	action := &CreateAction{Force: false}
	options := map[string]interface{}{"source": sourceDir}
	result := action.Execute(fsobj.TypeArchive, archivePath, options)

	if !result.Success {
		t.Fatalf("Expected success, got error: %v", result.Error)
	}

	if _, err := os.Stat(archivePath); os.IsNotExist(err) {
		t.Error("Archive should have been created")
	}
}

func TestCreateAction_ForceOverwrite(t *testing.T) {
	tempDir := t.TempDir()
	testPath := filepath.Join(tempDir, "testfile.txt")

	// Create existing file
	os.WriteFile(testPath, []byte("existing"), 0644)

	// Try without force - should fail
	action := &CreateAction{Force: false}
	result := action.Execute(fsobj.TypeFile, testPath, nil)

	if result.Success {
		t.Error("Should have failed without force flag")
	}

	// Try with force - should succeed
	action.Force = true
	result = action.Execute(fsobj.TypeFile, testPath, nil)

	if !result.Success {
		t.Fatalf("Should have succeeded with force flag: %v", result.Error)
	}
}

func TestListAction_Directory(t *testing.T) {
	tempDir := t.TempDir()

	// Create test files
	os.WriteFile(filepath.Join(tempDir, "file1.txt"), []byte("content1"), 0644)
	os.WriteFile(filepath.Join(tempDir, "file2.txt"), []byte("content2"), 0644)
	os.Mkdir(filepath.Join(tempDir, "subdir"), 0755)

	action := &ListAction{ShowHidden: false, LongFormat: false}
	result := action.Execute(tempDir)

	if !result.Success {
		t.Fatalf("Expected success, got error: %v", result.Error)
	}

	output, ok := result.Data.([]string)
	if !ok {
		t.Fatal("Expected string slice output")
	}

	if len(output) != 3 {
		t.Errorf("Expected 3 items, got %d", len(output))
	}
}

func TestListAction_LongFormat(t *testing.T) {
	tempDir := t.TempDir()

	// Create test file
	testFile := filepath.Join(tempDir, "testfile.txt")
	os.WriteFile(testFile, []byte("content"), 0644)

	action := &ListAction{ShowHidden: false, LongFormat: true}
	result := action.Execute(tempDir)

	if !result.Success {
		t.Fatalf("Expected success, got error: %v", result.Error)
	}

	output, ok := result.Data.([]string)
	if !ok {
		t.Fatal("Expected string slice output")
	}

	// Long format should contain more information
	if len(output[0]) <= 20 { // Arbitrary check for longer output
		t.Error("Long format should provide more detailed output")
	}
}

func TestCopyAction(t *testing.T) {
	tempDir := t.TempDir()

	// Create source file
	sourcePath := filepath.Join(tempDir, "source.txt")
	os.WriteFile(sourcePath, []byte("test content"), 0644)

	destPath := filepath.Join(tempDir, "dest.txt")

	action := &CopyAction{Force: false}
	result := action.Execute(sourcePath, destPath)

	if !result.Success {
		t.Fatalf("Expected success, got error: %v", result.Error)
	}

	// Check destination exists
	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		t.Error("Destination file should exist")
	}

	// Check content
	content, err := os.ReadFile(destPath)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if string(content) != "test content" {
		t.Errorf("Expected 'test content', got '%s'", string(content))
	}
}

func TestMoveAction(t *testing.T) {
	tempDir := t.TempDir()

	// Create source file
	sourcePath := filepath.Join(tempDir, "source.txt")
	os.WriteFile(sourcePath, []byte("test content"), 0644)

	destPath := filepath.Join(tempDir, "dest.txt")

	action := &MoveAction{Force: false}
	result := action.Execute(sourcePath, destPath)

	if !result.Success {
		t.Fatalf("Expected success, got error: %v", result.Error)
	}

	// Check source no longer exists
	if _, err := os.Stat(sourcePath); !os.IsNotExist(err) {
		t.Error("Source file should not exist after move")
	}

	// Check destination exists
	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		t.Error("Destination file should exist")
	}
}

func TestDeleteAction(t *testing.T) {
	tempDir := t.TempDir()

	// Create test file
	testPath := filepath.Join(tempDir, "testfile.txt")
	os.WriteFile(testPath, []byte("content"), 0644)

	action := &DeleteAction{Force: true}
	result := action.Execute(testPath)

	if !result.Success {
		t.Fatalf("Expected success, got error: %v", result.Error)
	}

	// Check file no longer exists
	if _, err := os.Stat(testPath); !os.IsNotExist(err) {
		t.Error("File should not exist after deletion")
	}
}

func TestInfoAction(t *testing.T) {
	tempDir := t.TempDir()

	// Create test file
	testPath := filepath.Join(tempDir, "testfile.txt")
	content := []byte("test content")
	os.WriteFile(testPath, content, 0644)

	action := &InfoAction{ShowAll: true}
	result := action.Execute(testPath)

	if !result.Success {
		t.Fatalf("Expected success, got error: %v", result.Error)
	}

	info, ok := result.Data.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map output")
	}

	// Check required fields
	if info["path"] != testPath {
		t.Errorf("Expected path %s, got %v", testPath, info["path"])
	}

	if info["type"] != fsobj.TypeFile {
		t.Errorf("Expected type %s, got %v", fsobj.TypeFile, info["type"])
	}

	if info["size"] != int64(len(content)) {
		t.Errorf("Expected size %d, got %v", len(content), info["size"])
	}
}

func TestCompareAction(t *testing.T) {
	tempDir := t.TempDir()

	// Create two identical files
	file1Path := filepath.Join(tempDir, "file1.txt")
	file2Path := filepath.Join(tempDir, "file2.txt")
	content := []byte("identical content")

	os.WriteFile(file1Path, content, 0644)
	os.WriteFile(file2Path, content, 0644)

	action := &CompareAction{Deep: false}
	result := action.Execute(file1Path, file2Path)

	if !result.Success {
		t.Fatalf("Expected success, got error: %v", result.Error)
	}

	comparison, ok := result.Data.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map output")
	}

	if comparison["equal"] != true {
		t.Error("Files should be equal")
	}
}

func TestBackupAction(t *testing.T) {
	tempDir := t.TempDir()

	// Create source file
	sourcePath := filepath.Join(tempDir, "source.txt")
	os.WriteFile(sourcePath, []byte("backup this"), 0644)

	// Create backup directory
	backupDir := filepath.Join(tempDir, "backups")
	os.Mkdir(backupDir, 0755)

	action := &BackupAction{Timestamp: true}
	result := action.Execute(sourcePath, backupDir)

	if !result.Success {
		t.Fatalf("Expected success, got error: %v", result.Error)
	}

	// Check that backup was created
	backupEntries, err := os.ReadDir(backupDir)
	if err != nil {
		t.Fatalf("Failed to read backup directory: %v", err)
	}

	if len(backupEntries) != 1 {
		t.Errorf("Expected 1 backup, got %d", len(backupEntries))
	}
}
