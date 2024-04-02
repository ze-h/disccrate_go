package main

import (
	"fmt"
	"testing"

	"github.com/irlndts/go-discogs"
)

func TestDiscogsSearch(t *testing.T) {
	DISCOGS_CONFIG = "discogs.cfg"
	err := initApi()
	if err != nil {
		t.Fatal(err)
	}
	album, err := CLIENT.Search(discogs.SearchRequest{Barcode: "093624877721"}) // American Idiot EU CD
	if err != nil {
		t.Fatal(err)
	}

	if album.Results[0].Title != "Green Day - American Idiot" {
		t.Fatal("FAIL: ", album.Results[0].Title)
	}
}

func TestLongInput(t *testing.T) {
	DISCOGS_CONFIG = "discogs.cfg"
	err := initApi()
	if err != nil {
		t.Fatal(err)
	}
	album, err := CLIENT.Search(discogs.SearchRequest{Barcode: "093624353126"}) // Hey Man Nice Shot Maxi Single
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(album.Results[0].Title)
}