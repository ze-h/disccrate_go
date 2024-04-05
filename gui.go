package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	_ "github.com/go-sql-driver/mysql"
)

// run the gui userloop
func gui(db *sql.DB) {
	var usr_temp, usr, pass string
	var collection [][]string
	var album, album_tmp [8]string

	app := tview.NewApplication()
	login_form := tview.NewForm()
	home_screen := tview.NewList()
	auto_record_popup := tview.NewModal()
	auto_record_popup_pass := tview.NewModal()
	add_record_popup := tview.NewModal()
	add_record_popup_pass := tview.NewModal()
	rem_record_popup := tview.NewModal()
	rem_record_popup_pass := tview.NewModal()
	incorrect_login := tview.NewModal()
	add_record_auto := tview.NewForm()
	add_record := tview.NewForm()
	remove_record := tview.NewForm()
	collection_table := tview.NewTable()

	incorrect_login.
		SetText("Incorrect username or password.").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				app.SetRoot(login_form, true)
			}
		})

	auto_record_popup.
		SetText("Adding record failed! See log for details.").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				app.SetRoot(add_record_auto, true)
			}
		})

	auto_record_popup_pass.
		SetText("Record added successfully!").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				app.SetRoot(add_record_auto, true)
			}
		})

	add_record_popup.
		SetText("Adding record failed! See log for details.").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				app.SetRoot(add_record, true)
			}
		})

	add_record_popup_pass.
		SetText("Record added successfully!").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				app.SetRoot(add_record, true)
			}
		})

	rem_record_popup.
		SetText("Removing record failed! See log for details.").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				app.SetRoot(remove_record, true)
			}
		})

	rem_record_popup_pass.
		SetText("Record removed successfully!").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				app.SetRoot(remove_record, true)
			}
		})

	login_form.
		AddInputField("Username", "", 16, nil, func(str string) {
			usr_temp = str
		}).
		AddPasswordField("Password", "", 16, '*', func(str string) {
			pass = str
		}).
		AddButton("Submit", func() {
			usr = usr_temp
			v, err := verify(db, usr, fmt.Sprintf("%x", md5.Sum([]byte(pass))))
			if err != nil {
				panic(err)
			}
			if v {
				app.SetRoot(home_screen, true)
			} else {
				app.SetRoot(incorrect_login, false)
			}
		}).
		AddButton("Quit", func() {
			app.Stop()
		}).
		SetBorder(true).
		SetTitle("DiscCrate").
		SetTitleAlign(tview.AlignCenter)

	home_screen.
		AddItem("Add Record", "Manually add record to collection", 'a', func() {
			app.SetRoot(add_record, true)
		}).
		AddItem("Add Record (Automatic)", "Add record to collection using UPC", 'b', func() {
			app.SetRoot(add_record_auto, true)
		}).
		AddItem("Remove Record", "Remoce record from collection using UPC", 'c', func() {
			app.SetRoot(remove_record, true)
		}).
		AddItem("See Collection", "Display record collection", 'd', func() {
			app.SetRoot(collection_table, true)
			records, err := getRecords(db, usr)
			iferr(err)
			collection, err = recordRowsToArray(records)
			iferr(err)
			for i, s := range []string{"title", "artist", "medium", "format", "label", "genre", "year", "UPC"} {
				collection_table.SetCell(0, i, tview.NewTableCell(s).SetAlign(tview.AlignCenter))
			}
			for i, s := range collection {
				for j, ss := range s {
					collection_table.SetCell(i+1, j, tview.NewTableCell(ss).SetAlign(tview.AlignCenter))
				}
			}
			app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				if event.Rune() == 'q' {
					app.SetRoot(home_screen, true)
				}
				return event
			})
		}).
		AddItem("Exit", "", 'e', func() {
			app.Stop()
		}).
		SetBorder(true).
		SetTitle("DiscCrate: Main Menu").
		SetTitleAlign(tview.AlignCenter)

	add_record_auto.
		AddInputField("UPC", "", 32, nil, func(str string) {
			album_tmp[7] = str
		}).
		AddButton("Submit", func() {
			album = album_tmp
			album, err := getAlbumInfo(album[7])
			if err != nil {
				app.SetRoot(auto_record_popup, false)
				log.Println(err)
			} else {
				_, err = addRecord(db, album, usr)
				if err != nil {
					app.SetRoot(auto_record_popup, false)
					log.Println(err)
				} else {
					app.SetRoot(auto_record_popup_pass, false)
				}
			}
		}).
		AddButton("Quit", func() {
			app.SetRoot(home_screen, true)
		}).
		SetBorder(true).
		SetTitle("DiscCrate: Add Record (Auto)").
		SetTitleAlign(tview.AlignCenter)

	add_record.
		AddInputField("Title", "", 64, nil, func(str string) {
			album_tmp[0] = str
		}).
		AddInputField("Artist", "", 48, nil, func(str string) {
			album_tmp[1] = str
		}).
		AddInputField("Year", "", 8, nil, func(str string) {
			album_tmp[6] = str
		}).
		AddInputField("Genre", "", 16, nil, func(str string) {
			album_tmp[5] = str
		}).
		AddInputField("Medium", "", 16, nil, func(str string) {
			album_tmp[2] = str
		}).
		AddInputField("Format", "", 16, nil, func(str string) {
			album_tmp[3] = str
		}).
		AddInputField("Label", "", 32, nil, func(str string) {
			album_tmp[4] = str
		}).
		AddInputField("UPC", "", 32, nil, func(str string) {
			album_tmp[7] = str
		}).
		AddButton("Submit", func() {
			album = album_tmp
			_, err := addRecord(db, album, usr)
			if err != nil {
				app.SetRoot(add_record_popup, false)
				log.Println(err)
			} else {
				app.SetRoot(add_record_popup_pass, false)
			}
		}).
		AddButton("Quit", func() {
			app.SetRoot(home_screen, true)
		}).
		SetBorder(true).
		SetTitle("DiscCrate: Add Record").
		SetTitleAlign(tview.AlignCenter)

	remove_record.
		AddInputField("UPC", "", 32, nil, func(str string) {
			album_tmp[7] = str
		}).
		AddButton("Submit", func() {
			album = album_tmp
			_, err := deleteRecord(db, album[7], usr)
			if err != nil {
				app.SetRoot(rem_record_popup, false)
				log.Println(err)
			} else {
				app.SetRoot(rem_record_popup_pass, false)
			}
		}).
		AddButton("Quit", func() {
			app.SetRoot(home_screen, true)
		}).
		SetBorder(true).
		SetTitle("DiscCrate: Remove Record").
		SetTitleAlign(tview.AlignCenter)

	collection_table.
		SetSelectable(true, false).
		SetBorder(true).
		SetTitle("DiscCrate: Collection View (Press q to return)").
		SetTitleAlign(tview.AlignCenter)

	// run application window
	if err := app.SetRoot(login_form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
