package customer

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(Customer) (Customer, error)
	FindAll() (Customers, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(customer Customer) (Customer, error) {
	err := r.db.Create(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) FindAll() (Customers, error) {
	var customers Customers
	err := r.db.Find(&customers).Error
	if err != nil {
		return customers, err
	}

	return customers, nil
}
