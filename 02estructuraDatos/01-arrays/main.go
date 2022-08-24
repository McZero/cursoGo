package main

import "fmt"

func main() {
	// arrays
	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	fmt.Println(numeros)
	fmt.Println(numeros[1])

	// slicen
	digitos := []int{1, 2, 3}
	fmt.Println(digitos, len(digitos), cap(digitos))

	digitos = append(digitos, 4, 5, 6)
	fmt.Println(digitos, len(digitos), cap(digitos))

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(digitos), cap(digitos), digitos)

	// iterar colecciones
	for _, elemento := range digitos {
		fmt.Println(elemento)
	}
}
