package utils

import (
	"github.com/google/uuid"
	"net/mail"
	"regexp"
	"unicode"
)

func IsUUIDValid(id uuid.UUID) bool {
	if id == uuid.Nil {
		return false
	}
	regex := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return regex.MatchString(id.String())
}

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
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

	return hasLetter && hasNumber && len(password) >= 8
}
