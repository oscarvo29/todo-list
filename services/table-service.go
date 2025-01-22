package services

import (
	"fmt"
	"reflect"
	"time"

	"github.com/oscarvo29/todo-list/models"
	"github.com/oscarvo29/todo-list/utils"
)

func InstansiateTable(tasks []models.Task) *models.Table {
	table := models.NewTable()

	for _, task := range tasks {
		if !table.IsColumnSet() {
			columns := GetFieldsNamesFromStruct(task)
			table.AddColumn(columns)
		}

		row := GetFieldValuesFromStructAsStrings(task)
		table.AddRow(row)
	}

	return table
}

func GetFieldsNamesFromStruct(t any) []string {
	var fieldNames []string

	val := reflect.TypeOf(t)
	for index := range val.NumField() {
		name := val.Field(index).Name
		fieldNames = append(fieldNames, name)
	}

	return fieldNames
}

func GetFieldValuesFromStructAsStrings(t any) []string {
	var fieldValues []string

	val := reflect.ValueOf(t)
	for index := range val.NumField() {
		field := val.Field(index)

		if field.Type() == reflect.TypeOf(time.Time{}) {
			if fieldTime, ok := field.Interface().(time.Time); ok {
				formattedTime := utils.CompareTimes(fieldTime)
				fieldString := fmt.Sprintf("%v", formattedTime)
				fieldValues = append(fieldValues, fieldString)
				continue
			}
		}

		fieldString := fmt.Sprintf("%v", field)
		fieldValues = append(fieldValues, fieldString)
	}

	return fieldValues
}
