package main

import (
	"bufio"
	"crypto/md5"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/term"
	//"github.com/irlndts/go-discogs"
)

func main() {
	body, err := os.ReadFile("./db.cfg")
	if err != nil {
		fmt.Println("db.cfg not found in local directory.")
		return
	}
	_, err = sql.Open("mysql", string(body))
	if err != nil {
		fmt.Println(err)
		return
	}

	usr, pass, err := login()

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Println(usr)
	fmt.Printf("0x%x\n", pass)

}

// login loop, prompt uname and pass
func login() (string, [16]byte, error) {
	var err error
	var uname string
	var password [16]byte
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
	password = md5.Sum(bword)
	fmt.Println()
	return stripFormatting(uname), password, nil
}

// remove lf and cr from string
func stripFormatting(str string) string {
	newStr := str
	newStr = strings.ReplaceAll(newStr, "\n", "")
	newStr = strings.ReplaceAll(newStr, "\r", "")
	return newStr
}
