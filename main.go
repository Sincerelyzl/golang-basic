package main

import (
	"fmt"
	"interface/shape"
)

func main() {
	rectangle := shape.Rectangle{Width: 10, Length: 20}
	triangle := shape.Triangle{Base: 10, Height: 20}
	circle := shape.Circle{Radius: 10}

	fmt.Println("Area of rectangle is : ", shape.CalArea(rectangle))
	fmt.Println("Area of triangle is : ", shape.CalArea(triangle))
	fmt.Println("Area of circle is : ", shape.CalArea(circle))

	shape.GetInfo(rectangle)
	shape.GetInfo(triangle)
	shape.GetInfo(circle)

}
