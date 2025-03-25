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

func (us *UserService) Authenticate(email, password string) (*User, error) {
	email = strings.ToLower(email)
	user := User{Email: email}

	row := us.DB.QueryRow("SELECT id, password_hash FROM users WHERE email=$1", email)
	err := row.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		fmt.Println("password does not match")
		return nil, fmt.Errorf("invalid password %w", err)
	}
	return &user, nil
}
