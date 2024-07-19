package helper

import (
	"strings"
)

/* Method to validate the user input */
func ValidateUserEmail(email string) bool {
	return strings.Contains(email, "@")
}
