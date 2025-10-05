package shared

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// ValidateWorkshopID checks if a workshop ID is valid (non-zero positive integer)
func ValidateWorkshopID(workshopID string) error {
	if workshopID == "" {
		return fmt.Errorf("workshop ID cannot be empty")
	}

	if workshopID == "0" {
		return fmt.Errorf("workshop ID cannot be zero")
	}

	id, err := strconv.ParseUint(workshopID, 10, 64)
	if err != nil {
		return fmt.Errorf("workshop ID must be a valid number: %w", err)
	}

	if id == 0 {
		return fmt.Errorf("workshop ID must be greater than zero")
	}

	return nil
}

// ValidateFilePath checks if a file path is safe and valid
func ValidateFilePath(path string) error {
	if path == "" {
		return fmt.Errorf("file path cannot be empty")
	}

	// Clean the path to resolve any .. or . components
	cleanPath := filepath.Clean(path)

	// Check for potentially dangerous path traversal
	if strings.Contains(cleanPath, "..") {
		return fmt.Errorf("path traversal not allowed: %s", path)
	}

	return nil
}

// ValidateUsername checks if a Steam username is valid
func ValidateUsername(username string) error {
	if username == "" {
		return fmt.Errorf("username cannot be empty")
	}

	if len(username) < 3 {
		return fmt.Errorf("username must be at least 3 characters long")
	}

	if len(username) > 64 {
		return fmt.Errorf("username must be no more than 64 characters long")
	}

	// Basic validation for Steam username characters
	validUsernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`)
	if !validUsernameRegex.MatchString(username) {
		return fmt.Errorf("username contains invalid characters (only letters, numbers, underscore, period, and hyphen allowed)")
	}

	return nil
}

// SanitizePath sanitizes a file path by cleaning it and removing quotes
func SanitizePath(path string) string {
	// Remove surrounding quotes
	path = strings.Trim(path, `"'`)

	// Clean the path
	return filepath.Clean(path)
}
