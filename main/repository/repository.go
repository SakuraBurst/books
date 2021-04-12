package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/SakuraBurst/books.git/main/helpers"
	"github.com/SakuraBurst/books.git/main/models"
)

type Repository struct {
	Database *sql.DB
}

func (r Repository) GetAllFromDatabase(sl *[]models.InstanseMaker, inst models.InstanseMaker) error {
	table := getInstanseTable(inst)
	var query = fmt.Sprintf("SELECT * FROM %v", table)
	fmt.Println(query)
	rows, err := r.Database.Query(query)
	if err != nil {
		return err
	}
	for rows.Next() {
		copy, err := inst.NewInstanseFromDB(rows)
		if err != nil {
			return err
		} else {
			*sl = append(*sl, copy)
		}

	}
	defer rows.Close()
	return nil
}

func (r Repository) GetOneFromDatabase(value interface{}, key string, instanse models.InstanseMaker) (models.InstanseMaker, error) {
	if len(key) == 0 {
		key = "id"
	}
	table := getInstanseTable(instanse)
	query := fmt.Sprintf(`SELECT * FROM %v WHERE %v = $1`, table, key)
	row := r.Database.QueryRow(query, value)
	newInstanse, err := instanse.NewInstanseFromDB(row)
	if err != nil {
		return nil, err
	}

	return newInstanse, nil
}

func (r Repository) WriteToTheDatabase(newInstanse models.InstanseMaker, body io.ReadCloser) error {
	newInstanse, err := helpers.MakeNewInstanse(newInstanse, body)

	if err != nil {
		return err
	} else {
		fields := models.GetFields(newInstanse)
		query := generateAddQuery(getInstanseTable(newInstanse), fields)
		_, err := r.Database.Exec(query, fields.BdValues...)
		if err != nil {
			return err
		}
		return nil
	}

}

func (r Repository) UpdateFromDatabase(newInstanse models.InstanseMaker, body io.ReadCloser, id string) error {
	table := getInstanseTable(newInstanse)
	if r.checkDatabaseForIdExisting(id, table) {
		newInstanse, err := helpers.MakeNewInstanse(newInstanse, body)
		if err != nil {
			fmt.Println("error")
			return err
		} else {
			fields := models.GetFields(newInstanse)
			query := generateReplaceQuery(table, fields, id)
			r.Database.Exec(query, fields.BdValues...)
			return nil

		}
	} else {
		return errors.New("id does not exist")
	}

}

func (r Repository) DeleteFromDatabase(id, table string) error {
	if r.checkDatabaseForIdExisting(id, table) {
		insertString := fmt.Sprintf(`DELETE FROM %v WHERE id = $1`, table)
		r.Database.Exec(insertString, id)
		return nil
	} else {
		return errors.New("id does not exist")
	}

}

func (r Repository) checkDatabaseForIdExisting(id, table string) bool {
	sqlStmt := fmt.Sprintf(`SELECT * FROM %v WHERE id = $1`, table)
	err := r.Database.QueryRow(sqlStmt, id).Scan()
	return err != sql.ErrNoRows
}

func getInstanseTable(instanse models.InstanseMaker) string {
	instType := fmt.Sprintf("%T", instanse)
	return models.Tables[instType]
}

func generateAddQuery(table string, fields models.Fields) string {
	fieldsStr := ""
	valuessStr := "("
	dollarCounter := 0
	for i := range fields.BdValues {
		dollarCounter++

		if i+1 == len(fields.BdValues) {
			valuessStr = valuessStr + fmt.Sprintf("$%v)", dollarCounter)
		} else {
			valuessStr = valuessStr + fmt.Sprintf("$%v,", dollarCounter)
		}

	}
	fieldsStr = strings.Join(fields.BdFields, ", ")
	return fmt.Sprintf("INSERT INTO %v(%v) VALUES%v", table, fieldsStr, valuessStr)

}

func generateReplaceQuery(table string, fields models.Fields, id string) string {
	replaceStr := ""
	dollarCounter := 0
	for i, v := range fields.BdFields {
		dollarCounter++
		if i+1 == len(fields.BdFields) {
			replaceStr = replaceStr + fmt.Sprintf("%v = $%v WHERE id = %v", v, dollarCounter, id)
		} else {
			replaceStr = replaceStr + fmt.Sprintf("%v = $%v,", v, dollarCounter)
		}

	}
	return fmt.Sprintf("UPDATE %v SET %v", table, replaceStr)

}
