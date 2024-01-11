package user

import (
	"fmt"
	"gameJoin/pkg/domain/user"
)

func (Handler HandlerUser) Login(loginReq user.LoginRequest) user.LoginResponse {
	userLogin := Handler.Repository.InitLogin(loginReq)

	login, err := Handler.Repository.Login(userLogin)

	fmt.Println(login)
	if err != nil {
		var response user.LoginResponse
		return response
	}
	return user.LoginResponse{
		Code: 200,
	}
}
