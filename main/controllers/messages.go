package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SakuraBurst/books.git/main/models"
)

func ServerStop(err error) {
	if err != nil {
		log.Fatal(err)
	}

}

func SendErrorMessage(rw http.ResponseWriter, errorMessage error, errorStatus int) {
	rw.WriteHeader(errorStatus)
	var errorjson []byte
	if errorMessage != nil {
		errorMessage := models.Message{Text: errorMessage.Error(), Status: "error"}
		errorjson, _ = json.Marshal(errorMessage)
	} else {
		errorjson, _ = json.Marshal(models.ErrorMessage)
	}

	rw.Write(errorjson)
}
