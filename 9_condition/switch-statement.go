package main

import "fmt"

func switchStatement() {
    switch1()
    switch2()

    checkType(20)
    checkType("20")

    fall(2)
    
    
    

    // // Expression
    // switch x := category << 2; x - 1 {

    // }

    
}

func switch1() {
    var name string
    var category int = 1

    switch category {
        case 1:
            name = "Paper Book"
            fmt.Println(name)
        case 2:
            name = "eBook"
            fmt.Println(name)
        case 3, 4:
            name = "Blog"
            fmt.Println(name)
        default:
            name = "Other"
            fmt.Println(name)
    }
}

func switch2() {
    var name string
    var category int = 1

    switch {
    case category == 1:
        name = "Paper Book"
        fmt.Println(name)
    case category == 2:
        name = "eBook"
        fmt.Println(name)
    case category == 3 || category == 4:
        name = "Blog"
        fmt.Println(name)
    default:
        name = "Other"
        fmt.Println(name)
}
}

func checkType(i interface{}) {
    switch i.(type) {
    case int:
        fmt.Println("int")
    case string:
        fmt.Println("string")
        
    }
}

func fall(val int) {
    // golang은 항상 break 한다.
    // break 하지 않기 위해 fallthrough를 사용
    switch val {
    case 1:
        fmt.Println("1 이하")
        fallthrough
    case 2:
        fmt.Println("2 이하")
        fallthrough
    case 3:
        fmt.Println("3 이하")
        fallthrough
    default:
        fmt.Println("default")

    }
}