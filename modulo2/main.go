package main

import "fmt"

func main() {
	var dolar int = 10
	fmt.Printf("Ahora vale: %d\n", dolar)
	pasaUnDia(&dolar)
	fmt.Printf("Ahora vale: %d", dolar)

}

func pasaUnDia(valor *int) {
	*valor = 2 * *valor
}
