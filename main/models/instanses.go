package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type InstanseMaker interface {
	NewInstanseFromJson(body io.ReadCloser) InstanseMaker
	NewInstanseFromDB(row *sql.Rows) (InstanseMaker, error)
	IsValid() bool
}

type Fields struct {
	BdFields []string
	BdValues []interface{}
}

var Tables map[string]string = map[string]string{
	"models.Book": "books",
	"models.User": "books",
}

func GetFields(instanse InstanseMaker) Fields {
	instanseMap := getMap(instanse)
	var bdFields []string
	var bdValues []interface{}
	var fields Fields
	for key, value := range instanseMap {
		if key != "ID" {
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

// func GetReplaceFields(instanse InstanseMaker) ReplaceFields {
// 	instanseMap := getMap(instanse)
// 	var replaceMap ReplaceFields = make(ReplaceFields)
// 	for key, value := range instanseMap {
// 		if key != "ID" {
// 			switch typedValue := value.(type) {
// 			case string:
// 				replaceMap[key] = typedValue
// 			}
// 		}

// 	}
// 	return replaceMap
// }

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
