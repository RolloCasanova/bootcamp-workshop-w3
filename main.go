package main

import (
	"fmt"

	"github.com/RolloCasanova/bootcamp-challenge-w3/shape"
)

func main() {
	// create a new Square, Circle and Hexagon
	square := shape.NewSquare("S", "Square", 4)
	circle := shape.NewCircle("C", "Circle", 10.2)
	hexagon := shape.NewHexagon("H", "Hexagon", 5)

	// print shapes
	printShape(square)
	printShape(circle)
	printShape(hexagon)

	// change shape's value
	square.NewName("Bob")
	square.NewShapeType("Tilted rombus")
	square.NewSide(8.88)

	circle.NewName("CJ")
	circle.NewShapeType("Circulo")
	circle.NewRadius(1.01)

	hexagon.NewName("0x00")
	hexagon.NewShapeType("Bee's favorite shape")
	hexagon.NewSide(0.25)

	// print shapes again
	printShape(square)
	printShape(circle)
	printShape(hexagon)
}

func printShape(s shape.Shape) {
	fmt.Println(s.Info())
}
