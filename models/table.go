package models

import (
	"fmt"

	"github.com/oscarvo29/todo-list/utils"
)

type Table struct {
	Columns     []string
	Rows        [][]string
	ColumnWidth []int // This is to manage how much space each field in a row need
}

func NewTable() *Table {
	return &Table{
		Columns:     []string{},
		Rows:        [][]string{},
		ColumnWidth: []int{},
	}
}

func (t *Table) AddColumn(column []string) {
	for _, field := range column {
		t.ColumnWidth = append(t.ColumnWidth, len(field))
	}

	t.Columns = append(t.Columns, column...)
}

func (t *Table) IsColumnSet() bool {
	return len(t.Columns) >= 1
}

func (t *Table) AddRow(row []string) {
	for index, field := range row {
		if len(field) > t.ColumnWidth[index] {
			t.ColumnWidth[index] = len(field)
		}
	}

	t.Rows = append(t.Rows, row)
}

func (t *Table) PrintTable() {
	for index, column := range t.Columns {
		columnLen := len(column)

		sizeDiff := t.ColumnWidth[index] - columnLen
		fmt.Printf("%s", column)
		utils.PrintBlankSpaces(sizeDiff)
	}

	for _, row := range t.Rows {
		for index, field := range row {
			if index == 0 { // creates a new line, each time a new row is getting printed.
				fmt.Print("\n")
			}
			fieldLen := len(field)
			sizeDiff := t.ColumnWidth[index] - fieldLen

			fmt.Printf("%s", field)
			utils.PrintBlankSpaces(sizeDiff)
		}
	}

	fmt.Print("\n")
}
