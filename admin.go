package main

import (
	"bufio"
	"crypto/md5"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//runner for cli admin settings
func admin() {
	db_cfg, err := readConfig("db.cfg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("mysql", getVar("DSN", db_cfg))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Administrator mode active.")

	for adminMode(reader, db) {

	}
}

// adminmode loop
func adminMode(reader *bufio.Reader, db *sql.DB) bool {
	fmt.Print("Select an option:\n1 - Add user to database\n2 - Query users\n3 - Remove user from database\n0 - Exit\n>")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch stripFormatting(input) {
	case "1":
		fmt.Print("Enter the new user's name\n>")
		uname, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		uname = stripFormatting(uname)
		fmt.Print("Enter the new user's password\n>")
		pass, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		pass = stripFormatting(pass)
		_, err = db.Query("INSERT INTO users (username, password) VALUES (?, ?)", uname, fmt.Sprintf("%x", md5.Sum([]byte(pass))))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return true
	case "2":
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer rows.Close()
		fmt.Println("user\tpassword hash")
		for rows.Next() {
			var uname, pass string
			if err := rows.Scan(&uname, &pass); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("%s\t%s\n", uname, pass)
		}
		fmt.Println()
		return true
	case "3":
		fmt.Print("Enter the user's name\n>")
		uname, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		uname = stripFormatting(uname)
		_, err = db.Query("DELETE FROM records WHERE username = ?", uname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = db.Query("DELETE FROM users WHERE username = ?", uname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("User %s deleted successfully.\n", uname)
		return true
	case "0":
		return false
	}
	return true
}
