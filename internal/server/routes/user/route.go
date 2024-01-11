package user

import (
	repository "gameJoin/internal/data/repository/user"
	controller "gameJoin/internal/server/controller/user"
	handler "gameJoin/pkg/handlers/user"

	"github.com/labstack/echo/v4"
)

type RoutesUser struct {
	Controller controller.ControllerUser
	Repository handler.IHandlerUser
}

func (r *RoutesUser) RoutesInit(g *echo.Group) {
	r.Controller.UserHandler = handler.HandlerUser{
		Repository: repository.User{},
	}

}


func (r RoutesUser) RouteImplement (){


}