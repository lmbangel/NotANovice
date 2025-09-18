package attempt

import "context"

type AttemptService struct {
	repo AttemptRepository
}

func NewAttemptService(repo AttemptRepository) *AttemptService {
	return &AttemptService{repo: repo}
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
	return s.repo.CreateAttempt(ctx, params)
}
