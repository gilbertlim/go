package main

import (
	"fmt"
	"sync"
)

func anonymous() {
    // wait group
    // 여러 goroutine이 끝날 때까지 기다리는 역할
    var wait sync.WaitGroup
    wait.Add(2) // 몇 개의 goroutine을 기다릴 것인지 결정

    go func() { // anonymous
        defer wait.Done() // goroutine이 끝나면 Done() 호출
        fmt.Println("Hello")
    }()

    go func(msg string) {
        defer wait.Done()
        fmt.Println(msg)
    }("Hi")

    wait.Wait() // goroutine이 모두 끝날 때 까지 대기
}