package main

import (
	"testing"

	"github.com/irlndts/go-discogs"
)

func TestDiscogsSearch(t *testing.T) {
	err := init_api()
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
