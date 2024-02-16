package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// export DATABASE_URL='postgres://username:password@localhost:5432/database_name'
	// 从环境变量中获取数据库连接信息
	// export DATABASE_URL='postgres://root:pgsql@QWE1113!@localhost:5432/lenslocked'
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		// 如果创建连接池出错，打印错误信息并退出程序
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close() // 确保在程序退出时关闭连接池，释放资源

	err = dbpool.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ping failed: %v\n", err)
	}

	fmt.Println("Connected!")

	// Create database
	_, err = dbpool.Exec(context.Background(), "create database users;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
		os.Exit(1)
	}

	// Create table
	_, err = dbpool.Exec(context.Background(), "CREATE TABLE users (id SERIAL PRIMARY KEY,age INT,firstname TEXT NOT NULL,lastname TEXT NOT NULL,email TEXT UNIQUE);")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
		os.Exit(1)
	}

	// Insert data
	_, err = dbpool.Exec(context.Background(), "INSERT INTO users (age, firstname, lastname, email) VALUES(30, 'jon', 'jontom', 'jontom@test.com');")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
		os.Exit(1)
	}
	// Select users from database
	_, err = dbpool.Exec(context.Background(), "SELECT * FROM users;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
		os.Exit(1)
	}

}
