package user

import (
	"context"
	"database/sql"

	"github.com/lmbangel/_novice/internal/db"
)

type UserRepository interface {
	CreateUser(ctx context.Context, params CreateUserParams) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
	GetUserByUsername(ctx context.Context, u_name string) (*User, error)
	UpdateUser(ctx context.Context, params UpdateUserParams) (*User, error)
}

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Timestamp sql.NullTime
}

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UpdateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       int64  `json:"id"`
}

func fmtUser(user db.User) *User {
	return &User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func fmtUsers(users []db.User) []User {
	u_new := make([]User, len(users))

	for i, u := range users {
		u_new[i] = User{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		}
	}
	return u_new
}
