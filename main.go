package main

import (
	"log"

	"github.com/Sid5488/go-crud/src/controllers/routes"
	controllers "github.com/Sid5488/go-crud/src/controllers/users"
	"github.com/Sid5488/go-crud/src/models/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init dependencies
	service := services.NewUserDomainService()
	userController := controllers.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
