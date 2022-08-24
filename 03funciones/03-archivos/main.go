package main

import (
	"fmt"
	"os"
)

func main() {
	// funcion
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ups! al parecer no finalizao de forma correcta -", err)
		}
	}()

	if file, err := os.Open("hola.txt"); err != nil {
		panic("No es posible abrir el archivo")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo...")
			file.Close()
		}()

		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}
}
