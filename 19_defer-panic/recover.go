package main

import (
	"fmt"
	"os"
)

func Recover() {
    openFile2("Invalid.txt")

    // recover에 의해 문장이 실행됨
    fmt.Println("Done")
}

func openFile2(fn string) {
    // defer
    // panic 호출시 실행됨
    defer func() {
        // recover
        // panic 상태를 제거하여 정상 상태로 되돌림 (다음 문장 계속 실행)
        if r := recover(); r != nil {
            fmt.Println("OPEN ERROR", r)
        }
    }()
 
    f, err := os.Open(fn)
    if err != nil {
        panic(err)
    }
 
    defer f.Close()
}