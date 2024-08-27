package main

import (
	"github.com/RichieMuga/go-gin-template/database"
	"github.com/RichieMuga/go-gin-template/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
  // Initialize the logger
  logger.InitLogger()
  
	// Initialize router
  router := gin.Default()

	// Test api
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
    logger.Info("ping request sent")
	})

  // Connect database
  database.ConnectDatabase()

  // Start the server and check for errors
  router.Run()
}
