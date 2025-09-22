package quiz

import (
	"context"
	"time"

	"github.com/lmbangel/_novice/internal/db"
)

type QuizRepository interface {
	GetQuizes(ctx context.Context) ([]Quiz, error)
	GetQuizByID(ctx context.Context, id int64) (*Quiz, error)
	GetQuizOfTheDay(ctx context.Context) (*Quiz, error)
}

type Quiz struct {
	QuizID        int64     `json:"quiz_id"`
	ID            int64     `json:"id"`
	Question      string    `json:"question"`
	CorrectAnswer string    `json:"correct_answer"`
	Timestamp     time.Time `json:"timestamp"`
	IsActive      bool      `json:"is_active"`
	AAnswer       string    `json:"a_answer"`
	BAnswer       string    `json:"b_answer"`
	CAnswer       string    `json:"c_answer"`
	DAnswer       string    `json:"d_answer"`
}

func fmtQuizes(q []db.GetQuizesRow) []Quiz {
	qzs := make([]Quiz, len(q))

	for i, qui := range q {
		qzs[i] = Quiz{
			QuizID:        qui.QuizID,
			ID:            qui.ID.Int64,
			Question:      qui.Question.String,
			CorrectAnswer: qui.CorrectAnswer.String,
			Timestamp:     qui.Timestamp.Time,
			IsActive:      qui.IsActive.Bool,
			AAnswer:       qui.AAnswer.String,
			BAnswer:       qui.BAnswer.String,
			CAnswer:       qui.CAnswer.String,
			DAnswer:       qui.DAnswer.String,
		}
	}
	return qzs
}

//func fmtQuiz(qui db.GetQuizOfTheDayRow) *Quiz {
//	return &Quiz{
//		QuizID:        qui.QuizID,
//		ID:            qui.ID,
//		Question:      qui.Question,
//		CorrectAnswer: qui.CorrectAnswer,
//		Timestamp:     qui.Timestamp,
//		IsActive:      qui.IsActive,
//		AAnswer:       qui.AAnswer,
//		BAnswer:       qui.BAnswer,
//		CAnswer:       qui.CAnswer,
//		DAnswer:       qui.DAnswer,
//	}
//}

func fmtQuiz(qui any) *Quiz {
	switch q := qui.(type) {
	case db.GetQuizOfTheDayRow:
		return &Quiz{
			QuizID:        q.QuizID,
			ID:            q.ID.Int64,
			Question:      q.Question.String,
			CorrectAnswer: q.CorrectAnswer.String,
			Timestamp:     q.Timestamp.Time,
			IsActive:      q.IsActive.Bool,
			AAnswer:       q.AAnswer.String,
			BAnswer:       q.BAnswer.String,
			CAnswer:       q.CAnswer.String,
			DAnswer:       q.DAnswer.String,
		}
	case db.GetQuizesRow:
		return &Quiz{
			QuizID:        q.QuizID,
			ID:            q.ID.Int64,
			Question:      q.Question.String,
			CorrectAnswer: q.CorrectAnswer.String,
			Timestamp:     q.Timestamp.Time,
			IsActive:      q.IsActive.Bool,
			AAnswer:       q.AAnswer.String,
			BAnswer:       q.BAnswer.String,
			CAnswer:       q.CAnswer.String,
			DAnswer:       q.DAnswer.String,
		}
	case db.GetQuizByIDRow:
		return &Quiz{
			QuizID:        q.QuizID,
			ID:            q.ID.Int64,
			Question:      q.Question.String,
			CorrectAnswer: q.CorrectAnswer.String,
			Timestamp:     q.Timestamp.Time,
			IsActive:      q.IsActive.Bool,
			AAnswer:       q.AAnswer.String,
			BAnswer:       q.BAnswer.String,
			CAnswer:       q.CAnswer.String,
			DAnswer:       q.DAnswer.String,
		}
	default:
		return nil
	}
}
