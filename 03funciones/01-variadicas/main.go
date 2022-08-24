package main

import (
	"fmt"
	"strings"
)

// funcion de multiples resultados, con argumentos indeterminados
func sumar(nombre string, numeros ...int) (string, int) {
	var resultado int

	mensaje := fmt.Sprintf("%s, la suma de los numeros %v es:", nombre, numeros)

	for _, numero := range numeros {
		resultado += numero
	}

	return mensaje, resultado
}

// función recursiva
func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}

	f := n * factorial(n-1)

	return f
}

// Closures
func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {
	mensaje, result := sumar("Carlos", 4, 5, 7, 8)
	fmt.Println(mensaje, result)

	// Factorial
	var n int64
	fmt.Print("Introduce un numero para calcular su factorial:")
	fmt.Scanln(&n)
	fmt.Printf("Factorial: %v\n", factorial(n))

	// Funcion anónima
	func() {
		fmt.Println("Hola desde la función anónima")
	}()

	miFuncion := func(nombre string) string {
		return fmt.Sprintf("Hola %s, desde la función anónima", nombre)
	}

	fmt.Println(miFuncion("Carlos"))

	// Closures
	reapeat3 := repeat(10)
	fmt.Println(reapeat3("  C a r l o s  "))
	fmt.Println(reapeat3("LO KA"))
}
