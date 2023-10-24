package auth

import (
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string, salt []byte) string {
	combined := append([]byte(password), salt...)
	hashedPassword, _ := bcrypt.GenerateFromPassword(combined, bcrypt.DefaultCost)
	return hex.EncodeToString(hashedPassword)
}

func GenerateRandomSalt() []byte {
	salt := make([]byte, 16) // You can adjust the salt length
	return salt
}
