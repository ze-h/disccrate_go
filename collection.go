package main

import (
	"github.com/irlndts/go-discogs"
)

var CLIENT discogs.Discogs

func init_api() error {
	cfg, err := readConfig("discogs.cfg")
	if err != nil {
		return err
	}
	CLIENT, err = discogs.New(&discogs.Options{
		UserAgent: getVar("AGENT", cfg),
		Token:     getVar("TOKEN", cfg),
	})
	return err
}
