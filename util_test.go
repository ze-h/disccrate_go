package main

import (
	"fmt"
	"testing"
)

func TestCFG(t *testing.T){
	cfg, err := readConfig("test.cfg")
	if err != nil {
		t.Fatal(err)
	}
	if getVar("DATA1", cfg) == "element1" {
		fmt.Println("PASS DATA1")
	}else{
		fmt.Printf("FAIL DATA1: %s\n", getVar("DATA1", cfg))
		t.Fail()
	}
	if getVar("DATA2", cfg) == "element2" {
		fmt.Println("PASS DATA2")
	}else{
		fmt.Printf("FAIL DATA2: %s\n", getVar("DATA2", cfg))
		t.Fail()
	}
}