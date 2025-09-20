package question

import "context"

type QuestionService struct {
	repo QuestionRepository
}

func NewQuestionService(repo QuestionRepository) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) GenerateQuestion(ctx context.Context) (*Question, error) {
	q, err := s.repo.GenerateQuestion(ctx)
	if err != nil {
		return nil, err
	}
	return q, nil
}
