package helpers

import (
	"encoding/json"

	"github.com/SakuraBurst/books.git/main/models"
)

func DeletePasswordField(object models.InstanseMaker) (map[string]interface{}, error) {
	var mapWithoutPasswordField map[string]interface{}
	encoder, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(encoder, &mapWithoutPasswordField)
	if err != nil {
		return nil, err
	}
	delete(mapWithoutPasswordField, "password")
	return mapWithoutPasswordField, nil
}
