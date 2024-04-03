package main

import (
	"fmt"
	"os"
	"strings"
)

// remove lf and cr from string
func stripFormatting(str string) string {
	newStr := str
	newStr = strings.ReplaceAll(newStr, "\n", "")
	newStr = strings.ReplaceAll(newStr, "\r", "")
	return newStr
}

// read and split a cfg file
func readConfig(path string) ([]string, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(body), "\n"), nil
}

// get var_ from cfg array
func getVar(var_ string, cfg []string) string {
	for _, s := range cfg {
		if strings.Split(s, "=")[0] == var_ {
			return strings.Split(s, "=")[1]
		}
	}
	return ""
}

// iferr -> os.exit(1) + print
func iferr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// write string to file with name fn
func writeToFile(fn string, out string) error {
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	_, err = f.WriteString(out)
	f.Close()
	return err
}

// take a 2d array of records and convert it to a csv string
func recordsToCSVString(arr [][]string) string {
	out := ""
	out += fmt.Sprintln("title,artist,medium,format,label,genre,year,upc,")
	for i := range arr {
		arr[i][8] = ""
	}
	for _, row := range arr {
		for i := 0; i < 8; i++ {
			out += fmt.Sprintf("%s,", row[i])
		}
		out += fmt.Sprintln()
	}
	return out
}
