package api

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/kwamekyeimonies/sb-design-bk/config"
	"github.com/kwamekyeimonies/sb-design-bk/middleware"
	"github.com/kwamekyeimonies/sb-design-bk/router"
)

func ApiServer() {
	webServer := gin.Default()
	webServer.Use(middleware.CORSMiddleware())

	router.NonAuthenticatedRoutes(&webServer.RouterGroup)

	authenticatedRoutes := webServer.Group("/api/v1/")
	// authenticatedRoutes.Use(middleware.JWTMiddleware())
	router.AuthenticatedRoutes(authenticatedRoutes)

	config, err := config.LoadInitializer("")
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}

	webServer.Run(":" + config.SERVER_PORT)
}
