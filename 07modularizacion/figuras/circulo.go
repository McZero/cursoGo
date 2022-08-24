package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (c *Circulo) calcularArea() float64 {
	return math.Pi * (c.Radio * c.Radio)
}

func (c *Circulo) calcularPerimetro() float64 {
	return 2 * math.Pi * c.Radio
}
