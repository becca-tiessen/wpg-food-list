package app

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func MakeCreateListingEndpoint(svc Service) endpoint.Endpoint {
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

type submitVoteRequest struct {
	Vote Vote `json:"vote"`
}

type submitVoteResponse struct{}
