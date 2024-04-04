package main

import (
	"fmt"
	"testing"
)

func TestCFG(t *testing.T) {
	cfg, err := readConfig("test.cfg")
	if err != nil {
		t.Fatal(err)
	}
	if getVar("DATA1", cfg) == "element1" {
		fmt.Println("PASS DATA1")
	} else {
		fmt.Printf("FAIL DATA1: %s\n", getVar("DATA1", cfg))
		t.Fail()
	}
	if getVar("DATA2", cfg) == "element2" {
		fmt.Println("PASS DATA2")
	} else {
		fmt.Printf("FAIL DATA2: %s\n", getVar("DATA2", cfg))
		t.Fail()
	}
}

func TestCSV(t *testing.T) {
	data_1 := []string{"A", "A", "A", "A", "A", "A", "A", "A", "A"}
	data_2 := []string{"B", "B", "B", "B", "B", "B", "B", "B", "B"}
	data_3 := []string{"C", "C", "C", "C", "C", "C", "C", "C", "C"}
	data := [][]string{data_1, data_2, data_3}

	expected := "title,artist,medium,format,label,genre,year,upc,\nA,A,A,A,A,A,A,A,\nB,B,B,B,B,B,B,B,\nC,C,C,C,C,C,C,C,\n"
	actual := recordsToCSVString(data)

	fmt.Printf("EXPECTED: %s\nACTUAL: %s\n", expected, actual)

	if expected != actual {
		t.Fail()
	}
}
