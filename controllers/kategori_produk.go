package controllers

import (
	"net/http"
	"stok-gudang/config"
	"stok-gudang/models"

	"github.com/gin-gonic/gin"
)

// CreateKategori - Tambah kategori baru
func CreateKategori(c *gin.Context) {
	var kategori models.KategoriProduk
	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, kategori)
}

// GetAllKategori - Ambil semua kategori
func GetAllKategori(c *gin.Context) {
	var kategori []models.KategoriProduk
	if err := config.DB.Find(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, kategori)
}

// UpdateKategori - Perbarui kategori berdasarkan ID
func UpdateKategori(c *gin.Context) {
	id := c.Param("id")
	var kategori models.KategoriProduk

	if err := config.DB.First(&kategori, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, kategori)
}

// DeleteKategori - Hapus kategori berdasarkan ID
func DeleteKategori(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.KategoriProduk{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil dihapus"})
}
