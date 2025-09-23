package leaderboard

import (
	"context"
	"time"
)

type LeaderBoardRepository interface {
	GetLeaderBoard(ctx context.Context) ([]LeaderBoard, error)
	GetLeaderBoardByUserID(ctx context.Context, u_id int64) (*LeaderBoard, error)
	//UpdateUserScore(ctx context.Context, u_id int64) (*LeaderBoard, error)
}

type LeaderBoard struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	TotalScore  int64     `json:"total_score"`
	LastUpdated time.Time `json:"last_updated"`
}
