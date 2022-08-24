package main

import "fmt"

type Figura interface {
	calcularArea() float64
	calcularPerimetro() float64
}

type Cuadrado struct {
	lado float64
}

type Circulo struct {
	radio float64
}

func (c *Cuadrado) calcularArea() float64 {
	return c.lado * c.lado
}

func (c *Circulo) calcularArea() float64 {
	return pi * (c.radio * c.radio)
}

func (c *Cuadrado) calcularPerimetro() float64 {
	return 4 * c.lado
}

func (c *Circulo) calcularPerimetro() float64 {
	return 2 * pi * c.radio
}

func calcularAreaFigura(figura Figura) float64 {
	return figura.calcularArea()
}

func calcularPerimetroFigura(figura Figura) float64 {
	return figura.calcularPerimetro()
}

const pi float64 = 3.141592

func main() {
	cuadrado := Cuadrado{lado: 5.46}
	circulo := Circulo{radio: 3.5}

	fmt.Println("Area del cuadrado es:", calcularAreaFigura(&cuadrado))
	fmt.Println("Area del circulo es:", calcularAreaFigura(&circulo))

	fmt.Println("Perimetro del cuadrado es:", calcularPerimetroFigura(&cuadrado))
	fmt.Println("Perimetro del circulo es:", calcularPerimetroFigura(&circulo))
}
