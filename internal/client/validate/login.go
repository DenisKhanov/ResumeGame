package validate

import (
	errors2 "github.com/DenisKhanov/ResumeGame/internal/client/errors"
	"regexp"
)

const (
	minLenLogin = uint8(6)
	maxLenLogin = uint8(12)
)

// CheckLogin method checks the login against the criteria (length must be at least minLen and no more than maxLen characters,
// and it should not contain characters other than "-" and "_")
func CheckLogin(login string) error {
	var (
		hasMinLen     = false
		hasMaxLen     = false
		hasValidChars = false
	)

	// Check login length
	l := uint8(len(login))
	if l >= minLenLogin {
		hasMinLen = true
	}
	if l <= maxLenLogin {
		hasMaxLen = true
	}

	// Checking for valid characters
	isAlphaNum := regexp.MustCompile(`^[A-Za-z0-9_\-]+$`).MatchString
	if isAlphaNum(login) {
		hasValidChars = true
	}
	if !hasMinLen {
		return errors2.ErrLoginShort
	}
	if !hasMaxLen {
		return errors2.ErrLoginLong
	}
	if !hasValidChars {
		return errors2.ErrLoginChar
	}
	return nil
}
