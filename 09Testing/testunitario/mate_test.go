package testunitario

import "testing"

func TestSuma(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)

		if total != item.c {
			t.Errorf("Suma incorrecta, el resultado es %d y se esperaba %d", total, item.c)
		}

	}
}

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 2},
		{4, 2, 4},
		{25, 24, 25},
	}

	for _, item := range tabla {
		total := GetMax(item.a, item.b)

		if total != item.c {
			t.Errorf("Máximo incorrecto, el resultado es %d y se esperaba %d", total, item.c)
		}

	}
}

func TestFibo(t *testing.T) {
	tabla := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tabla {
		total := Fibonacci(item.n)

		if total != item.r {
			t.Errorf("Fibonacci incorrecto, el r8esultado es %d y se esperaba %d", total, item.r)
		}
	}
}
