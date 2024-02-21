package main

import (
	"fmt"
	"log"

	"github.com/baimiyishu13/lenslocked/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	fmt.Println("open database connection")

	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping database: ", err)
	}

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

	_, err = us.DB.Exec(`
	   Create TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
        email TEXT UNIQUE NOT NULL,
        passwordHash TEXT UNIQUE NOT NULL
	   );
	`)
	if err != nil {
		log.Fatal("failed to create users table: ", err)
	}

	user, err := us.Create("bob3@bob.com", "bob1@bob123!")
	if err != nil {
		log.Fatal("failed to create user: ", err)
	}
	fmt.Println(*user)
}
