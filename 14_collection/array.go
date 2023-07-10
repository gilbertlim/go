package main

import "fmt"

func arrayTest() {
    // 배열 선언
    var a [3]int
    a[0] = 1
    a[1] = 2
    a[2] = 3
    fmt.Println(a[1])


    // 배열 초기화
    var b = [3]int{1, 2, 3}
    var c = [...]int{1, 2, 3} // 배열 크기 자동화
    fmt.Println(b)
    fmt.Println(c)


    // 다차원 배열
    var d = [2][3]int{
        {1, 2, 3},
        {4, 5, 6},
    }
    fmt.Println(d)
}