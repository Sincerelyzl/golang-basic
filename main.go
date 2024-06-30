package main

import (
	"interface/creature"
)

func main() {
	// rectangle := shape.Rectangle{Width: 10, Length: 20}
	// triangle := shape.Triangle{Base: 10, Height: 20}
	// circle := shape.Circle{Radius: 10}
	// rectangle.Area()
	// fmt.Println("Area of rectangle is : ", shape.CalArea(rectangle))
	// fmt.Println("Area of triangle is : ", shape.CalArea(triangle))
	// fmt.Println("Area of circle is : ", shape.CalArea(circle))

	// shape.GetInfo(rectangle)
	// shape.GetInfo(triangle)
	// shape.GetInfo(circle)
	intAdd := 5000

	player1 := creature.NewPlayer(0.0)
	player2 := creature.NewPlayer(intAdd)

	// player1.ShowLifePoint()
	// player2.ShowLifePoint()

	creature.AttackEvent(&player1, &player2, 1000)

}
