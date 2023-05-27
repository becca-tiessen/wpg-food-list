package app

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func MakeSubmitVoteEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(submitVoteRequest)
		if !ok {
			return submitVoteResponse{}, errors.New("Invalid Request")
		}

		err := svc.SubmitVote(req.Vote)
		if err != nil {
			return nil, err
		}
		return submitVoteResponse{}, nil
	}
}

// instead of this, maybe GetBestOf(category) and then pass in 'alltime', 'brunch', 'dessert', 'pizza'
func MakeGetTopRestaurantEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		restaurants, err := svc.GetTopRestaurants()
		if err != nil {
			return nil, err
		}
		return getTopRestaurantsResponse{restaurants}, nil
	}
}

type getTopRestaurantsResponse struct {
	Restaurants []Restaurant `json:"restaurants"`
}

type submitVoteRequest struct {
	Vote Vote `json:"vote"`
}

type submitVoteResponse struct{}
