package main

import "fmt"

func main() {

    // anonymous function
    // 함수 전체를 변수에 할당하거나 다른 함수의 파라미터에 직접 정의되어 사용됨
    sum := func(n  ...int) int {
        s := 0
        for _, i := range n {
            s += i
        }
        return s
    }

    result := sum(1, 2, 3, 4, 5)
    fmt.Println(result)


    // first class function
    // 다른 함수의 파라미터로 전달하거나 다른 함수의 리턴 값으로 사용됨
    add := func(i int, j int) int {
        return i + j
    }

    r1 := calc(add, 10, 20)
    fmt.Println(r1)

    r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
    fmt.Println(r2)

    // type
    r3 := calc2(add, 10, 20)
    fmt.Println(r3)

}

func calc(f func(int, int) int, a int, b int) int {
    result := f(a, b)
    return result
}


// type
type calculator func(int, int) int

func calc2(f calculator, a int, b int) int {
    result := f(a, b)
    return result
}