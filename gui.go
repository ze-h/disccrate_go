package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"

	_ "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	_ "github.com/go-sql-driver/mysql"
)

// run the gui userloop
func gui(db *sql.DB) {
	var usr, pass string
	app := tview.NewApplication()
	login_form := tview.NewForm()
	incorrect_password_popup := tview.NewModal()
	collection_table := tview.NewTable()

	incorrect_password_popup.
		SetText("Incorrect username or password.").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				app.SetRoot(login_form, true)
			}
		})

	login_form.
		AddInputField("Username", "", 16, nil, func(str string) {
			usr = str
		}).
		AddPasswordField("Password", "", 16, '*', func(str string) {
			pass = str
		}).
		AddButton("Submit", func() {
			v, err := verify(db, usr, fmt.Sprintf("%x", md5.Sum([]byte(pass))))
			if err != nil {
				panic(err)
			}
			if v {
				login_form.SetTitle("TRUE")
			} else {
				app.SetRoot(incorrect_password_popup, false)
			}
		}).
		AddButton("Quit", func() {
			app.Stop()
		}).
		SetBorder(true).
		SetTitle("DiscCrate").
		SetTitleAlign(tview.AlignCenter)

	collection_table.
			SetBorder(true). 
			SetTitle("DiscCrate").
			SetTitleAlign(tview.AlignCenter)

	// run application window
	if err := app.SetRoot(login_form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
