package main

import (
	"fmt"
	"math"
)

type Rect struct {
    width, height float64
}

// Rect type에 대한 Shape interfafce 구현
func (r Rect) area() float64 {
    return r.width * r.height
}

func (r Rect) perimeter() float64 {
    return 2 * (r.width + r.height)
}

type Circle struct {
    radius float64
}

// Circle type에 대한 Shape interfafce 구현
func (c Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 { 
    return 2 * math.Pi * c.radius
}    


func main() {
    // inteface
    r := Rect{10., 20.}
    c := Circle{10}
    
    showArea(r)
    showArea(r, c)


    // empty inteface
    var x interface{}
    x = 1
    x = "Tom"
    printIt(x)

    
    // type assertion
    var y interface{} = 5
    i := y
    j := y.(int)

    fmt.Println(i) // 5
    fmt.Println(j) // 5
    // y가 nil이 아니고 int 타입에 속한다면, 그 값을 리턴
    // nil이거나 속하지 않는 다면 RuntinmeError 발생
}


func showArea(shapes ...Shape) {
    for _, s := range shapes {
        a := s.area()
        p := s.perimeter()
        fmt.Println(a, p)
    }
}


// empty interface
// 메서드를 전혀 갖지 않는 빈 인터페이스
// 어떤 타입도 담을 수 있음
func printIt(v interface{}) {
    fmt.Println(v)
}