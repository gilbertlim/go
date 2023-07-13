package main

import (
	"fmt"
	"time"
)

/*
* Channel Select
* 여러 개의 채널을 기다리면서 준비된(데이터를 보내온) 채널을 실행하는 기능
 */

func channelSelect() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			fmt.Println("run1 완료")

		case <-done2:
			fmt.Println("run2 완료")
			break EXIT
			// EXIT label로 이동후 자신이 빠져나온 루프의 다음 문장을 실행함
			// 없기 때문에 main() 함수 종료
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
