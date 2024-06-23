package main

import (
	"fmt"
	"math"
)

func main() {

	a := 50
	b := 10

	// swap(&a, &b)

	// fmt.Println(a)
	// fmt.Println(b)

	fmt.Println("Max : ", max(a, b))
	fmt.Println("Sum : ", sum(a, b))
	fmt.Println("IsEven : ", isEven(a))
	fmt.Println("divide : ", divide(a, b))
	fmt.Println("power : ", power(a, b))

}

func swap(a, b *int) {
	olda := *a
	*a = *b
	*b = olda
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b

}

func sum(a int, b int) int {
	sum := a + b
	return sum
}

func isEven(a int) bool {
	if evenCheck := a % 2; evenCheck == 0 {
		return true
	}
	return false
}

func divide(a int, b int) float64 {
	divder := (float64(a)) / float64(b)
	return divder
}

func power(base int, exponent int) int {
	powerer := math.Pow(float64(base), float64(exponent))
	return int(powerer)
}
