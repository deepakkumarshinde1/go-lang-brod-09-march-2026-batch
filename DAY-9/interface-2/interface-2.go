package main

import "fmt"

type Shape interface {
	Area() string
}
type Rectangle struct {
	height float32
	width  float32
}

func (r Rectangle) Area() string {
	return fmt.Sprintf(" The area of Rectangle is %v", r.height*r.width)
}

type Circle struct {
	rad float32
}

func (c Circle) Area() string {
	return fmt.Sprintf("Area of circle is = %v", 3.14*c.rad*c.rad)
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}
func main() {
	rect := Rectangle{
		width:  100,
		height: 20,
	}
	circle := Circle{
		rad: 30,
	}
	printArea(rect)
	printArea(circle)
}
