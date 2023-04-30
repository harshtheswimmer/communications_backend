package main

import (
	"communications_backend/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.Default()

	routes.TestRoutes(router)

	router.Run(": " + port)
}
