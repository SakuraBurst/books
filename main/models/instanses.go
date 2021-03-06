package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type Scans interface {
	Scan(dest ...interface{}) error
}
type InstanseMaker interface {
	NewInstanseFromJson(body io.ReadCloser) (InstanseMaker, error)
	NewInstanseFromDB(row Scans) (InstanseMaker, error)
	IsValid() bool
}

type Fields struct {
	BdFields []string
	BdValues []interface{}
}

var Tables map[string]string = map[string]string{
	"models.Book": "books",
	"models.User": "users",
}

func GetFields(instanse InstanseMaker) Fields {
	instanseMap := getMap(instanse)
	var bdFields []string
	var bdValues []interface{}
	var fields Fields
	for key, value := range instanseMap {
		if key != "ID" && key != "id" {
			bdFields = append(bdFields, key)
			switch typedValue := value.(type) {
			case string:
				value := fmt.Sprintf("%v", typedValue)
				bdValues = append(bdValues, value)
			}
		}

	}
	fields.BdFields = bdFields
	fields.BdValues = bdValues
	return fields
}

func getMap(instanse InstanseMaker) map[string]interface{} {
	byteJson, err := json.Marshal(instanse)
	if err != nil {
		log.Fatal(err)
	}
	var dummy map[string]interface{}
	err = json.Unmarshal(byteJson, &dummy)
	if err != nil {
		log.Fatal(err)
	}
	return dummy
	// var dummy interface{}
	// return dummy.(map[string]interface{})
}
