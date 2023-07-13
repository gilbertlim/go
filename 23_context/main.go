package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
* Context Package
* 작업 명세서
* 작업 가능 시간, 작업 취소 등 작업의 흐름 제어
* 새로운 goroutine으로 작업을 시작할 때 일정 시간 동안만 작업을 지시하거나 외부에서 작업을 취소할 때 사용
 */

var wg sync.WaitGroup

func main() {
	cancel()
	deadline()
	timeout()
	value()
	wrapping()
}

// PrintTick 매 초 "tick" 문자를 화면에 표시하는 함수
// Done 함수 채널에 데이터가 들어오면 종료
func PrintTick(ctx context.Context) {
	tick := time.Tick(time.Second) // Tick 함수로 생성한 채널
	for {
		select {
		case <-ctx.Done(): // Done 함수 channel에서 데이터가 들어오면
			fmt.Println("Done:", ctx.Err())
			wg.Done()
			return
		case <-tick: // tick 채널에 데이터가 들어오면
			fmt.Println("tick")
		}
	}
}

//tick
//tick
//Done: context canceled
