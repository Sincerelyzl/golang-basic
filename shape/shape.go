package shape

import (
	"fmt"
	"io"
	"math"
	"os"
)

type Shape interface {
	Area() float64
	Info()
}

// Rectangle struct
type Rectangle struct {
	Width, Length float64
}

// Triangle struct
type Triangle struct {
	Base, Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

// Area method for Triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Info method for Rectangle
func (r Rectangle) Info() {
	fmt.Printf("Type : Rectangle have Length: %f , Width: %f \n", r.Length, r.Width)

}

// Info method for Triangle
func (t Triangle) Info() {
	info := fmt.Sprintf("Type : Triangle have Base: %f ,  Height: %f \n", t.Base, t.Height)
	io.WriteString(os.Stdout, info)
}

// Info method for Circle
func (c Circle) Info() {
	info := fmt.Sprintf("Type : Circle have Radius: %f \n", c.Radius)
	io.WriteString(os.Stdout, info)
}

// CalArea function
func CalArea(s Shape) float64 {
	return s.Area()
}

func GetInfo(s Shape) {
	s.Info()
}
