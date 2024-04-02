package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSQLConnection(t *testing.T) {
	DB_CONFIG = "db.cfg"

	db_cfg, err := readConfig(DB_CONFIG)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("mysql", getVar("DSN", db_cfg))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
