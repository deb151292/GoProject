// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type InputStr struct {
	StudentName string `json:"StudentName"`
	Pass        bool   `json:"Pass"`
}

type Paths struct {
	DestinationPath     string `json:"Destination_path"`
	FileDestinationPath string `json:"File_Destination_path"`
}

type SuccessMsg struct {
	OutputStr string `json:"OutputStr"`
}

type Response struct {
	Err     bool   `json:"err"`
	Massage string `json:"massage"`
}
