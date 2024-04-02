package main

import (
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