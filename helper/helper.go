package helper

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"
)

// generateRequestID creates a string in the format YYYYMMDDHHII concatenated with a random alphanumeric string.
// generateRequestID creates a string in the format YYYYMMDDHHII concatenated with a random alphanumeric string.
func GenerateRequestID(randomStringLength int) (string, error) {
	// Set the location to Africa/Lagos
	location, err := time.LoadLocation("Africa/Lagos") // Africa/Lagos (GMT+1)
	if err != nil {
		return "", err
	}

	// Get the current time in the desired format
	now := time.Now().In(location)
	timePart := now.Format("200601021504")

	// Generate a random alphanumeric string of the desired length
	randomString, err := generateRandomString(randomStringLength)
	if err != nil {
		return "", err
	}

	// Concatenate the time-based string with the random string
	fullRequestID := timePart + randomString
	return fullRequestID, nil
}

// generateRandomString creates a random alphanumeric string of the specified length.
func generateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}

	return string(bytes), nil
}

func GetDate() string {
	currentDate := time.Now()

	// Format the date
	formattedDate := currentDate.Format("02 January, 2006")

	// Print the formatted date
	return formattedDate
}

func GetTime() string {
	currentTime := time.Now()

	// Format the time
	formattedTime := currentTime.Format("3:04 PM")

	// Print the formatted time
	return formattedTime
}

// Utility function to format and escape special characters
func FormatAndEscape(format string, args ...interface{}) string {
	// Format the string with given arguments
	formattedString := fmt.Sprintf(format, args...)

	// Escape special characters, e.g., quotes
	escapedString := strings.ReplaceAll(formattedString, `"`, `\"`)

	// You can escape more characters as needed
	// escapedString = strings.ReplaceAll(escapedString, `'`, `\'`)  // For single quotes

	return escapedString
}
