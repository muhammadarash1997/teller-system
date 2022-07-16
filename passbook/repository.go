package passbook

import (
	"gorm.io/gorm"
	"teller-system/customer"
)

type Repository interface {
	FindByAccountId(uint, string, string) (Passbooks, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByAccountId(accountId uint, startDate string, endDate string) (Passbooks, error) {
	var passbooks Passbooks
	err := r.db.Raw("SELECT transaction_date, description, CASE WHEN debit_credit_status = 'C' THEN amount END AS credit, CASE WHEN debit_credit_status = 'D' THEN amount END AS debit FROM transactions WHERE account_id = ? AND transaction_date BETWEEN ? AND ? ORDER BY transaction_date", accountId, startDate, endDate).Scan(&passbooks).Error
	if err != nil {
		return passbooks, err
	}

	return passbooks, nil
}

func (r *repository) FindAll() (customer.Customers, error) {
	var customers customer.Customers
	err := r.db.Find(&customers).Error
	if err != nil {
		return customers, err
	}

	return customers, nil
}
