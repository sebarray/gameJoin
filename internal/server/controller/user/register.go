package user

import (
	"gameJoin/pkg/domain/user"

	"github.com/labstack/echo/v4"
)

func (Controller ControllerUser) Registry(c echo.Context) error {
	var user user.UserRegistryRequest
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	userResponse := Controller.UserHandler.CreateUserRegistry(user)
	return c.JSON(userResponse.Code, userResponse.Message)
}
