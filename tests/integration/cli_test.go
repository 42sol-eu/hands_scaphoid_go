//go:build integration

package integration

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestCommandsBuilt checks that all expected commands can be built
func TestCommandsBuilt(t *testing.T) {
	actionCommands := []string{"create", "list", "copy", "move", "delete", "info", "compare", "backup"}
	objectCommands := []string{"directory", "file", "link", "archive"}
	
	// Build all commands
	for _, cmd := range actionCommands {
		t.Run("build_"+cmd, func(t *testing.T) {
			buildCmd := exec.Command("go", "build", "-o", "/tmp/fs-test-"+cmd, "./cmd/"+cmd)
			if err := buildCmd.Run(); err != nil {
				t.Fatalf("Failed to build %s: %v", cmd, err)
			}
		})
	}
	
	for _, cmd := range objectCommands {
		t.Run("build_"+cmd, func(t *testing.T) {
			buildCmd := exec.Command("go", "build", "-o", "/tmp/fs-test-"+cmd, "./cmd/"+cmd)
			if err := buildCmd.Run(); err != nil {
				t.Fatalf("Failed to build %s: %v", cmd, err)
			}
		})
	}
}

// TestCreateDirectoryCommand tests the create command for directories
func TestCreateDirectoryCommand(t *testing.T) {
	tempDir := t.TempDir()
	testPath := filepath.Join(tempDir, "testdir")
	
	// Build create command
	buildCmd := exec.Command("go", "build", "-o", "/tmp/fs-test-create", "./cmd/create")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build create command: %v", err)
	}
	
	// Run create directory command
	cmd := exec.Command("/tmp/fs-test-create", "directory", testPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Create command failed: %v, output: %s", err, output)
	}
	
	// Check that directory was created
	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		t.Errorf("Directory should have been created at %s", testPath)
	}
	
	// Check output message
	if !strings.Contains(string(output), "Successfully created") {
		t.Errorf("Expected success message, got: %s", output)
	}
}

// TestDirectoryCreateCommand tests the directory command for creation
func TestDirectoryCreateCommand(t *testing.T) {
	tempDir := t.TempDir()
	testPath := filepath.Join(tempDir, "testdir")
	
	// Build directory command
	buildCmd := exec.Command("go", "build", "-o", "/tmp/fs-test-directory", "./cmd/directory")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build directory command: %v", err)
	}
	
	// Run directory create command
	cmd := exec.Command("/tmp/fs-test-directory", "create", testPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Directory create command failed: %v, output: %s", err, output)
	}
	
	// Check that directory was created
	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		t.Errorf("Directory should have been created at %s", testPath)
	}
	
	// Check output message
	if !strings.Contains(string(output), "Successfully created") {
		t.Errorf("Expected success message, got: %s", output)
	}
}

// TestListCommand tests the list command
func TestListCommand(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create some test files and directories
	os.WriteFile(filepath.Join(tempDir, "file1.txt"), []byte("content1"), 0644)
	os.WriteFile(filepath.Join(tempDir, "file2.txt"), []byte("content2"), 0644)
	os.Mkdir(filepath.Join(tempDir, "subdir"), 0755)
	
	// Build list command
	buildCmd := exec.Command("go", "build", "-o", "/tmp/fs-test-list", "./cmd/list")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build list command: %v", err)
	}
	
	// Run list command
	cmd := exec.Command("/tmp/fs-test-list", tempDir)
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("List command failed: %v", err)
	}
	
	outputStr := string(output)
	
	// Check that all items are listed
	if !strings.Contains(outputStr, "file1.txt") {
		t.Error("Should list file1.txt")
	}
	if !strings.Contains(outputStr, "file2.txt") {
		t.Error("Should list file2.txt")
	}
	if !strings.Contains(outputStr, "subdir") {
		t.Error("Should list subdir")
	}
}

// TestCopyCommand tests the copy command
func TestCopyCommand(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create source file
	sourcePath := filepath.Join(tempDir, "source.txt")
	content := []byte("test content for copy")
	os.WriteFile(sourcePath, content, 0644)
	
	destPath := filepath.Join(tempDir, "dest.txt")
	
	// Build copy command
	buildCmd := exec.Command("go", "build", "-o", "/tmp/fs-test-copy", "./cmd/copy")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build copy command: %v", err)
	}
	
	// Run copy command
	cmd := exec.Command("/tmp/fs-test-copy", sourcePath, destPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Copy command failed: %v, output: %s", err, output)
	}
	
	// Check that file was copied
	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		t.Error("Destination file should exist")
	}
	
	// Check content
	copiedContent, err := os.ReadFile(destPath)
	if err != nil {
		t.Fatalf("Failed to read copied file: %v", err)
	}
	
	if string(copiedContent) != string(content) {
		t.Errorf("Content mismatch. Expected: %s, Got: %s", content, copiedContent)
	}
}

// TestInfoCommand tests the info command
func TestInfoCommand(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create test file
	testPath := filepath.Join(tempDir, "testfile.txt")
	content := []byte("test content")
	os.WriteFile(testPath, content, 0644)
	
	// Build info command
	buildCmd := exec.Command("go", "build", "-o", "/tmp/fs-test-info", "./cmd/info")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build info command: %v", err)
	}
	
	// Run info command
	cmd := exec.Command("/tmp/fs-test-info", testPath)
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Info command failed: %v", err)
	}
	
	outputStr := string(output)
	
	// Check that output contains expected information
	if !strings.Contains(outputStr, "testfile.txt") {
		t.Error("Should contain filename")
	}
	if !strings.Contains(outputStr, "file") {
		t.Error("Should contain type information")
	}
}

// TestMoveCommand tests the move command
func TestMoveCommand(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create source file
	sourcePath := filepath.Join(tempDir, "source.txt")
	content := []byte("test content for move")
	os.WriteFile(sourcePath, content, 0644)
	
	destPath := filepath.Join(tempDir, "dest.txt")
	
	// Build move command
	buildCmd := exec.Command("go", "build", "-o", "/tmp/fs-test-move", "./cmd/move")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build move command: %v", err)
	}
	
	// Run move command
	cmd := exec.Command("/tmp/fs-test-move", sourcePath, destPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Move command failed: %v, output: %s", err, output)
	}
	
	// Check that source file no longer exists
	if _, err := os.Stat(sourcePath); !os.IsNotExist(err) {
		t.Error("Source file should not exist after move")
	}
	
	// Check that destination file exists
	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		t.Error("Destination file should exist")
	}
	
	// Check content
	movedContent, err := os.ReadFile(destPath)
	if err != nil {
		t.Fatalf("Failed to read moved file: %v", err)
	}
	
	if string(movedContent) != string(content) {
		t.Errorf("Content mismatch. Expected: %s, Got: %s", content, movedContent)
	}
}

// TestDeleteCommand tests the delete command
func TestDeleteCommand(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create test file
	testPath := filepath.Join(tempDir, "testfile.txt")
	os.WriteFile(testPath, []byte("to be deleted"), 0644)
	
	// Build delete command
	buildCmd := exec.Command("go", "build", "-o", "/tmp/fs-test-delete", "./cmd/delete")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build delete command: %v", err)
	}
	
	// Run delete command
	cmd := exec.Command("/tmp/fs-test-delete", testPath, "--force")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Delete command failed: %v, output: %s", err, output)
	}
	
	// Check that file no longer exists
	if _, err := os.Stat(testPath); !os.IsNotExist(err) {
		t.Error("File should not exist after deletion")
	}
}

// TestBothCommandSyntaxes tests that both action-object and object-action syntaxes work
func TestBothCommandSyntaxes(t *testing.T) {
	tempDir := t.TempDir()
	
	// Test action-object syntax: create directory
	testPath1 := filepath.Join(tempDir, "testdir1")
	buildCmd := exec.Command("go", "build", "-o", "/tmp/fs-test-create", "./cmd/create")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build create command: %v", err)
	}
	
	cmd := exec.Command("/tmp/fs-test-create", "directory", testPath1)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Create directory command failed: %v", err)
	}
	
	// Test object-action syntax: directory create
	testPath2 := filepath.Join(tempDir, "testdir2")
	buildCmd = exec.Command("go", "build", "-o", "/tmp/fs-test-directory", "./cmd/directory")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build directory command: %v", err)
	}
	
	cmd = exec.Command("/tmp/fs-test-directory", "create", testPath2)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Directory create command failed: %v", err)
	}
	
	// Both directories should exist
	if _, err := os.Stat(testPath1); os.IsNotExist(err) {
		t.Error("Directory created with action-object syntax should exist")
	}
	if _, err := os.Stat(testPath2); os.IsNotExist(err) {
		t.Error("Directory created with object-action syntax should exist")
	}
}