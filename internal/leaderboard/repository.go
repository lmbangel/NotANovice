package leaderboard

import (
	"context"
	"time"

	"github.com/lmbangel/_novice/internal/db"
)

type LeaderBoardRepository interface {
	GetLeaderBoard(ctx context.Context) ([]LeaderBoard, error)
	GetLeaderBoardByUserID(ctx context.Context, u_id int64) (*LeaderBoard, error)
	//UpdateUserScore(ctx context.Context, u_id int64) (*LeaderBoard, error)
}

type LeaderBoard struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	TotalScore  int64     `json:"total_score"`
	LastUpdated time.Time `json:"last_updated"`
}

func fmtLeaderBoardList(l_board []db.GetLeaderBoardRow) []LeaderBoard {
	newLeaderBoard := make([]LeaderBoard, len(l_board))

	for i, l := range l_board {
		newLeaderBoard[i] = LeaderBoard{
			Username:    l.Username.String,
			Email:       l.Email.String,
			ID:          l.ID,
			UserID:      l.UserID,
			TotalScore:  l.TotalScore.Int64,
			LastUpdated: l.LastUpdated.Time,
		}
	}
	return newLeaderBoard
}
func fmtLeaderBoard(l db.GetLeaderBoardByUserIDRow) *LeaderBoard {
	return &LeaderBoard{
		Username:    l.Email.String,
		Email:       l.Email.String,
		ID:          l.ID,
		UserID:      l.UserID,
		TotalScore:  l.TotalScore.Int64,
		LastUpdated: l.LastUpdated.Time,
	}
}
