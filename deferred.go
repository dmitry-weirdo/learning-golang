package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("main 1")

	// will defer to after the execution of main
	defer fmt.Println("deferred 1")

	fmt.Println("main 2")

	// deferred functions are LIFO, so "deferred 2" will be executed before "deferred 1"
	defer fmt.Println("deferred 2")
}

func dbExecute() {
	// we define close near the resource opening
	// then, since defer functions are LIFO, first `rows` and then `db` will be closed.
	db, _ := sql.Open("driverName", "connection string")
	defer db.Close() // can itself produce an error

	rows, _ := db.Query("some sql query")
	defer rows.Close() // can itself produce an error
}
