package user

import (
	repository "gameJoin/internal/data/repository/user"
	"gameJoin/pkg/domain/user"
)

type HandlerUser struct {
	Repository repository.IUser
}

type IHandlerUser interface {
	Login(login user.LoginRequest) user.LoginResponse
	CreateUserRegistry(userRegistry user.UserRegistryRequest) user.UserRegistryResponse
}
