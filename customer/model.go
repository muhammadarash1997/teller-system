package customer

// Entity
type Customer struct {
	AccountId  int    `json:"account_id"`
	Name       string `json:"name"`
	TotalPoint int64    `json:"total_point"`
}

type Customers []Customer