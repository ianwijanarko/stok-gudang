package main

import (
	"log"
	"stok-gudang/config"
	"stok-gudang/routes"
)

func main() {
	// Inisialisasi database
	_, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Gagal menghubungkan ke database: %v", err)
	}

	// Setup rute
	r := routes.SetupRoutes()

	// Menjalankan server
	log.Println("Server berjalan di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
