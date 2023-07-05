package main

import "fmt"

func main() {
    a := 1
    b := 2
    var c = (a + b) / 5;
    c++;
    fmt.Println(c)
    fmt.Println(a == b)
    fmt.Println(a != b)
    fmt.Println(a <= b)
    fmt.Println(a != b && a <= b)
    fmt.Println(a != b || a > b)
    fmt.Println(!true)
    a += 3
    fmt.Println(a)
}