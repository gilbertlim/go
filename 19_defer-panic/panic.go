package main

import (
	"fmt"
	"os"
)

func Panic() {
    openFile("Invalid.txt")

    // panic 동작 시 아래 문장 실행 X
    fmt.Println("Done")
}

func openFile(fn string) {
    f, err := os.Open(fn)
    if err != nil {
        // panic
        // 현재 함수를 즉시 멈추고 defer 함수를 모두 실행한 후 즉시 리턴
        panic(err)
    }

    defer f.Close()
}