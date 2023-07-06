package main

import "fmt"

func ifStatement() {
    k := 3

    if k == 1 {
        fmt.Println(k)
    } else if k == 2 {
        fmt.Println(k)
    } else {
        fmt.Println("else")
    }

    if val := k * 2; val < 9 {
        fmt.Println(val)
    }
}