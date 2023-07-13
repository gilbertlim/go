package main

import (
	"fmt"
	"time"
)

/* goroutine
Go 런타임이 관리하는 Lightweight 논리적 Thread
'go' 키워드를 사용하여 함수를 호출하면, runtime 시 새로운 goroutine을 실행함
비 동기적으로 함수 루틴을 실행하기 때문에 코드를 동시에 실행하는데 사용됨

(참고)
goroutine은 OS 쓰레드보다 훨씬 가볍게 비동기 Concurrent 처리를 구현하기 위하여 만든 것으로, 기본적으로 Go 런타임이 자체 관리한다.
Go 런타임 상에서 관리되는 작업단위인 여러 goroutine들은 종종 하나의 OS 쓰레드 1개로도 실행되곤 한다.
즉, Go루틴들은 OS 쓰레드와 1 대 1로 대응되지 않고, Multiplexing으로 훨씬 적은 OS 쓰레드를 사용한다.
메모리 측면에서도 OS 쓰레드가 1 메가바이트의 스택을 갖는 반면, goroutine은 이보다 훨씬 작은 몇 킬로바이트의 스택을 갖는다(필요시 동적으로 증가).
Go 런타임은 Go루틴을 관리하면서 Go 채널을 통해 Go루틴 간의 통신을 쉽게 할 수 있도록 하였다.
*/

func main() {

    say("Sync")

    go say("Async1")
    go say("Async2")
    go say("Async3")

    time.Sleep(time.Second * 3)

    anonymous()

    multipleCPU()
}

func say(s string) {
    for i := 0; i < 10; i++ {
        fmt.Println(s, "***", i)
    }
}