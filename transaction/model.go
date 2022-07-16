package transaction

import "time"

// Entity
type Transaction struct {
	AccountId         uint      `json:"account_id"`
	TransactionDate   time.Time `json:"transaction_date"`
	Description       string    `json:"description"`
	DebitCreditStatus string    `json:"debit_credit_status"`
	Amount            int64     `json:"amount"`
}
