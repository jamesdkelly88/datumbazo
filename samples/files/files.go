package main

import (
	"fmt"
	"io"
	"os"
)

// streamlining error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	ReadFileIntoMemory("data.csv")
	ReadBytesFromFile("data.csv", 10)
	ReadBytesFromFile("data.csv", 100)
	ReadBytesFromFileOffset("data.csv", 10, 0)
	ReadBytesFromFileOffset("data.csv", 10, 5)
}

func ReadFileIntoMemory(path string) {
	fmt.Printf("\nReading %s with os.ReadFile\n\n", path)
	dat, err := os.ReadFile(path)
	check(err)
	fmt.Print(string(dat))
}

func ReadBytesFromFile(path string, number int) {
	fmt.Printf("\nReading %d bytes from %s with os.Open and File.Read\n\n", number, path)
	f, err := os.Open(path)
	check(err)

	b1 := make([]byte, number)
	n1, err := f.Read(b1)
	check(err)

	f.Close()

	fmt.Printf("Read %d bytes:\n%s\n", n1, string(b1[:n1]))

}

func ReadBytesFromFileOffset(path string, number int, offset int64) {
	fmt.Printf("\nReading %d bytes from %s starting at %d with os.Open, io.SeekStart and File.Read\n\n", number, path, offset)
	f, err := os.Open(path)
	check(err)

	o1, err := f.Seek(offset, io.SeekStart) // io.SeekCurrent uses previous position, io.SeekEnd uses end of file (offset must be negative)
	check(err)

	b1 := make([]byte, number)
	n1, err := f.Read(b1)
	check(err)

	f.Close()

	fmt.Printf("Read %d bytes from offset %d:\n%s\n", n1, o1, string(b1[:n1]))

}
