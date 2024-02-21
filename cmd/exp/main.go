package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/baimiyishu13/lenslocked/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (c PostgresConfig) Starting() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "root",
		Password: "pgsql@QWE1113!",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
	fmt.Println(cfg.Starting())
	db, err := sql.Open("pgx", cfg.Starting())
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping database: ", err)
	}
	fmt.Println("database connected")

	us := models.UserService{
		DB: db,
	}

	// _, err = us.DB.Exec(`
	// 	Create TABLE users (
	// 		id SERIAL PRIMARY KEY,
	// 		email TEXT UNIQUE NOT NULL,
	// 		passwordHash TEXT UNIQUE NOT NULL
	// 	);
	// `)
	if err != nil {
		log.Fatal("failed to create database: ", err)
	}

	user, err := us.Create("bob2@bob.com", "bob2@bob123!")
	if err != nil {
		log.Fatal("failed to create user: ", err)
	}
	fmt.Println(*user)
}
