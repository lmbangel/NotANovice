package quiz

import (
	"context"
	"database/sql"

	"github.com/lmbangel/_novice/internal/db"
)

type sqliteQuizRepository struct {
	db *sql.DB
}

func NewSQLiteQuizRepository(db *sql.DB) QuizRepository {
	return &sqliteQuizRepository{db: db}
}

func (r *sqliteQuizRepository) GetQuizes(ctx context.Context) ([]Quiz, error) {
	q := db.New(r.db)

	qzs, err := q.GetQuizes(ctx)
	if err != nil {
		return nil, err
	}
	return fmtQuizes(qzs), nil
}

func (r *sqliteQuizRepository) GetQuizByID(ctx context.Context, id int64) (*Quiz, error) {
	q := db.New(r.db)
	qz, err := q.GetQuizByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return fmtQuiz(qz), nil
}

func (r *sqliteQuizRepository) GetQuizOfTheDay(ctx context.Context) (*Quiz, error) {
	q := db.New(r.db)
	qz, err := q.GetQuizOfTheDay(ctx)
	if err != nil {
		return nil, err
	}
	return fmtQuiz(qz), nil
}
func (r *sqliteQuizRepository) CreateNewQuiz(ctx context.Context, params CreateNewQuizParams) (*Quiz, error) {
	q := db.New(r.db)

	qz, err := q.CreateNewQuiz(ctx, db.CreateNewQuizParams{QID: params.QID, AID: params.AID})
	if err != nil {
		return nil, err
	}
	return fmtQuiz(qz), nil
}
