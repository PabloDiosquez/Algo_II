package main

import "fmt"

func main() {
	fmt.Println("Hola Mundo!")

	a, b := 5, 9
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	a, b = b, a

	var (
		x int     = 10
		y float64 = 12.45
		z uint    = 3
	)

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)

}
