package passbook

type Service interface {
	FindByAccountId(accountId uint, startDate string, endDate string) (Passbooks, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindByAccountId(accountId uint, startDate string, endDate string) (Passbooks, error) {
	return s.repository.FindByAccountId(accountId, startDate, endDate)
}
