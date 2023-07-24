package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// reader
	f, _ := os.Open("./test.csv")

	reader := csv.NewReader(bufio.NewReader(f))

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, record := range records {
		for j := range record {
			fmt.Printf("%s ", records[i][j])
		}
		fmt.Println()
	}

	// writer

	of, oErr := os.Create("./output.csv")
	if oErr != nil {
		panic(oErr)
	}

	wr := csv.NewWriter(bufio.NewWriter(of))
	wr.Write([]string{"name", "city"})
	wr.Write([]string{"jin1", "newyork"})
	wr.Write([]string{"jin2", "florida"})
	//wr.WriteAll()
	wr.Flush()
}
