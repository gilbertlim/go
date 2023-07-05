package main // 소속된 패키지 명
// main package : 프로그램의 시작 패키지, return 없는 main() 함수 필요

import (
	"fmt" // formatted I/O Library, C언어의 stdio와 같은 표준 입출력 라이브러리
)

func main() { // 메인 함수
	fmt.Println("Hello, Word!")
}