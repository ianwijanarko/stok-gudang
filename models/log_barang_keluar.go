package models

import "time"

type LogBarangKeluar struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Jumlah    int       `json:"jumlah"`
	IDProduk  uint      `json:"idproduk"`
	Produk    Produk    `gorm:"foreignKey:IDProduk" json:"produk"`
	CreatedAt time.Time `json:"created_at"`
	IDUser    uint      `json:"iduser"`
	User      User      `gorm:"foreignKey:IDUser" json:"user"`
}
