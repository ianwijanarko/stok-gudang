package controllers

import (
	"net/http"
	"stok-gudang/config"
	"stok-gudang/models"

	"github.com/gin-gonic/gin"
)

// CreateProduk - Tambah produk baru
func CreateProduk(c *gin.Context) {
	var produk models.Produk
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah kategori produk ada
	var kategori models.KategoriProduk
	if err := config.DB.First(&kategori, produk.IDKategoriProduk).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori produk tidak ditemukan"})
		return
	}

	if err := config.DB.Create(&produk).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, produk)
}

// GetAllProduk - Ambil semua produk
func GetAllProduk(c *gin.Context) {
	var produk []models.Produk
	if err := config.DB.Find(&produk).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, produk)
}

// UpdateProduk - Perbarui produk berdasarkan ID
func UpdateProduk(c *gin.Context) {
	id := c.Param("id")
	var produk models.Produk

	if err := config.DB.First(&produk, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah kategori produk ada
	var kategori models.KategoriProduk
	if err := config.DB.First(&kategori, produk.IDKategoriProduk).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori produk tidak ditemukan"})
		return
	}

	if err := config.DB.Save(&produk).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, produk)
}

// DeleteProduk - Hapus produk berdasarkan ID
func DeleteProduk(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Produk{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}
