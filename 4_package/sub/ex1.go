package sub

import "fmt"

func Test() {
	fmt.Println("this is ex1 in sub")
	Test2() // 같은 패키지의 함수는 바로 사용
}