package vdf

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuild(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create test files
	infoFile := filepath.Join(tempDir, "info.txt")
	if err := os.WriteFile(infoFile, []byte("Test Title\nTest Description\nLine 2 of description"), 0644); err != nil {
		t.Fatal(err)
	}

	changelogFile := filepath.Join(tempDir, "changelog.txt")
	if err := os.WriteFile(changelogFile, []byte("Test changelog content"), 0644); err != nil {
		t.Fatal(err)
	}

	item := VDFItem{
		WorkshopID:    "123456789",
		ContentFolder: tempDir,
	}

	result := Build(item)

	// Check if the result contains expected content
	expectedStrings := []string{
		`"workshopitem"`,
		`"appid"`,
		`"244850"`,
		`"publishedfileid"`,
		`"123456789"`,
		`"contentfolder"`,
		tempDir,
		`"visibility"`,
		`"0"`,
		`"title"`,
		`"Test Title"`,
		`"description"`,
		`"Test Description`,
		`Line 2 of description"`,
		`"changenote"`,
		`"Test changelog content"`,
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(result, expected) {
			t.Errorf("Build() result should contain %q, but got:\n%s", expected, result)
		}
	}
}

func TestBuildWithEmptyWorkshopID(t *testing.T) {
	item := VDFItem{
		WorkshopID:    "",
		ContentFolder: "/tmp/test",
	}

	result := Build(item)
	if result != "" {
		t.Errorf("Build() should return empty string for empty WorkshopID, got: %s", result)
	}
}

func TestReadme(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Test case 1: Valid info.txt file
	infoFile := filepath.Join(tempDir, "info.txt")
	content := "Test Workshop Title\nThis is the description\nSecond line of description"
	if err := os.WriteFile(infoFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	found, title, description := Readme(tempDir)
	if !found {
		t.Error("Readme() should find the info.txt file")
	}
	if title != "Test Workshop Title" {
		t.Errorf("Expected title 'Test Workshop Title', got '%s'", title)
	}
	expectedDesc := "This is the description\nSecond line of description"
	if description != expectedDesc {
		t.Errorf("Expected description '%s', got '%s'", expectedDesc, description)
	}

	// Test case 2: Non-existent file
	emptyDir := t.TempDir()
	found, title, description = Readme(emptyDir)
	if found {
		t.Error("Readme() should not find info.txt in empty directory")
	}
	if title != "" || description != "" {
		t.Error("Readme() should return empty strings when file not found")
	}
}

func TestChangelog(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Test case 1: Valid changelog.txt file
	changelogFile := filepath.Join(tempDir, "changelog.txt")
	content := "Version 1.2.0\n- Fixed bugs\n- Added new features"
	if err := os.WriteFile(changelogFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	found, changelog := Changelog(tempDir)
	if !found {
		t.Error("Changelog() should find the changelog.txt file")
	}
	if changelog != content {
		t.Errorf("Expected changelog '%s', got '%s'", content, changelog)
	}

	// Test case 2: Non-existent file
	emptyDir := t.TempDir()
	found, changelog = Changelog(emptyDir)
	if found {
		t.Error("Changelog() should not find changelog.txt in empty directory")
	}
	if changelog != "" {
		t.Error("Changelog() should return empty string when file not found")
	}
}
