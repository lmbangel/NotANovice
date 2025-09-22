package attempt

import (
	"context"
	"database/sql"

	"github.com/lmbangel/_novice/internal/quiz"
)

type AttemptService struct {
	repo    AttemptRepository
	quizzes *quiz.QuizService
}

func NewAttemptService(repo AttemptRepository, quizzes *quiz.QuizService) *AttemptService {
	return &AttemptService{repo: repo, quizzes: quizzes}
}

func (s AttemptService) GetAttempts(ctx context.Context) ([]Attempt, error) {
	return s.repo.GetAttempts(ctx)
}
func (s AttemptService) GetAttemptByID(ctx context.Context, id int64) (*Attempt, error) {
	return s.repo.GetAttemptByID(ctx, id)
}

func (s AttemptService) GetAttemptsByUserID(ctx context.Context, u_id int64) ([]Attempt, error) {
	return s.repo.GetAttemptsByUserID(ctx, u_id)
}

func (s AttemptService) CreateAttempt(ctx context.Context, params CreateAttemptParams) (*Attempt, error) {
	quizQuestion, err := s.quizzes.GetQuizByID(ctx, params.QuizID)
	if err != nil {
		return nil, err
	}
	if quizQuestion.CorrectAnswer == params.Answer {
		params.IsCorrect = sql.NullBool{
			Bool:  true,
			Valid: true,
		}
	}
	return s.repo.CreateAttempt(ctx, params)
}
