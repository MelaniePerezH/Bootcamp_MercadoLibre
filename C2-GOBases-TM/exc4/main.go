package main

import (
	"errors"
	"fmt"
)

func main() {
	const (
		minimo   = "minimo"
		promedio = "promedio"
		maximo   = "maximo"
	)

	minFunc, err := operacion(minimo)
	promFunc, err := operacion(promedio)
	maxFunc, err := operacion(maximo)

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println("El valor minimo es ", valorMinimo)
	fmt.Println("El valor maximo es ", valorMaximo)
	fmt.Println("El valor promedio es", valorPromedio)

	if err != nil {
		fmt.Println("No es posible acceder a tu solicitud")
	}
}
func minFunc(num ...int) int {
	minimo := 1000
	for _, i := range num {
		if i < minimo {
			minimo = i
		}
	}
	return minimo
}

func maxFunc(num ...int) int {
	maximo := 0
	for _, i := range num {
		if i > maximo {
			maximo = i
		}
	}
	return maximo
}
func promFunc(num ...int) int {
	almacenar := 0
	for _, i := range num {
		almacenar = almacenar + i
	}
	res := almacenar / len(num)
	return res

}
func operacion(operacion string) (func(...int) int, error) {
	switch operacion {
	case "minimo":
		return minFunc, nil
	case "promedio":
		return promFunc, nil
	case "maximo":
		return maxFunc, nil
	default:
		return nil, errors.New("Operación inválida")
	}
}
