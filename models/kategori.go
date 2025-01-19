package models

type KategoriProduk struct {
	ID           uint   `gorm:"primaryKey"`
	NamaKategori string `gorm:"size:50"`
}
