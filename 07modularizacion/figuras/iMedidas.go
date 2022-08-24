package figuras

import "fmt"

type Figura interface {
	calcularArea() float64
	calcularPerimetro() float64
}

func CalcularMedidas(figura Figura) string {
	return fmt.Sprintf("Perímetro: %f\nArea: %f\n", figura.calcularPerimetro(), figura.calcularArea())

}
