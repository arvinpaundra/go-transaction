package dto

type (
	GetUserRequest struct {
		ID int `form:"user_id"`
	}
	GetUserResponse struct {
		ID    int    `json:"user_id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	InsertUserRequest struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
	}

	UpdateUserRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)
