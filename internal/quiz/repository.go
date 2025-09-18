package quiz

import (
	"context"
	"database/sql"

	"github.com/lmbangel/_novice/internal/db"
)

type QuizRepository interface {
	GetQuizes(ctx context.Context) ([]Quiz, error)
	GetQuizByID(ctx context.Context, id int64) (*Quiz, error)
	GetQuizOfTheDay(ctx context.Context) (*Quiz, error)
}

type Quiz struct {
	ID          int64        `json:"id"`
	QID         int64        `json:"q_id"`
	AID         int64        `json:"a_id"`
	Date        sql.NullTime `json:"date"`
	IsActive    sql.NullBool `json:"is_active"`
	OptionsJson string       `json:"options_json"`
}

func fmtQuizes(q []db.Quiz) []Quiz {
	qzs := make([]Quiz, len(q))

	for i, qui := range q {
		qzs[i] = Quiz{
			ID:          qui.ID,
			QID:         qui.QID,
			AID:         qui.AID,
			Date:        qui.Date,
			IsActive:    qui.IsActive,
			OptionsJson: qui.OptionsJson,
		}
	}
	return qzs
}

func fmtQuiz(qui db.Quiz) *Quiz {
	return &Quiz{
		ID:          qui.ID,
		QID:         qui.QID,
		AID:         qui.AID,
		Date:        qui.Date,
		IsActive:    qui.IsActive,
		OptionsJson: qui.OptionsJson,
	}
}
