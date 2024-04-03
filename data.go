package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// verify username and password
func verify(db *sql.DB, username string, password string) (bool, error) {
	rows, err := db.Query("SELECT password FROM users WHERE username = ?", username)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	var pswd string

	for rows.Next() {
		if err := rows.Scan(&pswd); err != nil {
			return false, err
		}
	}

	return (pswd == password), nil
}

// given an array containing an album's information, add to username's collection in db
func addRecord(db *sql.DB, album [8]string, username string) (*sql.Rows, error) {
	return db.Query("INSERT INTO records (title, artist, medium, format, label, genre, year, upc, username) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", album[0], album[1], album[2], album[3], album[4], album[5], album[6], album[7], username)
}

// return all records belonging to user in db
func getRecords(db *sql.DB, username string) (*sql.Rows, error) {
	return db.Query("SELECT * FROM records WHERE username=?", username)
}

// delete record with given upc and username
func deleteRecord(db *sql.DB, upc string, username string) (*sql.Rows, error) {
	return db.Query("DELETE FROM records WHERE username = ? AND upc = ?", username, upc)
}

func findRecord(db *sql.DB, upc string, username string) (string, error) {
	rows, err := db.Query("SELECT title FROM records WHERE username = ? AND upc = ?", username, upc)
	if err != nil {
		return "", err
	}
	var title string

	for rows.Next() {
		if err := rows.Scan(&title); err != nil {
			return "", err
		}
	}

	return title, nil
}
