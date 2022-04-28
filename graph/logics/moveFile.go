package logics

import (
	"log"
	"os"
)

func Move(spath, dpath string) error {
	// oldLocation := "/var/www/html/test.txt"
	// newLocation := "/var/www/html/src/test.txt"
	err := os.Rename(spath, dpath)

	if err != nil {
		log.Fatal(err)
	}
	return err
}
