package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgres() *sql.DB {
	host := getEnv("POSTGRES_HOST", "localhost")
	port := getEnv("POSTGRES_PORT", "5432")
	user := getEnv("POSTGRES_USER", "dev")
	pass := getEnv("POSTGRES_PASSWORD", "devpassword")
	dbname := getEnv("POSTGRES_DB", "devdb")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	log.Println("connected to Postgres")
	return db
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
