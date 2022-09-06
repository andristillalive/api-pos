package logincontroller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"pos-backend/config"
	"pos-backend/functions/loginquery"
	"pos-backend/models"
)

var db *sql.DB

func ReadAllUser() models.ResponseUserAccounts {
	// Build connection string
	connString := config.ConnString()

	var err error
	var users []models.UserAccount
	var result models.ResponseUserAccounts

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		result.Message = err.Error()
		result.UserAccounts = users
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
		result.Message = err.Error()
		result.UserAccounts = users
	}
	fmt.Printf("Connected!\n")

	// Read employees
	err, users = loginquery.QueryReadAllUser(db, ctx)
	if err != nil {
		log.Fatal("Error reading Employees: ", err.Error())
		result.Message = err.Error()
		result.UserAccounts = users
	}
	fmt.Printf("Read %d row(s) successfully.\n", len(users))

	result.Message = ""
	result.UserAccounts = users
	return result
}

func Login(username string, password string) models.ResponseUserAccount {
	// Build connection string
	connString := config.ConnString()

	var err error
	var user models.UserAccount
	var result models.ResponseUserAccount

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		result.Message = err.Error()
		result.UserAccount = user
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
		result.Message = err.Error()
		result.UserAccount = user
	}
	fmt.Printf("Connected!\n")

	// Read login
	msg, user1 := loginquery.QueryUserLogin(username, password, db, ctx)
	if err != nil {
		log.Fatal("Login Failed : ", err.Error())
		result.Message = msg
		result.UserAccount = user1
		fmt.Printf(msg)
	} else {
		result.UserAccount = user1
		fmt.Printf("Login with username : %s and nik : %s successfully.\n", user1.Nama, user1.UserNumber)
	}
	return result
}
