package main

import (
	"context"
	"time"
)

/*
* WithTimeout
* 작업 흐름(goroutine)을 "얼마 동안" 유지할지 결정할 때 사용
 */

func timeout() {
	wg.Add(1)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	go PrintTick(ctx)

	time.Sleep(time.Second * 5)
	cancel()

	wg.Wait()
}

//tick
//tick
//tick
//Done: context deadline exceeded
