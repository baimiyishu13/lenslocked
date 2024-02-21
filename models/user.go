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
	// TODO: implement
	email = strings.TrimSpace(email)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash the password: %v", err)
	}
	passwordHash := string(hashedBytes)

	user := User{
		Email:        email,
		PasswordHash: passwordHash,
	}

	// 插入数据
	row := us.DB.QueryRow(`
	INSERT INTO users (email, passwordHash)
	VALUES ($1, $2) RETURNING id;`, email, passwordHash)
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %v", err)
	}
	return &user, nil
}

func (us *UserService) Update(user *User) error {
	// TODO: implement
	return nil
}
