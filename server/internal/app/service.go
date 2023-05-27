package app

type Service interface {
	SubmitVote(v Vote) (err error)
	GetTopRestaurants() ([]Restaurant, error)
}

type service struct {
	voteRepo       VoteRepo
	restaurantRepo RestaurantRepo
}

func NewService(voteRepo VoteRepo, restaurantRepo RestaurantRepo) Service {
	return &service{
		voteRepo:       voteRepo,
		restaurantRepo: restaurantRepo,
	}
}

func (s *service) SubmitVote(v Vote) error {
	return s.voteRepo.SubmitVote(v)
}

func (s *service) GetTopRestaurants() ([]Restaurant, error) {
	return s.restaurantRepo.GetTopRestaurants()
}
