package main

import (
	"bufio"
	"crypto/md5"
	"database/sql"
	"fmt"
	"os"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/term"
	//"github.com/irlndts/go-discogs"
)

func main() {
	db_cfg, err := readConfig("db.cfg")
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := sql.Open("mysql", getVar("DSN", db_cfg))
	if err != nil {
		fmt.Println(err)
		return
	}

	usr, pass, err := login()

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	rows, err := db.Query("SELECT password FROM users WHERE username = ?", usr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var pswd string

	for rows.Next() {
		if err := rows.Scan(&pswd); err != nil {
			fmt.Println(err)
			return
		}
	}

	if pswd != pass {
		fmt.Println("Incorrect username or password.")
		return
	}

	fmt.Printf("Welcome, %s.\n", usr)
}

// prompt uname and pass
func login() (string, string, error) {
	var err error
	var uname string
	var password string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username\n>")
	uname, err = reader.ReadString('\n')
	if err != nil {
		return "", password, err
	}

	fmt.Print("Enter your password\n>")
	bword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", password, err
	}
	password = fmt.Sprintf("%x", md5.Sum(bword))
	fmt.Println()
	return stripFormatting(uname), password, nil
}
