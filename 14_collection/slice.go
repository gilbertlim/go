package main

import "fmt"

func sliceTest() {

    // 슬라이스
    // 배열과 달리 고정된 크기를 미리지정하지 않을 수 있고, 크기 변경 가능
    // 배열의 부분만을 추출 가능
    var a [] int // 슬라이스 선언
    if a == nil {
        fmt.Println("Nil Slice")
    }
    fmt.Println(len(a), cap(a)) // 0 0

    a = []int{1, 2, 3}
    a[1] = 10
    fmt.Println(a)


    // make
    // 슬라이스의 Length(슬라이스 길이), Capacity(내부 배열의 최대 길이) 지정 가능
    s := make([]int, 5, 10) //type, length, capacity
    fmt.Println(len(s), cap(s))


    // sub-slice
    t := []int{0, 1, 2, 3, 4, 5}
    t = t[2:5] // [2 3 4]
    fmt.Println(t)
    t = t[1:] // [3 4]
    fmt.Println(t)
    t = t[:1] // [3]
    fmt.Println(t)


    // append, copy
    u := []int{0, 1}
    u = append(u, 2) // [0 1 2]
    u = append(u, 3, 4, 5) // [0 1 2 3 4 5]
    fmt.Println(u)

    
    // append의 내부적인 로직
    sliceA := make([]int, 0, 3)
    for i := 1; i <= 15; i++ {
        sliceA = append(sliceA, i)
        fmt.Println(len(sliceA), cap(sliceA))
            // 1 3
            // 2 3
            // 3 3
            // 4 6
            // 5 6
            // 6 6
            // 7 12
            // 8 12
            // 9 12
            // 10 12
            // 11 12
            // 12 12
            // 13 24
            // 14 24
            // 15 24
    }
    fmt.Println(sliceA)
    // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]


    // slice 간 append
    sliceB := []int{1, 2, 3}
    sliceC := []int{4, 5, 6}

    sliceB = append(sliceB, sliceC...) // ellipsis(...)는 해당 슬라이스의 collection을 표현
    fmt.Println(sliceB)


    // copy
    source := []int{0, 1, 2}
    fmt.Println(len(source), cap(source)) // 3, 3
    target := make([]int, len(source), cap(source) * 2) // 3, 6
    
    copy(target, source)

    fmt.Println(target)
    fmt.Println(len(target), cap(target))
}