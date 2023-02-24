package handlers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

// main method
func main() {
	//sql connection
	db, err = sql.Open("mysql", "root:admin@123@tcp(127.0.0.1:3306)/intg_stat")
	if err != nil {
		panic(err.Error())
	}
	//functions for apis//
	defer db.Close()
}
