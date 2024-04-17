package main

import (
	"bufio"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/term"
)

// runner for the program's command line interface
func cli(db *sql.DB) {
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
		fmt.Printf("Password: %s\n", pass)
		return
	}

	fmt.Printf("Welcome, %s.\n", usr)
	for usermode(usr, reader, db) {
	}
	fmt.Println("Goodbye!")
}

// usermode loop
func usermode(username string, reader *bufio.Reader, db *sql.DB) bool {
	fmt.Print("Select an option:\n1 - Add record to collection\n2 - Remove from collection\n3 - Query collection\n4 - Export collection\n0 - Exit\n>")
	input, err := reader.ReadString('\n')
	iferr(err)
	switch stripFormatting(input) {
	case "1":
		fmt.Print("Enter 1 for manual, 2 for automatic\n>")
		input_2, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var album [8]string
		if stripFormatting(input_2) == "1" {
			album, err = promptAlbum(reader)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			album, err = promptUPC(reader)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		_, err = addRecord(db, album, username)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return true
	case "2":
		fmt.Print("Enter the UPC to remove\n>")
		upc, err := reader.ReadString('\n')
		iferr(err)
		upc = stripFormatting(upc)

		title, err := findRecord(db, upc, username)

		iferr(err)

		if title == "" {
			fmt.Printf("No record with UPC %s found.\n", upc)
			return true
		}

		fmt.Printf("Are you sure you want to remove %s from your collection? (y/N)\n>", title)
		confirm, err := reader.ReadString('\n')
		iferr(err)

		confirm = strings.ToUpper(stripFormatting(confirm))

		if confirm == "Y" {
			_, err = deleteRecord(db, upc, username)
			iferr(err)
			fmt.Printf("Record with UPC %s deleted successfully.\n", upc)
		}
		return true
	case "3":
		rows, err := getRecords(db, username)
		iferr(err)
		defer rows.Close()
		fmt.Println("title\tartist\tmedium\tformat\tlabel\tgenre\tyear\tupc")
		arr, err := recordRowsToArray(rows)
		iferr(err)
		for i := range arr {
			arr[i][8] = ""
		}
		for _, row := range arr {
			for _, s := range row {
				fmt.Printf("%s\t", s)
			}
			fmt.Println()
		}
		return true
	case "4":
		fmt.Print("Enter the CSV filename to export to\n>")
		fn, err := reader.ReadString('\n')
		iferr(err)

		fn = stripFormatting(fn)
		rows, err := getRecords(db, username)
		iferr(err)

		defer rows.Close()
		arr, err := recordRowsToArray(rows)
		iferr(err)

		err = writeToFile(fn, recordsToCSVString(arr))
		iferr(err)

		return true
	case "0":
		return false
	}
	return true
}

// prompt each part of the album manually
func promptAlbum(reader *bufio.Reader) ([8]string, error) {
	fields := [8]string{"title", "artist", "medium", "format", "label", "genre", "year", "UPC"}
	var out [8]string
	for i, v := range fields {
		fmt.Printf("Enter the %s\n>", v)
		input, err := reader.ReadString('\n')
		if err != nil {
			return fields, err
		}
		out[i] = stripFormatting(input)
	}
	return out, nil
}

// prompt each part of the album using only the UPC
func promptUPC(reader *bufio.Reader) ([8]string, error) {
	var out [8]string
	fmt.Print("Enter the UPC\n>")
	input, err := reader.ReadString('\n')
	if err != nil {
		return out, err
	}
	return getAlbumInfo(stripFormatting(input))
}

// prompt uname and pass
func login(reader *bufio.Reader) (string, string, error) {
	var err error
	var uname string
	var password string
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
	password = fmt.Sprintf("%x", sha256.Sum256([]byte(bword)))
	fmt.Println()
	return stripFormatting(uname), password, nil
}
