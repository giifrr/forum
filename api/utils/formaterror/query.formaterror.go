/*
	This file used for formatting error messages because database errors
*/
package formaterror

import "strings"

var errorMessages = make(map[string]string)

var err error

func FormatError(errstring string) map[string]string {
	if strings.Contains(errstring, "username") {
		errorMessages["Taken_username"] = "Username Already Taken"
	}

	if strings.Contains(errstring, "record") {
		errorMessages["No_record"] = "No record found"
	}

	if strings.Contains(errstring, "password") {
		errorMessages["Incorrect_password"] = "Incorrect password"
	}

	if len(errorMessages) > 0 {
		return errorMessages
	}

	return nil
}
