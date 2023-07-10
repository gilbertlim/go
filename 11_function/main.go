package main

import (
	"fmt"
)

func main() {
    msg := "parameter"
    
    say(msg)
    fmt.Println("main()", msg)
    
    say2(&msg) // 주소 값 전달
    fmt.Println("main()", msg)

    say3("Those", "are", "variadic", "values")

    fmt.Println(say4(1, 2, 3, 4, 5))
    
    fmt.Println(say5(1, 2, 3, 4, 5))

    fmt.Println(say6(1, 2, 3, 4, 5))
}


// Pass By Value
// 값이 함수로 전달됨
// 인자 값은 함수 내에서 변경되더라도 main() 에서는 변함 없음
func say(msg string) {
    fmt.Println("say()", msg)
}


// Pass By Reference
// say2 함수의 인자는 포인터 타입임
// 주소 값을 전달 받고 그 주소에 있는 값을 변경함
func say2(msg *string) {
    fmt.Println("say2()", *msg)
    *msg = "changed parameter"
}


// Variadic Function (가변 인자)
func say3(msg ...string) {
    for _, s := range msg {
        fmt.Println(s)
    }
}


// return
func say4(nums ...int) int {
    s := 0
    for _, n := range nums {
        s += n
    }
    return s
}

// multiple return
func say5(nums ...int) (int, int) {
    sum := 0
    count := 0
    for _, n := range nums {
        sum += n
        count++
    }
    return count, sum
}

// named return
func say6(nums ...int) (count int, sum int) {
    // count = 1
    for _, n := range nums {
        sum += n
        count++
    }
    return
}