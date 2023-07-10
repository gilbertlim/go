package main

import "fmt"

func main() {
 
    // closure: 함수 바깥에 있는 변수를 참조하는 함수 값
    // 함수 바깥의 변수를 함수 안으로 끌어들인 듯이 변수를 읽음
    next := nextValue()

    fmt.Println(next())
    fmt.Println(next())
    fmt.Println(next())

    anotherNext := nextValue()
    fmt.Println(anotherNext())
    fmt.Println(anotherNext())
}

func nextValue() func() int {
    i := 0

    // func() int 익명함수를 리턴
    return func() int {
        i++ // 익명 함수 밖의 i를 참조하고 있음
        return i
    }
}