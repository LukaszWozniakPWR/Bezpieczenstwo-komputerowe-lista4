package database

import "database/sql"
import (
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var database *sql.DB

func Initialize() {
	var err error
	database, err = sql.Open("sqlite3", "./database.db")
	errorCheck(err, func(err error) { log.Fatal("Could not open database") })
}

func RecreateDatabase() {
	deleteTables()
	createTables()
}

func deleteTables() {
	deleteUserCredetialsTable()
	deleteTransactionsTable()
}

func createTables() {
	createUserCredentialsTable()
	createTransactionsTable()
}

func executeStatement(query string, args ...interface{}) {
	statement, err := database.Prepare(query)
	errorCheck(err, func(err error) { log.Fatal("Error executing statement ", query, "\n", err) })
	statement.Exec(args...)
}

func queryDatabase(query string) *sql.Rows {
	rows, err := database.Query(query)
	errorCheck(err, func(err error) { log.Fatal("Error in database query ", query, "\n", err) })
	return rows
}

func errorCheck(err error, function func(err error)) {
	if err != nil {
		function(err)
	}
}
