package main

import "fmt"

// struct
// 필드들의 집합체
type person struct {
    name string
    age int
}

type dict struct {
    data map[int]string
}

func newDict() *dict {
    d := dict{} // 객체 생성
    d.data = map[int]string{} // 필드 초기화
    return &d // 포인터 리턴
}

func main() {
    // struct 선언
    p := person{} // person 객체 생성

    // 필드 값 설정
    p.name = "Lim"
    p.age = 10
    fmt.Println(p)


    // struct 객체 생성
    var p1 = person{"Bob", 20}

    p2 := person{name: "Jack", age: 10}
    
    p3 := new(person) // 객체의 포인터 리턴 (*person)
    p3.name = "Lee"
    p3.age = 5

    fmt.Println(p1)
    fmt.Println(p2)
    fmt.Println(p3)
    fmt.Println(&p3) // address
    fmt.Println(*p3) // value


    // constructor
    dic := newDict() // 생성자 호출
    dic.data[1] = "A"
    fmt.Println(dic)
}