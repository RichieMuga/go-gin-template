package main

import (
	"github.com/gin-gonic/gin"
	"github.com/RichieMuga/go-gin-template/database"
)

func main() {
	router := gin.Default()

	// Test api
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

  // Connect database
  database.ConnectDatabase()

  // Start the server and check for errors
  router.Run()
}
