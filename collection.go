package main

import (
	"github.com/irlndts/go-discogs"
)

var CLIENT discogs.Discogs

func init_api(){
	CLIENT, err := discogs.New(&discogs.Options{
        UserAgent: "Some Name",
        Token:     "Some Token",
    })
}