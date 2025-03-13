package main

import (
	"log"

	// "net/http"
	// "os"

	"cardinal-re/internal/server"
	// _ "github.com/lib/pq" // PostgreSQL driver
)

// Konfigurasi database dari environment variables atau langsung di kode ( malas kalo langsung ke production)
// var (
// 	DB_HOST     = "localhost"
// 	DB_PORT     = "5432"
// 	DB_USER     = ""
// 	DB_PASSWORD = ""
// 	DB_NAME     = "ctf_competition"
// )

func main() {
	log.Println("🚀 Starting ... Tunggu bentar")

	// // Buat koneksi ke database
	// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal("❌ Gagal konek ke database:", err)
	// }
	// defer db.Close()

	// // Cek apakah koneksi database berhasil
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal("❌ Database tidak merespons:", err)
	// }

	// fmt.Println("✅ Database Connected!")

	server.StartServer()
}
