package quiz

import (
	"context"
)

type QuizService struct {
	repo QuizRepository
}

func NewQuizService(repo QuizRepository) *QuizService {
	return &QuizService{repo: repo}
}

func (s *QuizService) GetQuizes(ctx context.Context) ([]Quiz, error) {
	return s.repo.GetQuizes(ctx)
}

func (s *QuizService) GetQuizByID(ctx context.Context, id int64) (*Quiz, error) {
	return s.repo.GetQuizByID(ctx, id)
}

func (s *QuizService) GetQuizOfTheDay(ctx context.Context) (*Quiz, error) {
	return s.repo.GetQuizOfTheDay(ctx)
}