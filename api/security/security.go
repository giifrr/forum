package security

import "golang.org/x/crypto/bcrypt"

func Hash(s string) ([]byte, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return res, err
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
