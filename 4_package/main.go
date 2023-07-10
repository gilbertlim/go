package main

// main package는 go compiler에 의해 특별하게 인식됨
// 패키지 명이 main인 경우, go compiler가 패키지를 executble program으로 만듬
// main 패키지 안의 main() 함수가 Entry Point가 됨

import (
	// go compiler는 패키지를 import할 때 패키지 종류에 따라 GOROOT or GOPATH 환경 변수를 검색함
	// GOROOT/pkg: 표준 패키지
	// GOPATH/pkg: 사용자 or 3rd party 패키지

	"fmt"
	
	subAlias "package_test/sub" // <alias> "<module_name>/<package_name>"

	"github.com/google/uuid" // external package

	_ "package_test/init" // 패키지의 init() 함수만 호출하는 경우 _ 사용
)

func main() {
	fmt.Println("sub package의 Test() 함수를 실행합니다.")
	subAlias.Test() // public


	fmt.Println("외부 패키지 uuid의 New() 함수를 실행합니다.")
 	fmt.Println(uuid.New()) // public

	
	// sub.test3() // 함수가 소문자로 시작할 경우 실행 불가 (non public)



}