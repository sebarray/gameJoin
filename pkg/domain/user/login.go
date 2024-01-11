package user

type LoginRequest struct {
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Body LoginBodyResponse `json:"body"`
	Code int               `json:"code"`
}

type LoginBodyResponse struct {
	Message string `json:"message"`
	Jwt     string `json:"jwt"`
}
