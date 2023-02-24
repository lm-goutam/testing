package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

// main method
func dbb() {
	//sql connection
	db, err = sql.Open("mysql", "root:admin@123@tcp(127.0.0.1:3306)/intg_stat")
	if err != nil {
		fmt.Println("Database not found", err)
		return
	}
	//functions for apis//
	defer db.Close()
}
