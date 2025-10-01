package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	http.HandleFunc("/", data)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func data(w http.ResponseWriter, r *http.Request) {
	// v1 - read file and print directly to response
	dat, err := os.ReadFile("data.csv")
	check(err)
	fmt.Fprintln(w, "v1")
	fmt.Fprintf(w, "%s\n", dat)

	// v2 - load file as csv
	f, err := os.Open("data.csv")
	check(err)
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	check(err)
	fmt.Fprintln(w, "v2")
	fmt.Fprintf(w, "%s\n", data)

	// v3 - load into object list
	obj := convertToObjects(data)
	fmt.Fprintln(w, "v3")
	for _, o := range obj {
		fmt.Fprintf(w, "A = %d, B = %d\n", o.A, o.B)
	}

	// v4 - convert to json
	jsonData, err := json.MarshalIndent(obj, "", "  ")
	check(err)
	fmt.Fprintln(w, "v4")
	fmt.Fprintf(w, "%s\n", jsonData)

	// TODO: try to use dynamic parsing instead of struct
	// https://stackoverflow.com/questions/20768511/unmarshal-csv-record-into-struct-in-go
}

type demo struct {
	A int
	B int
}

func convertToObjects(data [][]string) []demo {
	var list []demo
	for i, line := range data {
		if i > 0 { // skip headers
			var d demo
			d.A, _ = strconv.Atoi(line[0])
			d.B, _ = strconv.Atoi(line[1])
			list = append(list, d)
		}
	}
	return list
}

/* output:

v1
A,B
1,4
2,5
3,6
v2
[[A B] [1 4] [2 5] [3 6]]
v3
A = 1, B = 4
A = 2, B = 5
A = 3, B = 6
v4
[
  {
    "A": 1,
    "B": 4
  },
  {
    "A": 2,
    "B": 5
  },
  {
    "A": 3,
    "B": 6
  }
]
*/
