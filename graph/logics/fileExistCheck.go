package logics

import "os"

func DoesFileExist(fileName string) bool {

	var err bool

	_, errors := os.Stat(fileName)
	if errors != nil {
		err = true
	} else {
		err = false
	}
	return err
}
