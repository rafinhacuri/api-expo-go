package utils

import (
	"errors"
	"regexp"
	"unicode"
)

func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("a senha deve ter pelo menos 6 caracteres")
	}

	var hasUpper, hasLower, hasNumber, hasSymbol bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(string(char)):
			hasSymbol = true
		}
	}

	if !hasUpper {
		return errors.New("a senha deve conter pelo menos uma letra maiúscula")
	}
	if !hasLower {
		return errors.New("a senha deve conter pelo menos uma letra minúscula")
	}
	if !hasNumber {
		return errors.New("a senha deve conter pelo menos um número")
	}
	if !hasSymbol {
		return errors.New("a senha deve conter pelo menos um símbolo")
	}

	return nil
}
