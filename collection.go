package main

import (
	"strings"

	"github.com/irlndts/go-discogs"
)

var CLIENT discogs.Discogs

// initialize the discogs api using the info stored in discogs.cfg
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

// use the discogs api to get information on an album using a UPC as input
func getAlbumInfo(upc string) ([8]string, error) {
	var album_info [8]string // title, artist, medium, format, label, genre, year, upc
	res, err := CLIENT.Search(discogs.SearchRequest{Barcode: upc})
	if err != nil {
		return album_info, err
	}
	album := res.Results[0]
	album_info[0] = strings.Split(album.Title, " - ")[1]
	album_info[1] = strings.Split(album.Title, " - ")[0]
	album_info[2] = album.Format[0]
	album_info[3] = album.Format[1]
	album_info[4] = album.Label[0]
	album_info[5] = album.Genre[0]
	album_info[6] = album.Year
	album_info[7] = upc
	return album_info, nil
}
