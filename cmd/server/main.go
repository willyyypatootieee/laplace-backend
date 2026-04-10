package main

import (
	"laplace-backend/internal/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	docs.SwaggerInfo.Title = "Acost Backend RestAPI"
	docs.SwaggerInfo.Description = "This is a simple API for Laplace Backend."
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:" + port
	docs.SwaggerInfo.BasePath = "/api/v1"
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.RegisterRoutes(r)

	log.Println("Server running on : " + port)
	r.Run(":" + port)
}
