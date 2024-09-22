package dto

type LoginResponse struct {
	Id       *int    `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Age      *int    `json:"age"`
	Token    struct {
		Token     string `json:"token"`
		ExpiredAt string `json:"expired_at"`
	} `json:"token"`
}
