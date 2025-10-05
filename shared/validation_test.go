package shared

import (
	"path/filepath"
	"testing"
)

func TestValidateWorkshopID(t *testing.T) {
	tests := []struct {
		name        string
		workshopID  string
		expectError bool
	}{
		{"Valid ID", "123456789", false},
		{"Empty ID", "", true},
		{"Zero ID", "0", true},
		{"Invalid characters", "abc123", true},
		{"Negative number", "-123", true},
		{"Very large number", "999999999999999999", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateWorkshopID(tt.workshopID)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateWorkshopID(%q) error = %v, expectError %v", tt.workshopID, err, tt.expectError)
			}
		})
	}
}

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name        string
		username    string
		expectError bool
	}{
		{"Valid username", "testuser123", false},
		{"Valid with underscore", "test_user", false},
		{"Valid with period", "test.user", false},
		{"Valid with hyphen", "test-user", false},
		{"Empty username", "", true},
		{"Too short", "ab", true},
		{"Too long", "a" + string(make([]byte, 65)), true},
		{"Invalid characters", "test@user", true},
		{"Valid minimum length", "abc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUsername(tt.username)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateUsername(%q) error = %v, expectError %v", tt.username, err, tt.expectError)
			}
		})
	}
}

func TestValidateFilePath(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		expectError bool
	}{
		{"Valid path", "test/path", false},
		{"Empty path", "", true},
		{"Path traversal", "../../../etc/passwd", true},
		{"Valid absolute path", "/home/user/file", false},
		{"Windows path", "C:\\Users\\test\\file.txt", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFilePath(tt.path)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateFilePath(%q) error = %v, expectError %v", tt.path, err, tt.expectError)
			}
		})
	}
}

func TestSanitizePath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Remove quotes", `"test/path"`, filepath.Clean("test/path")},
		{"Remove single quotes", "'test/path'", filepath.Clean("test/path")},
		{"Clean path", "test/../other/./path", filepath.Clean("other/path")},
		{"No changes needed", "test/path", filepath.Clean("test/path")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizePath(tt.input)
			if result != tt.expected {
				t.Errorf("SanitizePath(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	// Test with a file that should exist
	if !FileExists("validation.go") {
		t.Error("FileExists should return true for existing file validation.go")
	}

	// Test with a file that should not exist
	if FileExists("nonexistent_file_12345.txt") {
		t.Error("FileExists should return false for non-existent file")
	}
}