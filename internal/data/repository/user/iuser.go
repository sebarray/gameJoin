package user

import (
	"gameJoin/internal/data/repository/user/dto"
	"gameJoin/pkg/domain/user"
)

type User struct {
}

type IUser interface {
	InitLogin(u user.LoginRequest) dto.User
	Login(userDTO dto.User) (*dto.User, error)
	CreateUserRegistry(userRegistry user.UserRegistryRequest) (user.UserRegistryRequest, error)
}
