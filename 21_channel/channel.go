package main

import "fmt"

/*
* Go Channel
* 데이터를 주고 받는 통로
* make() 함수로 미리 생성해야 함
* 채널 연산자(<-)를 통해 데이터를 보내고 받음
* goroutine 사이에서 데이터를 주고 받는데 사용됨
* 상대편이 준비될 때까지 Channel에서 대기함으로 써 별도의 Lock을 걸지 않고 데이타를 동기화하는데 사용
*/
 
func channel() {
    // int channel 생성
    ch := make(chan int)   

    go func() {
        ch <- 123 // channel에 값을 보냄
    }()

    var i int = <- ch // channel로부터 값을 받음
    fmt.Println(i)

    

    done := make(chan bool)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
        done <- true
    } ()

    // goroutine이 끝날 때 까지 대기(데이터를 보내줄 때까지)
    <-done 
}