package app

import (
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type Repo interface {
	SubmitVote(v Vote) error
}

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) Repo {
	return &repo{db: db}
}

func (r *repo) SubmitVote(v Vote) error {
	now := time.Now()
	sql, args, err := sq.Insert("Votes").
		SetMap(sq.Eq{
			"user_id":       v.UserID,
			"restaurant_id": v.RestaurantID,
			"factor_id":     v.FactorID,
			"factor_rating": v.FactorRating,
			"created_at":    now,
		}).
		ToSql()
	if err != nil {
		return errors.New("error building submit vote query")
	}

	_, err = r.db.Exec(sql, args...)
	if err != nil {
		return errors.New("error submitting vote")
	}

	return nil
}
