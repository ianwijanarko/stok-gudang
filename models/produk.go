package models

type Produk struct {
	ID               uint   `gorm:"primaryKey"`
	NamaProduk       string `gorm:"size:50"`
	IDKategoriProduk uint   `gorm:"not null"`
}
