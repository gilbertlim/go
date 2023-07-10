package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    f, err := os.Open("./test.sh")

    switch err.(type) {
        default:
            fmt.Println("ok")
            fmt.Println(f.Name())
        case error: // error interface
            log.Fatal(err.Error())  // Fatal: 에러 메시지 출력 및 os.Exit(1)를 호출하여 프로그램 종료
    }
    

}