package logics

import (
	"log"
	"os"
)

func Delete(s string) error {
	e := os.Remove(s)

	if e != nil {
		log.Println(e)
		log.Fatal(e)
	}
	return e
}
