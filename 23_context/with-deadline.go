package main

import (
	"context"
	"time"
)

/*
* WithDeadline
* 작업 흐름(goroutine)을 "언제까지" 유지할지 결정할 때 사용
 */

func deadline() {
	wg.Add(1)

	d := time.Now().Add(3 * time.Second)                         // 3초 설정
	ctx, cancel := context.WithDeadline(context.Background(), d) // 3초 후 컨텍스트가 종료되도록 데드라인 지정

	go PrintTick(ctx)

	time.Sleep(time.Second * 5)
	cancel()
	// 3초 deadline으로 인해 3초 이후 컨텍스트 종료
	// main goroutine은 5초 후에 프로그램이 종료됨
	// deadline으로 컨텍스트를 종료시키더라도 cancel() 함수는 필수

	wg.Wait()
}

//tick
//tick
//tick
//Done: context deadline exceeded
