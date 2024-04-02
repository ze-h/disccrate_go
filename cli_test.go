package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSQLConnection(t *testing.T) {
	body, err := os.ReadFile("./db.cfg")
	if err != nil {
		fmt.Println("db.cfg not found in local directory.")
		t.Fatal(err)
	}
	db, err := sql.Open("mysql", string(body))
	if err != nil {
		t.Fatal(err)
	}
	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var pswd string
		if err := rows.Scan(&pswd); err != nil {
			t.Fatal(err)
		}
		fmt.Println(pswd)
	}
}