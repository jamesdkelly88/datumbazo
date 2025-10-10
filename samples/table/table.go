package main

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
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

func (t *Table) GetColumnWidths() []int {

	var widths []int

	for _, col := range t.Columns {
		w := t.GetColumnWidth(col.Name)
		widths = append(widths, w)
	}

	return widths
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
			l = utf8.RuneCountInString(i.(string))
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
	widths := t.GetColumnWidths()
	var divider string
	for _, w := range widths {
		divider += ("+" + strings.Repeat("-", w+2))
	}
	divider += "+"

	// headers
	fmt.Println(divider)
	for idx, col := range t.Columns {
		fmt.Printf("| %-*s ", widths[idx], col.Name)
	}
	fmt.Print("|\n")
	fmt.Println(divider)
	// data
	for _, row := range t.Rows {
		for idx, col := range t.Columns {
			i := row.Data[idx]
			var v string
			switch col.Type {
			case "int":
				v = strconv.Itoa(i.(int))
			case "varchar":
				v = i.(string)
			default:
				v = ""
			}
			fmt.Printf("| %-*s ", widths[idx], v)
		}
		fmt.Print("|\n")
	}
	fmt.Println(divider)
}

func (t *Table) ToJson() string {
	var data []map[string]any
	for _, row := range t.Rows {
		m := make(map[string]any)
		for idx, col := range t.Columns {
			m[col.Name] = row.Data[idx]
		}
		data = append(data, m)
	}
	j, _ := json.Marshal(data)
	return string(j)
}

func main() {

	// create
	t := Table{Name: "Test"}
	// add columns
	t.AddColumn("ID", "int")
	t.AddColumn("Name", "varchar")
	// add data
	t.AddRow([]any{4645746, "D"})
	t.AddRow([]any{5, "Eagles"})
	t.AddRow([]any{62, "Fish"})
	// inspect
	fmt.Printf("Table name: %s\n", t.Name)
	fmt.Printf("Columns: %d\n", len(t.Columns))
	fmt.Printf("Rows: %d\n", len(t.Rows))
	fmt.Printf("IDs: %v (longest = %d)\n", t.GetValues("ID"), t.GetColumnWidth("ID"))
	fmt.Printf("Names: %v (longest = %d)\n", t.GetValues("Name"), t.GetColumnWidth("Name"))
	// output as text
	t.Print()
	// output as json
	fmt.Println(t.ToJson())

	// another one
	t2 := Table{Name: "MASH"}
	t2.AddColumn("First Name", "varchar")
	t2.AddColumn("Last Name", "varchar")
	t2.AddColumn("Rank", "varchar")
	t2.AddColumn("First Season", "int")
	t2.AddColumn("Last Season", "int")
	t2.AddRow([]any{"Henry", "Blake", "Lt. Col", 1, 3})
	t2.AddRow([]any{"Frank", "Burns", "Maj", 1, 5})
	t2.AddRow([]any{"Margaret", "Houlihan", "Maj", 1, 11})
	t2.AddRow([]any{"Hawkeye", "Pierce", "Capt", 1, 11})
	t2.AddRow([]any{"Trapper John", "MacIntyre", "Capt", 1, 3})
	t2.AddRow([]any{"Radar", "O'Reilly", "Cpl", 1, 8})
	t2.AddRow([]any{"Max", "Klinger", "Sgt", 1, 11})
	t2.AddRow([]any{"BJ", "Hunnicut", "Capt", 4, 11})
	t2.AddRow([]any{"Sherman", "Potter", "Col", 4, 11})
	t2.AddRow([]any{"Charles", "Winchester", "Maj", 6, 11})
	t2.Print()
	fmt.Println(t2.ToJson())
}
