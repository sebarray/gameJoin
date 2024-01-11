package user

import (
	"gameJoin/pkg/domain/user"
)

func (u HandlerUser) CreateUserRegistry(userRegistry user.UserRegistryRequest) user.UserRegistryResponse {
	_, err := u.Repository.CreateUserRegistry(userRegistry)
	if err != nil {
		return user.UserRegistryResponse{
			Code: 500,
		}
	}
	return user.UserRegistryResponse{
		Code: 200,
	}
}
