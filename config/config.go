package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

var (
	DB        *sql.DB
	RedisPool *redis.Pool
)

func Initialize() {
	err := godotenv.Load("./environment/cred.env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize MySQL
	dsn := os.Getenv("DB_DSN")
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to ping MySQL: %v", err)
	}
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE
	)`

	if _, err := DB.Exec(query); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Println("Database migrated successfully.")

	// Initialize Redis
	RedisPool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS_HOST")) // Example: localhost:6379
		},
	}
}
