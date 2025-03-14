package main

import "unicode"

func isValidPassword(password string) bool {
	var upperCount int
	var digitCount int

	if len(password) < 5 || len(password) > 12 {
		return false
	}

	for _, char := range password {
		if unicode.IsNumber(char) {
			digitCount++
		}
		if unicode.IsUpper(char) {
			upperCount++
		}
	}

	if upperCount > 0 && digitCount > 0 {
		return true
	}
	return false
}

/*

    At least 5 characters long but no more than 12 characters.
    Contains at least one uppercase letter.
    Contains at least one digit.

A string is really just a read-only slice of bytes. This means that you can use the same techniques you learned in the
previous lesson to iterate over the characters of a string.
Assignment

Implement the isValidPassword function by looping through each character in the password string.
Make sure the password is long enough and includes at least one uppercase letter and one digit.

Assume passwords consist of ASCII characters only.
*/
