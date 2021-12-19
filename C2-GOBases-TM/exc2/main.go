package main

import "fmt"

func main() {
	fmt.Println(promedio(2, 3, 5))
}

func promedio(calificaciones ...int) int {
	var suma int
	var respromedio int

	for _, calificacion := range calificaciones {
		suma += calificacion
		respromedio = suma / len(calificaciones)
	}
	return respromedio
}
