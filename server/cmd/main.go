package main

import (
	"context"
	"encoding/json"
	"net/http"

	"wpg-food-list/server/internal/app"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("mysql", "root:password@tcp(127.0.0.1:3318)/wpg-food-list?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	vRepo := app.NewVoteRepo(db)
	rRepo := app.NewRestaurantRepo(db)

	svc := app.NewService(vRepo, rRepo)

	voteHandler := httptransport.NewServer(
		app.MakeSubmitVoteEndpoint(svc),
		decodeRequest,
		encodeResponse,
	)

	restaurantHandler := httptransport.NewServer(
		app.MakeGetTopRestaurantEndpoint(svc),
		decodeRequest,
		encodeResponse,
	)

	http.Handle("/vote", voteHandler)
	http.Handle("/restaurants", restaurantHandler)

	http.ListenAndServe(":8080", nil)
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
