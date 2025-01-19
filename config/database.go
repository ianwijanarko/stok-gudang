package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase - Inisialisasi koneksi database
func ConnectDatabase() (*gorm.DB, error) {
	// Ubah sesuai dengan konfigurasi MySQL Anda
	dsn := "root:password@tcp(127.0.0.1:3306)/stok_gudang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db
	log.Println("Berhasil terhubung ke database!")
	return db, nil
}
