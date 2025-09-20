package question

import (
	"context"
	"database/sql"

	"github.com/lmbangel/_novice/internal/db"
)

type QuestionRepository interface {
	GenerateQuestion(ctx context.Context) (*Question, error)
}

type Question struct {
	ID            int64          `json:"id"`
	Question      string         `json:"question"`
	CorrectAnswer string         `json:"correct_answer"`
	Timestamp     sql.NullTime   `json:"timestamp"`
	IsActive      sql.NullBool   `json:"is_active"`
	AAnswer       string         `json:"a_answer"`
	BAnswer       string         `json:"b_answer"`
	CAnswer       string         `json:"c_answer"`
	DAnswer       sql.NullString `json:"d_answer"`
}

func fmtQuesion(q db.Question) *Question {
	return &Question{
		ID:            q.ID,
		Question:      q.Question,
		CorrectAnswer: q.CorrectAnswer,
		Timestamp:     q.Timestamp,
		IsActive:      q.IsActive,
		AAnswer:       q.AAnswer,
		BAnswer:       q.BAnswer,
		CAnswer:       q.CAnswer,
		DAnswer:       q.DAnswer,
	}
}
