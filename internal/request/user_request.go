package request

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required,max=100"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type UpdateUserRequest struct {
	Username *string `json:"username,omitempty" binding:"omitempty,max=100"`
	Email    *string `json:"email,omitempty" binding:"omitempty,email,max=100"`
	Password *string `json:"password,omitempty" binding:"omitempty,min=6,max=100"`
}