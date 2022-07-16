# Teller System
> This is a API for teller in charge of customer data and customer financial transactions.

## Usage
Create .env file in the root directory of your project. Add environment-specific variables on new lines in the form NAME=VALUE. For example:

```bash
- DB_HOST=localhost
- DB_PORT=3306
- DB_PASS=rahasia
- DB_NAME=trydatabase
```

## Requirements
You need [MySQL](https://www.mysql.com/) server running.

## Running App

```bash
go run main.go
```
You could run localhost:8080/api/docs in your browser for seeing documentation and testing this service 

## Technologies Used
- Go 1.18
- MySQL
- Gin Web Framework
- GORM
- Swagger 2.0

## Code Structure
The design contains several layers and components and very much similar to onion or clean architecture attempt.
