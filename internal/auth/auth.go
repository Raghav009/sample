package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func CreateHashedPassword(password string) (string, error) {
	passByte := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
