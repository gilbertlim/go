package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {
	member := Member{"Alex", 10, true}

	// encoding
	xmlBytes, err := xml.Marshal(member)
	if err != nil {
		panic(err)
	}

	xmlString := string(xmlBytes)
	fmt.Println(xmlString)

	// decoding
	var mem Member
	err2 := xml.Unmarshal(xmlBytes, &mem)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(mem.Name, mem.Age, mem.Active)

	// Read
	fp, err := os.Open("./test.xml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	data, err := io.ReadAll(fp)

	var members Members
	xmlerr := xml.Unmarshal(data, &members)
	if xmlerr != nil {
		panic(xmlerr)
	}

	fmt.Println(members)

}

type Member struct {
	Name   string
	Age    int
	Active bool
}

type Members struct {
	Member []Member
}
