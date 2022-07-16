package main

import (
	"net/http"
	"teller-system/customer"
	"teller-system/database"
	_ "teller-system/docs" // This line is necessary for go-swagger to find your docs!
	"teller-system/passbook"
	"teller-system/transaction"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db := database.StartConnection()

	customerRepository := customer.NewRepository(db)
	customerService := customer.NewService(customerRepository)
	customerController := customer.NewController(customerService)

	transactionRepository := transaction.NewRepository(db)
	transactionService := transaction.NewService(transactionRepository)
	transactionController := transaction.NewController(transactionService)

	passbookRepository := passbook.NewRepository(db)
	passbookService := passbook.NewService(passbookRepository)
	passbookController := passbook.NewController(passbookService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/api/docs", "./swagger")

	// swagger:route GET /api/test test testAPI
	// Test API
	//
	// responses:
	//		200: testAPI

	// For API Testing
	router.GET("/api/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	router.POST("/api/customer", customerController.InputCustomerDataHandler)
	router.GET("/api/point", customerController.FindAllPointsHandler)
	router.POST("/api/transaction", transactionController.InputTransaction)
	router.GET("/api/passbook/:account_id/:start_date/:end_date", passbookController.FindPassbooksByAccountIdHandler)

	router.Run()
}
