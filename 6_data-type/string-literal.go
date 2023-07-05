package main

import "fmt"

func stringLiteral() {
    fmt.Println("string literal")

    literal := `안녕
하
십\n
니까`

    doubleQuote := "안녕하십\n니까"

    fmt.Println(literal)
    fmt.Println()
    fmt.Println(doubleQuote)
    fmt.Println()
}
