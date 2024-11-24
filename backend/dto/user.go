package dto

// RegisterRequest 用户注册请求结构
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// LoginRequest 用户登录请求结构
type LoginRequest struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse 用户登录响应结构
type LoginResponse struct {
	Basic
	Token string `json:"token"`
}

type UserBasic struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// UserResponse 用户信息响应结构
type UserResponse struct {
	Basic
	Data UserBasic `json:"data"`
}

type UsersResponse struct {
	Basic
	Data []UserBasic `json:"data"`
}
