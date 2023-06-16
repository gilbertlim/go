package main // 소속된 패키지 명
// main package : 프로그램의 시작 패키지, return 없는 main() 함수 필요

import (
	"fmt" // formatted I/O Library, C언어의 stdio와 같은 표준 입출력 라이브러리
	"strconv"
)

func main() { // 메인 함수
	fmt.Println("Hello, Word!")

	variable()
	cast()
}

func variable() {
	var a int
	var f float32 = 11.
	a = 10
	f = 12.0
	fmt.Println(a, f)


	var i, j, k int = 1, 2, 3
	fmt.Println(i, j, k)

	var l = 1
	var m = "Hi"
	fmt.Println(l, m)

	const c int = 10
	const s string = "Hi"
	const (
		Apple = iota // 0
		Grape
		Orange
	)
	fmt.Println(c, s)
	fmt.Println(Apple, Grape, Orange)

	notype := 1
	fmt.Println("notype", notype)
}

func cast() {
	var v1 int8 = 10
	var v2 uint64

	v2 = uint64(v1)
	fmt.Println(v2)


	i, err := strconv.Atoi("-42")
	i = i + 2
	fmt.Println(i, err)

	j := strconv.Itoa(-42)
	j = j + "2"
	fmt.Println(j, err)

	f, err := strconv.ParseFloat("3.1415", 64)
	f = f + 0.8585
	fmt.Println(f)
}