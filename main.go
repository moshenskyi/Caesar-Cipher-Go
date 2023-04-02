package main

import (
	"github.com/gin-gonic/gin"
	"moshenskyi/caesar/pkg/handlers"
)

func main() {
	router := gin.Default()

	router.POST("/encrypt", handlers.Encrypt)

	router.Run("localhost:9090")
}
