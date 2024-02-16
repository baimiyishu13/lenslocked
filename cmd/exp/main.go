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
		// os.Exit(1)
		return
	}

	// Insert some date ...
	name := "','');DROP TABLE users CASCADE; --"
	email := "JonD@gmail.com"
	// queru := fmt.Sprintf(`
	// 	Insert into users (name, email)
	// 	VALUES ('%s', '%s');
	// `, name, email)
	// fmt.Printf("Execqueru: %s\n", queru)
	// _, err = dbpool.Exec(context.Background(), queru)
	_, err = dbpool.Exec(context.Background(), `
	INSERT INTO users (name, email)
	VALUES ($1, $2);`, name, email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
		os.Exit(1)
	}
}
