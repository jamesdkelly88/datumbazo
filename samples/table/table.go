package main

import (
	"fmt"
	"slices"
	"strconv"
)

type Row struct {
	RowId int64
	Data  []any
}

type Column struct {
	Name string
	Type string
}

type Table struct {
	Name    string
	Columns []Column
	Rows    []Row
}

func (t *Table) AddColumn(name string, datatype string) {
	c := Column{Name: name, Type: datatype}
	t.Columns = append(t.Columns, c)
}
func (t *Table) AddRow(data []any) {
	var rowId int64
	if len(t.Rows) == 0 {
		rowId = 1
	} else {
		lastRow := t.Rows[len(t.Rows)-1]
		rowId = lastRow.RowId + 1
	}
	newRow := Row{RowId: rowId, Data: data}
	t.Rows = append(t.Rows, newRow)
}
func (t *Table) GetValues(column string) []any {
	idx := slices.IndexFunc(t.Columns, func(c Column) bool { return c.Name == column })

	var values []any
	for _, item := range t.Rows {
		values = append(values, item.Data[idx])
	}

	return values
}

func (t *Table) GetColumnWidth(column string) int {
	idx := slices.IndexFunc(t.Columns, func(c Column) bool { return c.Name == column })
	coltype := t.Columns[idx].Type
	var w int
	for _, item := range t.Rows {
		var l int
		i := item.Data[idx]
		switch coltype {
		case "int":
			l = len(strconv.Itoa(i.(int)))
		case "varchar":
			l = len(i.(string))
		default:
			return -1
		}
		if l > w {
			w = l
		}
	}
	if w < len(column) {
		w = len(column) //rune count?
	}
	return w
}
func (t *Table) Print() {
	// Get width for each column
	// Print each header padded, separated by " | "
	// Print a line of "---"
	// Print each row, padding values
}

func main() {
	t := Table{Name: "Test"}

	t.AddColumn("ID", "int")
	t.AddColumn("Name", "varchar")

	t.AddRow([]any{4645746, "D"})
	t.AddRow([]any{5, "Eagles"})
	t.AddRow([]any{62, "Fish"})

	fmt.Printf("Table name: %s\n", t.Name)
	fmt.Printf("Columns: %d\n", len(t.Columns))
	fmt.Printf("Rows: %d\n", len(t.Rows))

	fmt.Printf("IDs: %v (longest = %d)\n", t.GetValues("ID"), t.GetColumnWidth("ID"))
	fmt.Printf("Names: %v (longest = %d)\n", t.GetValues("Name"), t.GetColumnWidth("Name"))

	// t.Print()

}
