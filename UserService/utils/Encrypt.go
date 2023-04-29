package utils

import "golang.org/x/crypto/bcrypt"

func Hash(str string) (string, error) {
	password := []byte(str)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CompareHash(str string, hashedStr string) bool {
	hashedPassword := []byte(hashedStr)
	password := []byte(str)
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}
