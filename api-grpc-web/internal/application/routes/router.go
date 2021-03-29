package routes

import (
	"api-grpc-web/internal/application/handler"
	"api-grpc-web/internal/domain/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(productService service.UsersService) *echo.Echo {
	userHandler := handler.CreateUserHandler(productService)

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/v1/users", userHandler.FindUser)

	return e
}