package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// runner for the program's continuous kiosk interface
func kiosk(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	usr, pass, err := login(reader)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	check, err := verify(db, usr, pass)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	if !check {
		fmt.Printf("Incorrect username or password\n")
		return
	}

	fmt.Printf("Welcome, %s.\n", usr)
	for kisok_loop(usr, reader, db) {
	}
	fmt.Println("Goodbye!")
}

// kiosk scanner loop
func kisok_loop(username string, reader *bufio.Reader, db *sql.DB) bool {
	var album [8]string
	album, err := promptUPC(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = addRecord(db, album, username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return true
}
