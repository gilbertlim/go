package main

import (
	"fmt"
	"strconv"
)

func typeConversion() {
	fmt.Println("type conversion")

	var v1 int8 = 10

	var v2 = uint64(v1)
	fmt.Println(v2)


	i, err := strconv.Atoi("-42")
	i = i + 2
	fmt.Println(i, err)

	j := strconv.Itoa(-42)
	j = j + "2"
	fmt.Println(j, err)

	f, err := strconv.ParseFloat("3.1415", 64)
	f = f + 0.8585
	fmt.Println(f, err)

	var g = "test"
	bytes := []byte(g)
	str := string(bytes)
	fmt.Println(bytes, str)
}