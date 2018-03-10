package datastore

import (
	"log"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func init() {
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}
