package models

type Message struct {
	Text   string `json:"year"`
	Status string `json:"status"`
}

func (err Message) Error() string {
	return err.Text
}

var ErrorMessage Message = Message{Text: "something went wrong", Status: "error"}
var SuccessMessage Message = Message{Text: "zaebis", Status: "success"}
