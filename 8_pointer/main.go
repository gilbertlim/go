package main

import "fmt"

// &variable : variable's memory address
// *variable : hold memory address and resolves it

func main() {
    var a = 5
    var b = &a
    fmt.Printf("Address of var a: %p\n", &a) // 0xc0000180f0
    fmt.Printf("Value of var a: %v\n", *b) // 5
    fmt.Println(*&a) // 5
    fmt.Println(a) // 5
}

// https://stackoverflow.com/questions/38172661/what-is-the-meaning-of-and