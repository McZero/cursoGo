package mensajes

import "fmt"

func Hola() {
	fmt.Println("Hola desde el paquete mensajes")
}

const mensaje = "Hola desde mi constante"

func Imprimir() {
	fmt.Println(mensaje)
}
