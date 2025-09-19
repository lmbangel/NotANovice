package user

import (
	"context"
	"database/sql"

	"github.com/lmbangel/_novice/internal/db"
)

type sqliteRepository struct {
	db *sql.DB
}

func NewSQLiteUserRepository(db *sql.DB) UserRepository {
	return &sqliteRepository{db: db}
}

func (r *sqliteRepository) CreateUser(ctx context.Context, params CreateUserParams) (*User, error) {
	q := db.New(r.db)

	user, err := q.CreateUser(ctx, db.CreateUserParams(params))
	if err != nil {
		return nil, err
	}
	return fmtUser(user), nil
}

func (r *sqliteRepository) GetUsers(ctx context.Context) ([]User, error) {
	q := db.New(r.db)

	users, err := q.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return fmtUsers(users), nil
}

func (r *sqliteRepository) GetUserByID(ctx context.Context, id int64) (*User, error) {
	q := db.New(r.db)

	user, err := q.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return fmtUser(user), nil
}

func (r *sqliteRepository) GetUserByUsername(ctx context.Context, u_name string) (*User, error) {
	q := db.New(r.db)

	user, err := q.GetUserByUsername(ctx, u_name)
	if err != nil {
		return nil, err
	}
	return fmtUser(user), nil
}

func (r *sqliteRepository) UpdateUser(ctx context.Context, params UpdateUserParams) (*User, error) {
	q := db.New(r.db)

	user, err := q.UpdateUser(ctx, db.UpdateUserParams(params))
	if err != nil {
		return nil, err
	}
	return fmtUser(user), nil
}
