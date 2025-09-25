package question

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/lmbangel/_novice/internal/db"
	"github.com/lmbangel/_novice/internal/quiz"
	"github.com/lmbangel/_novice/pkg/agents"
)

type sqliteQuestionRepository struct {
	db *sql.DB
}

func NewSQLiteQuestionRepository(db *sql.DB) QuestionRepository {
	return &sqliteQuestionRepository{db: db}
}

func (r *sqliteQuestionRepository) GetQuestions(ctx context.Context) ([]Question, error) {
	q := db.New(r.db)

	qs, err := q.GetQuestions(ctx)
	if err != nil {
		return nil, err
	}
	return fmtQuestions(qs), nil
}

func (r *sqliteQuestionRepository) GenerateQuestion(ctx context.Context) (*Question, error) {
	data, err := os.ReadFile("./internal/question/_prompt.md")
	if err != nil {
		return nil, err
	}

	sysPrompt := string(data)
	o := agents.Ollama{
		Url: os.Getenv("OLLAMA_URL") + "/api/generate",
		Request: &agents.Request{
			Model:  "llama3.2:latest",
			Prompt: sysPrompt,
			//Options: &agents.Options{
			//	Temperature: 0,
			//	TopP:        1,
			//	TopK:        1,
			//},
		},
	}
	response := o.Prompt()

	var newQuestion Question
	if err := ExtractJSON(response.Response, &newQuestion); err != nil {
		return nil, err
	}
	if newQuestion.CorrectAnswer == "" {
		return nil, errors.New("LLM returned blank correct_answer")
	}

	q := db.New(r.db)
	qsn, _ := q.CreateQuestion(ctx, db.CreateQuestionParams{
		Question:      newQuestion.Question,
		CorrectAnswer: newQuestion.CorrectAnswer,
		AAnswer:       newQuestion.AAnswer,
		BAnswer:       newQuestion.BAnswer,
		CAnswer:       newQuestion.CAnswer,
		DAnswer:       newQuestion.DAnswer,
	})

	quizRepo := quiz.NewSQLiteQuizRepository(r.db)
	if _, err := quizRepo.CreateNewQuiz(ctx, quiz.CreateNewQuizParams{QID: qsn.ID, AID: qsn.ID}); err != nil {
		return nil, err
	}

	return fmtQuesion(qsn), nil
}

func ExtractJSON(resp string, v interface{}) error {
	start := strings.Index(resp, "{")
	end := strings.LastIndex(resp, "}")
	if start == -1 || end == -1 || start >= end {
		return errors.New("no valid JSON object found")
	}
	jsonStr := resp[start : end+1]
	return json.Unmarshal([]byte(jsonStr), v)
}
