package utils

import (
	"github.com/google/uuid"
	"regexp"
	"unicode"
)

func IsUuidValid(id uuid.UUID) bool {
	if id == uuid.Nil {
		return false
	}
	regex := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return regex.MatchString(id.String())
}

func IsUrlValid(url string) bool {
	regex := regexp.MustCompile(`^(?:https?://)?(w{3}\.)?[\w_-]+((\.\w{2,}){1,3})(/([^/\n]+/?)*(\?[\w_-]+=[^?/&]*(&[\w_-]+=[^?/&]*)*)?)?$`)
	return regex.MatchString(url)
}

func IsEmailValid(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

func IsPasswordValid(password string) bool {
	hasLetter, hasNumber := false, false
	for _, char := range password {
		hasLetter = hasLetter || unicode.IsLetter(char)
		hasNumber = hasNumber || unicode.IsNumber(char)

		if hasLetter && hasNumber {
			break
		}
	}

	return hasLetter && hasNumber && len(password) >= 8 && len(password) <= 50
}
