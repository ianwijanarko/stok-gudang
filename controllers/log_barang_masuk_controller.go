package controllers

import (
	"net/http"
	"time"

	"stok-gudang/config"
	"stok-gudang/models"

	"github.com/gin-gonic/gin"
)

// LogBarangMasuk - Tambah log barang masuk
func LogBarangMasuk(c *gin.Context) {
	var log models.LogBarangMasuk
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi produk dan user
	var produk models.Produk
	var user models.User
	if err := config.DB.First(&produk, log.IDProduk).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}
	if err := config.DB.First(&user, log.IDUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	log.CreatedAt = time.Now()
	if err := config.DB.Create(&log).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan log barang masuk"})
		return
	}

	c.JSON(http.StatusOK, log)
}

// GetLogBarangMasuk - Ambil semua log barang masuk
func GetLogBarangMasuk(c *gin.Context) {
	var logs []models.LogBarangMasuk
	if err := config.DB.Preload("Produk").Preload("User").Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil log barang masuk"})
		return
	}
	c.JSON(http.StatusOK, logs)
}
