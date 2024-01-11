package user

type UserRegistryRequest struct {
	ID int `json:"id" gorm:"primaryKey"`
	LoginRequest
	NicknameDiscord string `json:"nicknameDiscord"`
	Age             int    `json:"age"`
	Avatar          string `json:"avatar"`
}

type UserRegistryResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
