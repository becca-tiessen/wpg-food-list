package app

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type RestaurantRepo interface {
	GetTopRestaurants() ([]Restaurant, error)
}

type restaurantRepo struct {
	db *sqlx.DB
}

func NewRestaurantRepo(db *sqlx.DB) RestaurantRepo {
	return &restaurantRepo{db: db}
}

func (r *restaurantRepo) GetTopRestaurants() ([]Restaurant, error) {
	return []Restaurant{
		{ID: 1, Name: "Sous Sol"},
		{ID: 2, Name: "Deer + Almond"},
		{ID: 3, Name: "Vera"},
	}, nil
}
