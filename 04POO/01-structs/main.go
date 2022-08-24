package main

import "fmt"

// Struct persona
type Persona struct {
	nombre string
	edad   int
}

// metodos de struct
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}

func (p *Persona) editarNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

// Herencia
type Empleado struct {
	Persona
	clave string
}

func main() {
	persona1 := Persona{"Carlos", 34}
	//fmt.Println(persona1)
	persona1.imprimir()

	persona1.editarNombre("Ernesto")
	//fmt.Println("Modificado:", persona1)
	persona1.imprimir()

	persona2 := Persona{
		nombre: "Juan",
		edad:   35,
	}

	//fmt.Println(persona2)
	persona2.imprimir()

	persona2.editarNombre("Juan Perez Jolote")
	persona2.imprimir()

	empleado := Empleado{}
	empleado.nombre = "Ejemplo"
	empleado.edad = 54
	empleado.clave = "0001"

	empleado.imprimir()
	fmt.Println(empleado)
}
