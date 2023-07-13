package main

import "runtime"

func multipleCPU() {
    // default: 1 CPU
    // 여러 개의 goroutine을 만들어도 1개 CPU에서 작업을 시분할하여 처리함(concurrent)
    // 여러 개 CPU를 가지는 경우 다중 CPU에서 Parallel 처리할 수 있음
    // 아래 함수에 CPU 개수를 인자로하여 호출
    runtime.GOMAXPROCS(4) 
}