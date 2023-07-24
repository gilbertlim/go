package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	member := Member{"Alex", 10, true}

	// encoding
	jsonBytes, err := json.Marshal(member)
	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	// decoding
	// 매칭되는 필드가 구조체에 없으면 값은 무시되고 처리됨
	var mem Member
	err2 := json.Unmarshal(jsonBytes, &mem)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(mem.Name, mem.Age, mem.Active)
}

type Member struct {
	Name   string
	Age    int
	Active bool
}
