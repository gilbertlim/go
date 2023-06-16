package main

import (
	"fmt"
	"package_test/sub" // <module_name>/<package_name>

	"github.com/google/uuid" // external package
)

func main() {
	fmt.Println("sub package의 Test() 함수를 실행합니다.")
	sub.Test()

	fmt.Println("외부 패키지 uuid의 New() 함수를 실행합니다.")
 	fmt.Println(uuid.New())
}