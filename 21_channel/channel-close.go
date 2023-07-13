package main

import "fmt"

/*
* Channel Close
* 송신만 불가, 수신은 가능
 */

func channelClose() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    
    close(ch)

    fmt.Println(<-ch) // 1
    // check(ch) 데이터 있음
    fmt.Println(<-ch) // 2

    check(ch) // 데이터 없음

    
    // 값이 없으면, int channel이므로 0 출력
    fmt.Println(<-ch) // 0
}

func check(ch chan int) {
    if val, success := <-ch; !success {
        fmt.Println("데이터 없음")
    } else {
        fmt.Println(val, success)
    }
}