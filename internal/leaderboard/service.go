package leaderboard

import "context"

type LeaderBoardService struct {
	repo LeaderBoardRepository
}

func NewLeaderBoardService(repo LeaderBoardRepository) *LeaderBoardService {
	return &LeaderBoardService{repo: repo}
}

func (s *LeaderBoardService) GetLeaderBoard(ctx context.Context) ([]LeaderBoard, error) {

	leaderBoard, err := s.repo.GetLeaderBoard(ctx)
	if err != nil {
		return nil, err
	}
	return leaderBoard, nil
}
func (s *LeaderBoardService) GetLeaderBoardByUserID(ctx context.Context, userID int64) (*LeaderBoard, error) {

	leaderBoard, err := s.repo.GetLeaderBoardByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return leaderBoard, nil
}
