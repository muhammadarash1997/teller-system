package database

import (
	"fmt"
	"os"
	"teller-system/customer"
	"teller-system/transaction"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartConnection() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Fail connect")
		panic(err.Error())
	}
	fmt.Println("Success connect")

	db.AutoMigrate(customer.Customer{})
	db.AutoMigrate(transaction.Transaction{})

	return db
}
