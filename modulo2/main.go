package main

import "fmt"

func main() {
	// Punteros
	var dolar int = 10
	fmt.Printf("Ahora vale: %d\n", dolar)
	pasaUnDia(&dolar)
	fmt.Printf("Ahora vale: %d\n", dolar)

	// Arreglos
	var arreglo [4]int = [4]int{11, 4, -12, 18}
	fmt.Println(arreglo[0])
	fmt.Println(arreglo[2])
	arreglo[2] = 10
	fmt.Println(arreglo)

	// numeros := [4]int{1, 2, 3}
	// fmt.Println(numeros)

	modificar(&arreglo)
	fmt.Println(arreglo)

	// Slices
	//numeros := [6]int{1,2,3,4,5}
	// var slice []int = numeros[1:3]
	// fmt.Println(slice)

	// Crea un array [5]int y hace un slice := array[:]
	// slice := []int{1,2,3,4,5}
	// fmt.Println(slice)

	slice := make([]int, 0)
	fmt.Println(slice)

}

// pasaUnDia duplica el valor apuntado por el puntero dado.
// Modifica el valor original al multiplicarlo por 2.
//
// Parámetros:
//   valor: Un puntero a un entero cuyo valor se duplicará.
//
// Ejemplo:
//   valor := 5
//   pasaUnDia(&valor)
//   // Ahora valor es 10.
func pasaUnDia(valor *int) {
	*valor = 2 * *valor
}

// modificar modifica el primer elemento del arreglo dado para que sea 0.
//
// Parámetros:
//   arreglo: Un puntero a un arreglo de 4 enteros.
//
// Ejemplo:
//   arr := [4]int{1, 2, 3, 4}
//   modificar(&arr)
//   // Ahora arr es [0 2 3 4].
func modificar(arreglo *[4]int) {
	arreglo[0] = 0
}

func modificar_slice(slice []int) {
	slice[0] = -1
}
