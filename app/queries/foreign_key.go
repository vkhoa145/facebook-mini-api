package queries

import (
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

func (q *Queries) IsForeignKeyExisted(value interface{}, field string, table interface{}) (bool, error) {
	if reflect.TypeOf(table).Kind() != reflect.Struct {
		return false, errors.New("Not a valid struct")
	}

	column := fmt.Sprintf("%s = ?", field)
	result := q.db.Where(column, value).First(&table)
	if result.Error == nil {
		return true, nil
	}

	if result.Error == gorm.ErrRecordNotFound {
		fmt.Println("Primary Key is not existed")
		return false, errors.New("Primary Key is not existed")
	} else {
		fmt.Println("Error occured:", result.Error)
		return false, errors.New("Unexpected Error occured")
	}
}
