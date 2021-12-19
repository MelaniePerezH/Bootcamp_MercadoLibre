package main

import (
	"errors"
	"fmt"
)

func main() {
	const (
		perro = "perro"
		gato  = "gato"
	)

	animalPerro, err := Animal(perro)
	animalGato, err := Animal(gato)

	var cantidad int
	var cantidades int
	cantidad += animalPerro(5)
	cantidades += animalGato(8)
	if err != nil {
		fmt.Println("No es posible acceder a tu solicitud")
	}
	fmt.Printf("Alimento necesario para los perros son %d Kg\n", cantidad)
	fmt.Printf("Alimento necesario para gatos son %d Kg\n", cantidades)
}
func perros(numPerros int) int {
	alimento := 10
	res := alimento * numPerros
	return res
}
func gatos(numGatos int) int {
	alimento := 5
	res := alimento * numGatos
	return res
}

func Animal(Animal string) (func(int) int, error) {
	switch Animal {
	case "perro":
		return perros, nil
	case "gato":
		return gatos, nil
	default:
		return nil, errors.New("Operación inválida")
	}
}
