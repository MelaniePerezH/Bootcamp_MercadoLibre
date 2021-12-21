package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var p, j int
	var aux int

	variable1 := rand.Perm(100)

	//ORDENAMIENTO POR INSERCION
	for p = 1; p < len(variable1); p++ {
		aux = variable1[p]
		j = p - 1
		for j >= 0 && aux < variable1[j] {
			variable1[j+1] = variable1[j]
			j--
		}
		variable1[j+1] = aux
	}
	fmt.Println(variable1)
	//ORDENAMIENTO POR SELECCION
	var h, k, menor, pos, tmp int
	for h = 0; h < len(variable1); h++ {
		menor = variable1[h]
		pos = h
		for k = h + 1; k < len(variable1); k++ {
			if variable1[k] < menor {
				menor = variable1[k]
				pos = k
			}
		}
		if pos != 0 {
			tmp = variable1[h]
			variable1[h] = variable1[pos]
			variable1[pos] = tmp
		}
	}
	fmt.Println(variable1)

	//variable2 := rand.Perm(1000)
	//variable3 := rand.Perm(10000)
}
