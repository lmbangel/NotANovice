package leaderboard

import (
	"context"
	"database/sql"

	"github.com/lmbangel/_novice/internal/db"
)

type sqliteLeaderBoardRepository struct {
	db *sql.DB
}

func NewSQLiteLeaderBoardRepository(db *sql.DB) LeaderBoardRepository {
	return &sqliteLeaderBoardRepository{db: db}
}

func (r *sqliteLeaderBoardRepository) GetLeaderBoard(ctx context.Context) ([]LeaderBoard, error) {
	q := db.New(r.db)

	leaderBoards, err := q.GetLeaderBoard(ctx)
	if err != nil {
		return nil, err
	}
	return fmtLeaderBoardList(leaderBoards), nil
}

func (r *sqliteLeaderBoardRepository) GetLeaderBoardByUserID(ctx context.Context, userID int64) (*LeaderBoard, error) {
	q := db.New(r.db)

	leaderBoard, err := q.GetLeaderBoardByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return fmtLeaderBoard(leaderBoard), nil
}
