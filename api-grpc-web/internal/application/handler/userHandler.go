package handler

import (
	"api-grpc-web/internal/domain/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserService service.UsersService
}

func CreateUserHandler(userService service.UsersService) UserHandler {
	return UserHandler{UserService: userService}
}

func (h *UserHandler) FindUser(ctx echo.Context) error {
	username := ctx.QueryParam("username")
	user, err := h.UserService.FindUser(ctx.Request().Context(), username)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}