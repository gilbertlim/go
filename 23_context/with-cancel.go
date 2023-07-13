package main

import (
	"context"
	"time"
)

/*
* WithCancel
* context가 cancel or timeout으로 종료되면 context의 Done이 호출됨
 */

func cancel() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	go PrintTick(ctx)

	time.Sleep(5 * time.Second)
	cancel() // context의 cancel을 호출하여 context를 종료시킴

	wg.Wait()
}

//tick
//tick
//tick
//tick
//tick
//Done: context canceled
