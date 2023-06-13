package main // 소속된 패키지 명

import (
	"fmt"
)

func main() { // 메인 함수
	fmt.Println("Hello, Word!")

	variable()
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

}