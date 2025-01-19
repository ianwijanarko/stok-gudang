package controllers

import (
	"net/http"

	"stok-gudang/config"

	"github.com/gin-gonic/gin"
)

type StokResponse struct {
	IDProduk    uint   `json:"id_produk"`
	NamaProduk  string `json:"nama_produk"`
	TotalMasuk  int    `json:"total_masuk"`
	TotalKeluar int    `json:"total_keluar"`
	Stok        int    `json:"stok"`
}

// ShowStok - Menampilkan stok setiap produk
func ShowStok(c *gin.Context) {
	var stok []StokResponse

	query := `
		SELECT 
			p.id AS id_produk,
			p.namaproduk,
			COALESCE(SUM(lbm.jumlah), 0) AS total_masuk,
			COALESCE(SUM(lbk.jumlah), 0) AS total_keluar,
			COALESCE(SUM(lbm.jumlah), 0) - COALESCE(SUM(lbk.jumlah), 0) AS stok
		FROM 
			produk p
		LEFT JOIN 
			log_barang_masuk lbm ON p.id = lbm.idproduk
		LEFT JOIN 
			log_barang_keluar lbk ON p.id = lbk.idproduk
		GROUP BY 
			p.id, p.namaproduk;
	`

	if err := config.DB.Raw(query).Scan(&stok).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data stok"})
		return
	}

	c.JSON(http.StatusOK, stok)
}
