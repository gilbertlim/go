package main

import "fmt"

/*
* Channel Range
* 수신자가 계속 루프를 돌면서 채널이 닫혔는 지 체크할 수 있음
 */

func channelRange() {
	fmt.Println("channel-range")
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	close(ch)

	/*for {
		if i, success := <-ch; success {
			fmt.Println(i)
		} else {
			break
		}
	}*/

	for i := range ch {
		fmt.Println(i)
	}
}
