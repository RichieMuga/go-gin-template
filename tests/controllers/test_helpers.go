package controllers

import (
	"github.com/RichieMuga/go-gin-template/routes"
	"github.com/gin-gonic/gin"
)

func SetupTestRouter() *gin.Engine {
  // Reduce noisy output for tests
  gin.SetMode(gin.TestMode)

  // Init router
  router := gin.Default()

  // Register routes
  routes.InitializeRoutes(router)

  return router
}
