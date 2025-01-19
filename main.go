package main

import (
	"stok-gudang/config"
	"stok-gudang/routes"
)

func main() {
	// Connect to database
	config.ConnectDatabase()

	// Setup routes
	r := routes.SetupRoutes()

	// Run server
	r.Run(":8080")
}
