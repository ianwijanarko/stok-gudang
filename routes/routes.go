package routes

import (
	"stok-gudang/controllers"
	"stok-gudang/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Routes untuk User (login dan logout)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)

	// Routes untuk Kategori (dengan autentikasi)
	r.POST("/kategori", middleware.AuthMiddleware(), controllers.CreateKategori)
	r.GET("/kategori", middleware.AuthMiddleware(), controllers.GetAllKategori)
	r.PUT("/kategori/:id", middleware.AuthMiddleware(), controllers.UpdateKategori)
	r.DELETE("/kategori/:id", middleware.AuthMiddleware(), controllers.DeleteKategori)

	// Routes untuk Produk (dengan autentikasi)
	r.POST("/produk", middleware.AuthMiddleware(), controllers.CreateProduk)
	r.GET("/produk", middleware.AuthMiddleware(), controllers.GetAllProduk)
	r.PUT("/produk/:id", middleware.AuthMiddleware(), controllers.UpdateProduk)
	r.DELETE("/produk/:id", middleware.AuthMiddleware(), controllers.DeleteProduk)

	// Routes untuk User
	r.GET("/user", middleware.AuthMiddleware(), controllers.GetAllUser)
	r.PUT("/user/:id", middleware.AuthMiddleware(), controllers.UpdateUser)
	r.DELETE("/user/:id", middleware.AuthMiddleware(), controllers.DeleteUser)

	// Routes untuk Log Barang Masuk
	r.POST("/log_masuk", middleware.AuthMiddleware(), controllers.LogBarangMasuk)
	r.GET("/log_masuk", middleware.AuthMiddleware(), controllers.GetLogBarangMasuk)

	// Routes untuk Log Barang Keluar
	r.POST("/log_keluar", middleware.AuthMiddleware(), controllers.LogBarangKeluar)
	r.GET("/log_keluar", middleware.AuthMiddleware(), controllers.GetLogBarangKeluar)

	r.GET("/stok", middleware.AuthMiddleware(), controllers.ShowStok)

	return r
}
