package main

import (
	"fmt"
	"paquetes/figuras"
	"paquetes/models"

	"github.com/donvito/hellomod"
)

func main() {
	/*
		mensajes.Hola()
		mensajes.Imprimir()
	*/

	cua := figuras.Cuadrado{Lado: 5}
	cir := figuras.Circulo{Radio: 5.25}

	fmt.Println("Cuadrado:", figuras.CalcularMedidas(&cua))
	fmt.Println("Circulo:", figuras.CalcularMedidas(&cir))

	p1 := models.Persona{}
	p1.New("Carlos", 34)

	fmt.Println(p1.GetNombre())
	p1.SetNombre("Ernesto")
	fmt.Println(p1)

	hellomod.SayHello()
}
