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
	if init_api() != nil {
		fmt.Println("Discogs init failed!")
		return
	}
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
	reader := bufio.NewReader(os.Stdin)
	for usermode(usr, reader, db) {
	}
	fmt.Println("Goodbye!")
}

// usermode loop
func usermode(username string, reader *bufio.Reader, db *sql.DB) bool {
	fmt.Print("Select an option:\n1 - Add record to collection\n2 - Query collection\n3 - Remove from collection\n>")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return false
	}
	switch stripFormatting(input) {
	case "1":
		fmt.Print("Enter 1 for manual, 2 for automatic\n>")
		input_2, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return false
		}
		var album [8]string
		if stripFormatting(input_2) == "1" {
			album, err = promptAlbum(reader)
			if err != nil {
				fmt.Println(err)
				return false
			}
		} else {
			album, err = promptUPC(reader)
			if err != nil {
				fmt.Println(err)
				return false
			}
		}
		_, err = db.Query("INSERT INTO records (title, artist, medium, format, label, genre, year, upc, username) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", album[0], album[1], album[2], album[3], album[4], album[5], album[6], album[7], username)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	case "2":
		rows, err := db.Query("SELECT * FROM records WHERE username=?", username)
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer rows.Close()
		fmt.Println("title\tartist\t\tmedium\tformat\tlabel\tgenre\tyear\tupc")
		for rows.Next() {
			var title, artist, medium, format, label, genre, year, upc, record_user string
			if err := rows.Scan(&title, &artist, &medium, &format, &label, &genre, &year, &upc, &record_user); err != nil {
				fmt.Println(err)
				return false
			}
			fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", title, artist, medium, format, label, genre, year, upc)
		}
		fmt.Println()
		return true
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
	return getAlbumInfo(input)
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
