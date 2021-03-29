package main

import (
	"api-grpc-web/internal/application/routes"
	"api-grpc-web/internal/domain/service"
	"github.com/labstack/gommon/log"
)

func main() {
	userService := service.NewUserService()

	router := routes.Route(userService)

	err := router.Start(":8080")
	if err != nil {
		log.Error("Error to start application.", err)
	}
}
