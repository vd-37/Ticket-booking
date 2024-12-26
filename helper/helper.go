package helper

import (
	"strings"
)

/* Method to validate the user input */
/* capitalize the first letter of the method(in Go) to make it exportable */
func ValidateUserEmail(email string) bool {
	return strings.Contains(email, "@")
}
