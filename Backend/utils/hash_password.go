package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain password using bcrypt with the default cost.
func HashPassword(password string) (string, error) {
	// bcrypt.DefaultCost is a recommended cost value (currently 10).
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err // Return error if password hashing fails.
	}
	return string(hashedBytes), nil
}

// CheckPasswordHash compares a plaintext password with a hashed password.
func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
