package main

import (
	"api-grpc-web/internal/application/routes"
	"api-grpc-web/internal/domain/service"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {
	userService := service.NewUserService()

	router := routes.Route(userService)

	port := os.Getenv("port")
	err := router.Start(":" + port)
	if err != nil {
		log.Error("Error to start application.", err)
	}
}
