package calc_test

import (
	"calc"
	"testing"
)

// 1. 패키지명을 다르게 설정
// 2. 파일 명을 *_test.go로 설정

func TestSum(t *testing.T) { // 예약된 메서드명
	s := calc.Sum(1, 2, 3)
	if s != 6 {
		t.Error("Wrong Result")
	}
}

/*
=== RUN   TestSum
--- PASS: TestSum (0.00s)
PASS
*/
