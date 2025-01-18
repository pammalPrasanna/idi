package dtos

type GetUsersParams struct {
	ID int64 `json:"id"`
}
type (
	FindUsersParams   struct{}
	CreateUsersParams struct{}
	UpdateUsersParams struct{}
	DeleteUsersParams struct{}
)

type User struct {
	ID int64 `json:"id"`
}

type FindUsersResponse struct {
	Users []*User `json:"users"`
}

type CreateUserResponse struct {
	UserID int64 `json:"user_id"`
}

type GetUserResponse struct {
	User *User `json:"user"`
}
