package dtos

type (
	User struct {
		ID             int64  `json:"id"`
		Username       string `json:"task"`
		Email          string `json:"description"`
		HashedPassword string `json:"-"`
		CreatedAt      string `json:"created_at"`
		UpdatedAt      string `json:"updated_at"`
	}
	GetUserParams struct {
		ID    int64
		Email string
	}
	FindUsersParams  struct{}
	CreateUserParams struct {
		Username string `json:"task"`
		Email    string `json:"description"`
		Password string `json:"password"`
	}
	UpdateUserParams struct {
		ID       int64  `json:"id"`
		Username string `json:"task"`
		Email    string `json:"description"`
	}
	DeleteUserParams struct {
		ID int64
	}
)
