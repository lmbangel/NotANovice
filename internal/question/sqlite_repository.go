package question

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"os"
	"strings"

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
	data, err := os.ReadFile("./internal/question/_prompt.md")
	if err != nil {
		return nil, err
	}

	sysPrompt := string(data)
	o := agents.Ollama{
		Url: "http://localhost:11434/api/generate",
		Request: &agents.Request{
			Model:  "llama3.2:latest",
			Prompt: sysPrompt,
		},
	}
	response := o.Prompt()

	//resp := strings.TrimPrefix(strings.TrimSpace(response.Response), "```json")
	//resp = strings.TrimPrefix(resp, "```")
	//resp = strings.TrimSuffix(resp, "```")

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
