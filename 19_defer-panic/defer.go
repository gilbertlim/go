package main

import (
	"fmt"
	"os"
)

func Defer() {
    f, err := os.Open("./defer.txt")
    if err != nil {
        panic(err)
    }

    // defer
    // 특정 문장 혹은 함수를 나중에(defer를 호출하는 함수, 여기서는 main()이 리턴하기 직전에) 실행하게 함
    // main 마지막에 실행
    defer f.Close()

    bytes := make([]byte, 1024)
    f.Read(bytes)
    fmt.Println(len(bytes))
}