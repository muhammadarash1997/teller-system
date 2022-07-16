package customer

type Service interface {
	Save(Customer) (Customer, error)
	FindAll() (Customers, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Save(customer Customer) (Customer, error) {
	customer, err := s.repository.Save(customer)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (s *service) FindAll() (Customers, error) {
	return s.repository.FindAll()
}
