package password

import (
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compares the password and hash version if equal
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

// HashPassword hash the password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10) // recommended default

	return string(bytes), err
}
