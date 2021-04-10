package controllers

import (
	"encoding/json"
	"net/http"
)

func (c Controler) responseEncoder(rw http.ResponseWriter) *json.Encoder {
	return json.NewEncoder(rw)

}
