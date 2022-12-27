package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user-generator/infra/log"
	"user-generator/internal/domain/dto"
	"user-generator/internal/service"
)

var logger = log.NewLogger()

type UserController struct {
	userService service.UserServiceAdapter
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (controller *UserController) CreateUser(ctx echo.Context) error {
	userReq := &dto.UserDTO{}

	binder := echo.DefaultBinder{}

	if err := binder.BindBody(ctx, userReq); err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Trace("init user service create..")
	err := controller.userService.CreateUser(ctx.Request().Context(), userReq)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return ctx.JSON(http.StatusCreated, "user created with success")

}

func (controller *UserController) Register(server *echo.Echo) {
	v1 := server.Group("v1")
	v1.POST("/testbank", controller.CreateUser)
}
