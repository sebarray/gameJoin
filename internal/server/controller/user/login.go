package user

import (
	"gameJoin/pkg/domain/user"

	"github.com/labstack/echo/v4"
)

func (Controller ControllerUser) Login(c echo.Context) error {
	var user user.LoginRequest
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	userResponse := Controller.UserHandler.Login(user)
	return c.JSON(userResponse.Code, userResponse.Body)
}
