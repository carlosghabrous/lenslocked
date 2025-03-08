package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, password string) (*User, error) {
	emailLower := strings.ToLower(email)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	hashedPassword := string(hashedBytes)
	user := User{Email: emailLower, PasswordHash: hashedPassword}
	row := us.DB.QueryRow(`
        INSERT INTO users (email, password_hash)
        VALUES ($1, $2) RETURNING id`, emailLower, hashedPassword)

	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("error scanning user id: %w", err)
	}

	return &user, nil

}
