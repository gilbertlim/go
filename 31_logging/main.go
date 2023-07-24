package main

import (
	"io"
	"log"
	"os"
)

var myLogger *log.Logger

func main() {
	// default logger
	log.SetFlags(3) // 날짜 시간 형식 설정
	// 0: logging
	// 1: 2023/07/24 logging
	// 2: 16:02:12 logging
	// 3: 2023/07/24 16:02:37 logging
	log.Println("logging")

	// custom logger
	// 로그 플래그를 활용하여 지정된 플래그 표시 가능
	//myLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	//myLogger.Println("myLogger")
	// INFO: 2023/07/24 16:10:15 /Users/mz01-lsjin/Study/go/31_logging/calc.go:23: myLogger

	// file logging
	fpLog, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()
	//myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	//myLogger.Println("fileLogger")

	// file logging - file logger(custom)를 defaultLogger로 설정
	//log.SetOutput(fpLog)
	//log.Println("fileLoggerWithDefaults")

	// multi writer: console과 file 모두 로그 출력
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	log.SetOutput(multiWriter)
	log.Println("multi writer")

}
