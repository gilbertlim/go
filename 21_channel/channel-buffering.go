package main

import "fmt"

/* Channel Buffering
* Unbuffered Channel: 수신자가 데이터를 받을 때까지 송신자가 데이터를 보내는 채널에 묶여 있음
* Buffered Channel: 수신자가 받을 준비가 되어있지 않더라도 지정된 버퍼만큼 데이터를 보내고 계속 다른 일을 수행할 수 있음
*/

func buffering() {
    // c := make(chan int)
    // c <- 1
    // fmt.Println(<-c) 
    // fatal error: all goroutines are asleep - deadlock!
    // main routine에서 channel에 값을 보내고, 수신자를 기다리지만 수신자 goroutine이 없기 때문에 에러 발생

    ch := make(chan int, 1) // buffer 개수
    ch <- 101 // 수신자가 없더라도 데이터를 보낼 수 있음
    fmt.Println(<-ch) 

}