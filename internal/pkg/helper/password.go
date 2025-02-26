package helper

import (
	"github.com/pebruwantoro/movie-festival-backend/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password+config.Get().SaltPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ValidatePassword(storedHash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password+config.Get().SaltPassword))
	return err == nil
}
