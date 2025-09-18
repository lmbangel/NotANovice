package attempt

import (
	"context"
	"database/sql"

	"github.com/lmbangel/_novice/internal/db"
)

type sqliteAttemptRepository struct {
	db *sql.DB
}

func NewSQLiteAttemptRepository(db *sql.DB) AttemptRepository {
	return &sqliteAttemptRepository{db: db}
}

func (r *sqliteAttemptRepository) CreateAttempt(ctx context.Context, params CreateAttemptParams) (*Attempt, error) {
	q := db.New(r.db)

	a, err := q.RecordAttempt(ctx, db.RecordAttemptParams(params))
	if err != nil {
		return nil, err
	}

	return fmtAttempt(&a), nil
}

func (r *sqliteAttemptRepository) GetAttempts(ctx context.Context) ([]Attempt, error) {
	q := db.New(r.db)

	a, err := q.GetAttempts(ctx)
	if err != nil {
		return nil, err
	}

	return fmtAttempts(a), nil
}

func (r *sqliteAttemptRepository) GetAttemptByID(ctx context.Context, id int64) (*Attempt, error) {
	q := db.New(r.db)

	a, err := q.GetAttemptByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return fmtAttempt(&a), nil
}

func (r *sqliteAttemptRepository) GetAttemptsByUserID(ctx context.Context, u_id int64) ([]Attempt, error) {
	q := db.New(r.db)

	a, err := q.GetAttemptsByUserID(ctx, u_id)
	if err != nil {
		return nil, err
	}
	return fmtAttempts(a), nil
}
