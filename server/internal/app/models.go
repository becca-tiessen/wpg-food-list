package app

type Restaurant struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type User struct {
	ID    int64  `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
}

type Vote struct {
	UserID       int64 `json:"user_id" db:"user_id"`
	RestaurantID int64 `json:"restaurant_id" db:"restaurant_id"`
	FactorID     int64 `json:"factor_id" db:"factor_id"`
	FactorRating int64 `json:"factor_rating" db:"factor_rating"`
}
