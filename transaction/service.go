package transaction

type Service interface {
	Save(Transaction) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Save(transaction Transaction) error {
	err := s.repository.Save(transaction)
	if err != nil {
		return err
	}

	if transaction.Description == "Beli Pulsa" || transaction.Description == "Bayar Listrik" {
		if transaction.Description == "Beli Pulsa" {
			var point int64
			if 0 <= transaction.Amount {
				point += 10000/1000 * 0
			}
			if 50001 <= transaction.Amount {
				point += 10000/1000 * 1
			}
			if 100000 <= transaction.Amount {
				point += (transaction.Amount-20000)/1000 * 2
			}

			customer, err := s.repository.FindCustomerById(transaction.AccountId)
			if err != nil {
				return err
			}
			
			customer.TotalPoint += point
			err = s.repository.UpdateCustomer(customer)
			if err != nil {
				return err
			}
		}

		if transaction.Description == "Bayar Listrik" {
			var point int64
			if 0 <= transaction.Amount {
				point += 50000/2000 * 0
			}
			if 50001 <= transaction.Amount {
				point += 50000/2000 * 1
			}
			if 100000 <= transaction.Amount {
				point += (transaction.Amount-100000)/2000 * 2
			}

			customer, err := s.repository.FindCustomerById(transaction.AccountId)
			if err != nil {
				return err
			}

			customer.TotalPoint += point
			err = s.repository.UpdateCustomer(customer)
			if err != nil {
				return err
			}
		}
	}

	return nil
}