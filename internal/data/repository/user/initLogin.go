package user

import (
	"gameJoin/internal/data/repository/user/dto"
	"gameJoin/pkg/domain/user"
)

func (User) InitLogin(u user.LoginRequest) dto.User {

	userDTO := dto.User{
		Email:        u.Email,
		PasswordHash: u.Password,
	}

	return userDTO
}
