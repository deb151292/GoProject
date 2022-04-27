package logics

import (
	"errors"
	"log"

	"github.com/deb151292/gqlgen-todos/graph/model"
	"github.com/skip2/go-qrcode"
)

func QrGenerator(input model.InputStr) (*model.SuccessMsg, error) {
	log.Println("a")
	var Output model.SuccessMsg
	var booleanData string

	if input.Pass == true {
		booleanData = "true"
	} else {
		booleanData = "false"
	}
	datastring := input.StudentName + booleanData
	err := qrcode.WriteFile(string(datastring), qrcode.Medium, 256, (input.StudentName + ".png"))

	if err != nil {
		err = errors.New("There is a problem while building QR code !")
		Output.OutputStr = "Error"
	} else {
		Output.OutputStr = "Processing..."
	}
	A := "D:/go lang/22-04-2022/GoProject/" + input.StudentName + ".png"
	B := "D:/" + input.StudentName + ".png"
	err1 := move(A, B)
	if err1 != nil {
		Output.OutputStr = "Processing Faild...!"
	} else {
		Output.OutputStr = "Download Success!"
	}
	return &Output, err
}
