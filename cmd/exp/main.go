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

	// Create table
	_, err = dbpool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE NOT NULL,
			email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			amonut INT NOT NULL,
			description TEXT NOT NULL
		);
		`)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
		os.Exit(1)
	}

	// Insert some date ...
	// // name := "','');DROP TABLE users CASCADE; --"
	name := "New User2"
	email := "NewUser2@gmail.com"

	id := 1
	row := dbpool.QueryRow(context.Background(), `
		SELECT name,email
		FROM users
		WHERE id=$1`, id)
	err = row.Scan(&name, &email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("User name: ", name, "User email: ", email)

	type order struct {
		ID          int
		UserID      int
		Amonut      int
		Description string
	}
	var orders []order
	userID := 1
	rows, err := dbpool.Query(context.Background(), `
        SELECT id, user_id, amonut, description
        FROM orders
        WHERE user_id=$1`, userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var o order
		err = rows.Scan(&o.ID, &o.UserID, &o.Amonut, &o.Description)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
			os.Exit(1)
		}
		orders = append(orders, o)
	}

	err = rows.Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(orders)
}
