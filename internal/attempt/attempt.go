package attempt

import (
	"context"
	"database/sql"

	"github.com/lmbangel/_novice/internal/db"
	_ "modernc.org/sqlite"
)

func CreateAttempt(a db.Attempt) (*db.Attempt, error) {
	conn, err := sql.Open("sqlite", "./quiz.db")
	if err != nil {
		return nil, err
	}

	q := db.New(conn)
	attempt, err := q.RecordAttempt(context.Background(), db.RecordAttemptParams{
		UserID:    a.UserID,
		QuizID:    a.QuizID,
		Answer:    a.Answer,
		IsCorrect: sql.NullBool{Valid: false},
	})
	if err != nil {
		return nil, err
	}

	return &attempt, nil
}

func GetAttempts() (*[]db.Attempt, error) {
	conn, err := sql.Open("sqlite", "./quiz.db")
	if err != nil {
		return nil, err
	}
	q := db.New(conn)
	a, err := q.GetAttempts(context.Background())
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func GetAttemptByID(id int64) (*db.Attempt, error) {
	conn, err := sql.Open("sqlite", "./quiz.db")
	if err != nil {
		return nil, err
	}

	q := db.New(conn)
	a, err := q.GetAttemptByID(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
