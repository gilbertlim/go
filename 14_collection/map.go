package main

import "fmt"

func mapTest() {
    // Nil Map
    var idMap map[int]string
    fmt.Println(idMap == nil)
    

    // NotNil Map
    var idMap2 = make(map[int]string)
    fmt.Println(idMap2 == nil)


    // Literal
    tickers := map[string]string {
        "A": "Apple",
        "B": "Banana",
        "C": "Citron",
    }
    fmt.Println(tickers)


    // Update Map
    // var m map[int]string
    var m = make(map[int]string)
    m[901] = "Apple"
    m[134] = "Grape"
    m[777] = "Tomato"

    str := m[134]
    fmt.Println(str)

    noData := m[999]
    fmt.Println(noData) // 공백

    delete(m, 777)
    fmt.Println(m)


    // Key Check
    _, exists := m[988]
    if !exists {
        fmt.Println("This key does not exist")
    }


    // For Loop
    for k, v := range m {
        fmt.Println(k, v)
    }
}