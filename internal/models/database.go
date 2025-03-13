package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// init db malas jir
func InitDB() {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("❌ Gagal koneksi ke database:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("❌ Database tidak merespon:", err)
	}

	fmt.Println("✅ Database Connected!")
}
