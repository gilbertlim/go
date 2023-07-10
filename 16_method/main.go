package main

import "fmt"

// Go는 OOP를 지원하지만,
// 클래스가 필드, 메서드를 함께 갖는 것과는 다른 방식이다.

// struct에는 필드만
type Rect struct {
    width, height int
}

// 어떤 struct를 위한 메서드인지 표시하여 사용(Receiver로 불림)
func (r Rect) area() int { // func (<alias> <struct>) <method_name> <return type>
    return r.width * r.height
}

// pointer receiver
func (r *Rect) area2() int {
    r.width++
    return r.width * r.height
}

func main() {
    rect := Rect{10, 20}
    area := rect.area()
    fmt.Println(area)

    area2 := rect.area2()
    fmt.Println(rect.width, area2) // 11 120

    area3 := rect.area()
    fmt.Println(area3)
}