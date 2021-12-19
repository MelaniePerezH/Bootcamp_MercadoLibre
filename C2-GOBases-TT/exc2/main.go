package main

import "fmt"

func main() {
	//Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
	anadirElementos(4.5, 9.6, 10.3)

	//Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
	for _, value := range matrix {
		if value != 0 {
			fmt.Printf("-%g\n", value)

		}
	}
	//Dimensiones de Matrix
	//columnas

	fmt.Println("Las filas en Matrix son:", len(matrix))

	//valor maximo
	almacenar := 0.0
	for _, value := range matrix {

		if value > almacenar {
			almacenar = value
		}

	}
	fmt.Println("El valor máximo encontrado en Matrix es : ", almacenar)

}

//Estructura que represente una matriz de datos.

var matrix [100]float64

func anadirElementos(values ...float64) {
	for contador, i := range values {
		matrix[contador] = i
	}

}
