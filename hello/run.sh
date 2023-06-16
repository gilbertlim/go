#!/bin/bash

# main package 파일을 컴파일 후 에 바이너리 파일을 남기지 않고, 메모리 위에서 실행
go run hello.go

# main 패키지 파일을 import 시킨 모든 라이브러리, 연결된 패키지와 함께 컴파일하며 바이너리 파일을 생성
go build hello.go
# 바이너리 파일 실행
./hello
