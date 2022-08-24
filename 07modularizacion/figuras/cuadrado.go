package figuras

type Cuadrado struct {
	Lado float64
}

func (c *Cuadrado) calcularArea() float64 {
	return c.Lado * c.Lado
}

func (c *Cuadrado) calcularPerimetro() float64 {
	return 4 * c.Lado
}
