package main

import (
	"fmt"
)

func main() {
	algo := "Algoritmos"
	fmt.Println("Hola Mundo!")
	fmt.Printf("Esto es %s %d\n", algo, 2)

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

	//fmt.Println("Suma:", sumar(3, 1))

	suma, resta := sumarYRestar(3, 4)

	fmt.Println("Suma:", suma)
	fmt.Println("Resta:", resta)

	// For - Ejemplo

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}

func sumar(a int, b int) int {
	return a + b
}

// func sumarYRestar(a, b int) (int, int) {
// 	return a+b, a-b
// }

// func sumarYRestar(a, b int) (int, int) {
// 	var suma, resta int
// 	suma = a+b
// 	resta = a-b
// 	return suma, resta
// }

func sumarYRestar(a, b int) (suma, resta int) {
	suma = a + b
	resta = a - b
	return
}
