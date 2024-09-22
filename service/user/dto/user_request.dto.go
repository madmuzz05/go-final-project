package dto

type UserRequest struct {
	Username *string `json:"username" binding:"required"`
	Email    *string `json:"email" binding:"required,email"`
	Password *string `json:"password" binding:"required,min=6"`
	Age      *int    `json:"age" binding:"required,gte=8"`
	Id       int     `json:"id"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}
