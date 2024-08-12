package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"qrCode/pkg/database"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, salt, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}

func SavePassword(password string) (string, string, error) {
	// generate salt for password
	salt := GenerateSalt()
	// hash password
	hashedPassword, err := HashPassword(password + salt)
	if err != nil {
		return "", "", err
	}

	// TODO Save to database and only return err
	return hashedPassword, salt, nil
}

func GenerateSalt() string {
	// generate random salt for password
	runeSet := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	salt := make([]rune, 8)
	for i := range salt {
		salt[i] = runeSet[rand.Intn(len(runeSet))]
	}
	return string(salt)
}

// users in the database

func AddUser(username, password, email string) error {
	// check if user already exists
	err := database.UserExists(username, email)
	if err != nil {
		return err
	}

	passwordHash, salt, err := SavePassword(password)
	if err != nil {
		return err
	}

	// add user to database
	err = database.AddUser(username, passwordHash, salt, email)
	if err != nil {
		return err
	}

	return nil
}

var ErrInvalidPassword = errors.New("invalid password")

func CheckPassword(username, password string) error {
	passwordHash, salt, err := database.GetUser(username)
	if err != nil {
		return err
	}

	if !CheckPasswordHash(password, salt, passwordHash) {
		return ErrInvalidPassword
	}

	return nil
}
