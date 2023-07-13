package main

import (
	"context"
	"fmt"
)

/*
* WithValue
* channel과 같이 sub goroutine에 데이터를 전달할 수 있음
 */

func value() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), "v", 3)

	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("v"); v != nil {
		n := v.(int)
		fmt.Println("Square:", n*n)
	}
	wg.Done()
}

// Square: 9
