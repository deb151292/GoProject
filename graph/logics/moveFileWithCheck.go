package logics

import (
	"errors"

	"github.com/deb151292/gqlgen-todos/graph/model"
)

func Moving(input *model.Paths) (*model.Response, error) {
	var response model.Response
	var erro error
	err := DoesFileExist(input.FileDestinationPath)
	if err == true {
		erro = errors.New("File Not Found!")
		response.Err = true
		response.Massage = "Sorry!"
		return &response, erro
	}

	err1 := Copy(input.FileDestinationPath, input.FileDestinationPath)
	if err1 != nil {
		err1 = errors.New("File can't be moved")
		response.Err = true
		response.Massage = "Sorry!"

		return &response, err1
	} else {
		err2 := Delete(input.FileDestinationPath)
		if err2 != nil {
			err2 = errors.New("File copied but couldnot be moved!")
			response.Err = true
			response.Massage = "Sorry!"

			return &response, err1
		} else {
			response.Err = false
			response.Massage = "Done."

			return &response, nil
		}
	}
}
