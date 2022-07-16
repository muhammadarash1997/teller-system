// Package classification Teller System API.
//
// This is a API for teller in charge of customer data and customer financial transactions
//
//   Schemes: http
//   Host: localhost:8080
//   BasePath: /
//   Version: 1.0.0
//   Contact: muhammadarash1997@gmail.com
//
//   Consumes:
//   - application/json
//
//   Produces:
//   - application/json
//
// swagger:meta
package docs

import (
	"teller-system/customer"
	"teller-system/passbook"
	"teller-system/transaction"
)

// Success testing API
// swagger:response testAPI
type testAPI struct {
	// in: Body
	Body struct {
		Message string `json:"message"`
	}
}

// swagger:response messageResponse
type messageResponse struct {
	// in: Body
	Body struct {
		Message string `json:"message"`
	}
}

// swagger:response inputCustomerData
type inputCustomerData struct {
	// in: Body
	Body customer.Customer
}

// swagger:response findAllPoints
type findAllPoints struct {
	// in: Body
	Body customer.Customers
}

// swagger:response findPassbooksCustomer
type findPassbooksCustomer struct {
	// in: Body
	Body passbook.Passbooks
}

// swagger:parameters inputCustomerData
type inputCustomerDataParams struct {
	// Customer account_id and customer name that needs to be registered
	// in: body
	// required: true
	Body struct {
		AccountId uint    `json:"account_id"`
		Name      string `json:"name"`
	}
}

// swagger:parameters inputTransaction
type inputTransactionParams struct {
	// Transaction object that needs to be stored
	// in: body
	// required: true
	Body transaction.Transaction
}

// swagger:parameters findPassbooks
type findPassbooksParams struct {
	// The account id for as customer parameter
	// in: path
	// required: true
	AccountId uint `json:"account_id"`
	// Start date for as parameter, ex = 2022-12-31
	// in: path
	// required: true
	StartDate string `json:"start_date"`
	// End date for as parameter, ex = 2022-12-31
	// in: path
	// required: true
	EndDate   string `json:"end_date"`
}
