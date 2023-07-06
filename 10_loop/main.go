package main

import "fmt"

func main() {
    
    // basic
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)


    // only condition
    n := 1
    for n < 100 { // while 같다.
        n *= 2
    }
    fmt.Println(n)


    // range
    names := []string{"a", "b", "c"}
    for index, name := range names {
        fmt.Println(index, name)
    }


    // break, continue, goto
    a := 1
    for a < 15 {
        if a == 5 {
            a += a
            continue
        }

        a++

        if a > 10 {
            break
        }
    }
    fmt.Println(a)
    if a == 11 {
        goto END
    }
// label
END:
    fmt.Println("END")


    // break label
    i := 0
L1:
    for { // infinite loop
        fmt.Println("i:", i)
        if i == 5 {
            break L1
        }

        i++
    }
    fmt.Println("break L1, i=", i)
}
