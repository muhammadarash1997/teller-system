package transaction

import (
	"teller-system/customer"

	"gorm.io/gorm"
)

type Repository interface {
	Save(Transaction) error
	UpdateCustomer(customer.Customer) error
	FindCustomerById(uint) (customer.Customer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(transaction Transaction) error {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateCustomer(customer customer.Customer) error {
	err := r.db.Raw("UPDATE customers SET account_id = ?, name = ?, total_point = ? WHERE account_id = ?", customer.AccountId, customer.Name, customer.TotalPoint, customer.AccountId).Scan(&customer).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindCustomerById(accountId uint) (customer.Customer, error) {
	var customer customer.Customer
	err := r.db.Raw("SELECT * FROM customers WHERE account_id = ? LIMIT 1", accountId).Scan(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}
