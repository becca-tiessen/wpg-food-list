package app

type Service interface {
	SubmitVote(v Vote) (err error)
}

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{
		repo: repo,
	}
}
func (s *service) SubmitVote(v Vote) error {
	return s.repo.SubmitVote(v)
}
