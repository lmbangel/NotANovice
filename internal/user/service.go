package user

import "context"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers(ctx context.Context) ([]User, error) {
	return s.repo.GetUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (*User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) GetUserByUsername(ctx context.Context, u_name string) (*User, error) {
	return s.repo.GetUserByUsername(ctx, u_name)
}

func (s *UserService) CreateUser(ctx context.Context, params CreateUserParams) (*User, error) {
	return s.repo.CreateUser(ctx, params)
}

func (s UserService) UpdateUser(ctx context.Context, params UpdateUserParams) (*User, error) {
	return s.repo.UpdateUser(ctx, params)
}
