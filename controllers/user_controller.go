package controllers

import (
	"net/http"
	"stok-gudang/config"
	"stok-gudang/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("rahasia123")

func GenerateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"nama": user.Nama,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 24 jam
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Login - Autentikasi user dan menghasilkan JWT
func Login(c *gin.Context) {
	var loginData struct {
		Nama     string `json:"nama"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	if err := config.DB.Where("nama = ?", loginData.Nama).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT
	token, err := GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Kirim token ke client
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Logout - Tidak perlu server-side state, cukup hapus token di client
func Logout(c *gin.Context) {
	// Implementasi logout cukup dengan menghapus token di client
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

// CreateUser - Tambah user baru
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUser - Ambil semua user
func GetAllUser(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser - Perbarui user berdasarkan ID
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password jika diubah
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
			return
		}
		user.Password = string(hashedPassword)
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser - Hapus user berdasarkan ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}
