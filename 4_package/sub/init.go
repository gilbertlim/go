package sub

import "fmt"

func init() { // 패키지 실행 시 처음으로 호출되는 함수
	fmt.Println("sub package의 init 함수가 호출되었습니다.")
}