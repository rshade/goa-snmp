package db

import (
	"log"
	badger "github.com/dgraph-io/badger/v3"
)

func badgerOpen() {
  // Open the Badger database located in the /tmp/badger directory.
  // It will be created if it doesn't exist.
  db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
  if err != nil {
	  log.Fatal(err)
  }
  defer db.Close()
  // Your code here…
}