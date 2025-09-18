package attempt

import (
	"context"
	"database/sql"

	"github.com/lmbangel/_novice/internal/db"
)

type AttemptRepository interface {
	CreateAttempt(ctx context.Context, p CreateAttemptParams) (*Attempt, error)
	GetAttempts(ctx context.Context) ([]Attempt, error)
	GetAttemptByID(ctx context.Context, id int64) (*Attempt, error)
	GetAttemptsByUserID(ctx context.Context, u_id int64) ([]Attempt, error)
}

type CreateAttemptParams struct {
	UserID    int64        `json:"user_id"`
	QuizID    int64        `json:"quiz_id"`
	Answer    string       `json:"answer"`
	IsCorrect sql.NullBool `json:"is_correct"`
}

type Attempt struct {
	ID        int64        `json:"id"`
	UserID    int64        `json:"user_id"`
	QuizID    int64        `json:"quiz_id"`
	Answer    string       `json:"answer"`
	IsCorrect sql.NullBool `json:"is_correct"`
	Timestamp sql.NullTime `json:"timestamp"`
}

func fmtAttempt(a *db.Attempt) *Attempt {
	return &Attempt{
		ID:        a.ID,
		UserID:    a.UserID,
		QuizID:    a.QuizID,
		Answer:    a.Answer,
		IsCorrect: a.IsCorrect,
		Timestamp: a.Timestamp,
	}
}

func fmtAttempts(db_a []db.Attempt) []Attempt {
	att := make([]Attempt, len(db_a))
	for i, a := range db_a {
		att[i] = Attempt{
			ID:        a.ID,
			UserID:    a.UserID,
			QuizID:    a.QuizID,
			Answer:    a.Answer,
			IsCorrect: a.IsCorrect,
			Timestamp: a.Timestamp,
		}
	}
	return att
}
