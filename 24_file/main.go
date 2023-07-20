package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fi, err := os.Open("./test.txt")
	if err != nil {
		panic(err) // 현재 함수를 즉시 멈추고 defer 함수를 모두 실행한 후 즉시 리턴
	}
	defer fi.Close()

	fo, err := os.Create("./test2.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 1024)

	for {
		cnt, err := fi.Read(buff) // 파일을 buff 변수 크기만큼 읽어 buff에 저장
		if err != nil && err != io.EOF {
			panic(err)
		}

		fmt.Println(buff)

		// [116 101 115 116 0 0 0 ... ]

		if cnt == 0 {
			break
		}

		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}

	// I/O 패키지
	// 입력 파일이 크기 않은 경우 편하게 사용 가능
	ioUtil()
}

func ioUtil() {

	bytes, err := ioutil.ReadFile("./io.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(bytes)

	err = ioutil.WriteFile("./io2.txt", bytes, 0666)
	// 파일이 존재하지 않을 경우 permission denied 발생하므로 권한 부여
	if err != nil {
		panic(err)
	}
}
