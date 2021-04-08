package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type InstanseMaker interface {
	NewInstanseFromJson(body io.ReadCloser)
	IsValid() bool
}

type Fields struct {
	BdFields string
	BdValues string
}

var Tables map[string]string = map[string]string{
	"*models.Book": "books",
	"*models.User": "books",
}

func GetFields(instanse InstanseMaker) Fields {
	instanseMap := getMap(instanse)
	var bdFields []string
	var bdValues []string
	var fields Fields
	for key, value := range instanseMap {
		if key != "ID" {
			bdFields = append(bdFields, key)
			switch typedValue := value.(type) {
			case string:
				value := fmt.Sprintf("'%v'", typedValue)
				bdValues = append(bdValues, value)
			}
		}

	}
	fields.BdFields = strings.Join(bdFields, ", ")
	fields.BdValues = strings.Join(bdValues, ", ")
	return fields
}

func getMap(instanse InstanseMaker) map[string]interface{} {
	byteJson, err := json.Marshal(instanse)
	if err != nil {
		log.Fatal(err)
	}
	var dummy interface{}
	err = json.Unmarshal(byteJson, &dummy)
	if err != nil {
		log.Fatal(err)
	}
	return dummy.(map[string]interface{})
}
