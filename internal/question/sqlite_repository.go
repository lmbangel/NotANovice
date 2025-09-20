package question

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lmbangel/_novice/internal/db"
	"github.com/lmbangel/_novice/pkg/agents"
)

type sqliteQuestionRepository struct {
	db *sql.DB
}

func NewSQLiteQuestionRepository(db *sql.DB) QuestionRepository {
	return &sqliteQuestionRepository{db: db}
}

func (r *sqliteQuestionRepository) GenerateQuestion(ctx context.Context) (*Question, error) {
	q := db.New(r.db)

	sysPrompt := `You are a discipleship leader (a shepherd).
					You want your student to know the foundation and most important things of the faith of Christianity as commented on by the writer of the book of Hebrews (Hebrews 5:12-14).
					Therefore, to test daily if your sheep/disciples are spending time learning and growing in the faith by particularly reading their Bible; have a quiz daily that is used to test the discipled for the day.
					It's a straightforward biblical question every day. Either part of a scripture, or question on context of verse, or question about naming a verse that talks about a specific topic in a specific kind of context.
					For example, How many times is Jesus recorded to be crying in the Bible.`

	userPrompt := "What is the quiz of the day ?"

	o := agents.Ollama{
		Url: "http://localhost:11434/api/generate",
		Request: &agents.Request{
			Model:  "llama3.2:latest",
			Prompt: fmt.Sprintf("%s %s", sysPrompt, userPrompt),
		},
	}
	response := o.Prompt()

	qsn, _ := q.CreateQuestion(ctx, db.CreateQuestionParams{
		Question:      response.Response,
		CorrectAnswer: response.Response,
		AAnswer:       response.Response,
		BAnswer:       response.Response,
		CAnswer:       response.Response,
		DAnswer:       sql.NullString{},
	})

	return fmtQuesion(qsn), nil
}
