package server

type DeleteDto struct {
	Id       string `json:"id" from:"id" query:"id"`
	FileName string `json:"filename" from:"filename" query:"filename"`
}
