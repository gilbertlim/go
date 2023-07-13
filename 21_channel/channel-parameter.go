package main

import "fmt"

/*
* Channel Parameter
* channel로 데이터를 송신만 할 것인지, 수신만 할 것인지 지정 가능
 */

func parameter() {
    ch := make(chan string, 1)
    sendChan(ch)
    receiveChan(ch)
}

func sendChan(ch chan<- string) {
    ch <- "Data"
    // test := <- ch // send-only channel error
}

func receiveChan(ch <-chan string) {
    data := <- ch
    fmt.Println(data)
}