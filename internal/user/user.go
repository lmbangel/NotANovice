package user

import (
	"context"
	"database/sql"

	"github.com/lmbangel/_novice/internal/db"
	_ "modernc.org/sqlite"
)

func GetUsers() ([]db.User, error) {
	conn, err := sql.Open("sqlite", "./quiz.db")
	if err != nil {
		return []db.User{}, err
	}

	q := db.New(conn)
	us, err := q.GetUsers(context.Background())
	if err != nil {
		return []db.User{}, err
	}
	return us, err
}
func GetUserByID(id int64) (db.User, error) {
	conn, err := sql.Open("sqlite", "./quiz.db")
	if err != nil {
		return db.User{}, err
	}

	q := db.New(conn)
	u, err := q.GetUserByID(context.Background(), id)
	if err != nil {
		return db.User{}, err
	}
	return u, nil
}

func CreateUser(u db.User) (db.User, error) {
	conn, err := sql.Open("sqlite", "./quiz.db")
	if err != nil {
		return db.User{}, err
	}

	q := db.New(conn)
	usr, err := q.CreateUser(context.Background(), db.CreateUserParams{
		Username: u.Username,
		Email:    u.Email,
	})

	if err != nil {
		return db.User{}, err
	}

	return usr, nil
}
