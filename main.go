package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if len(os.Args) == 3 {
		if os.Args[1] == "admin" {
			admin_cfg, err := readConfig("admin.cfg")
			if err != nil {
				fmt.Println("Admin config read error")
				fmt.Println(err)
				os.Exit(1)
			}
			sum := fmt.Sprintf("%x", md5.Sum([]byte(os.Args[2])))
			if sum == getVar("KEY", admin_cfg) {
				admin()
			} else {
				fmt.Println("Incorrect administrator password.")
			}
			return
		}
	}
	if initApi() != nil {
		fmt.Println("Discogs init failed!")
		os.Exit(1)
	}
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
	if len(os.Args) == 1 {
		cli(db)
	} else {
		fmt.Print("Invalid argument(s): ")
		for _, v := range os.Args {
			fmt.Printf("%s ", v)
		}
		fmt.Println()
		os.Exit(1)
	}
}
