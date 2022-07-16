package passbook

import "time"

// Domain
type Passbook struct {
	TransactionDate time.Time `json:"transaction_date"`
	Description     string    `json:"description"`
	Credit          int64     `json:"credit"`
	Debit           int64     `json:"debit"`
	Amount          int64     `json:"amount"`
}

type Passbooks []Passbook
