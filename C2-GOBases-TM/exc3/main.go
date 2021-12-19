package main

import "fmt"

func main() {
	fmt.Println(sueldo(1000, "A"))
}
func sueldo(minutos int, categoria string) int {
	horas := minutos / 60
	var salario int
	var salariomes int
	var porcentaje int

	if categoria == "C" {
		salario = horas * 1000
		return salario
	} else if categoria == "B" {
		salariomes = horas * 1500
		porcentaje = salariomes * 25 / 100
		salario = porcentaje + salariomes
		return salario
	} else if categoria == "A" {
		salariomes = horas * 3000
		porcentaje = salariomes * 50 / 100
		salario = salariomes + porcentaje
		return salario
	}
	return 0
}
