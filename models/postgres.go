package models

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Open 打开数据库连接
// Postgres Database 连接最终关闭
// db.Close()
func Open(config PostgresConfig) (*sql.DB, error) {

	db, err := sql.Open("pgx", config.Starting())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return db, err

}

// 连接数据库默认 - 开发
func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "root",
		Password: "pgsql@QWE1113!",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
}

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
