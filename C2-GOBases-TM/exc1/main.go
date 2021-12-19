package main

import "fmt"

func main() {
	fmt.Println(impuesto(60000))

}

func impuesto(salario int) int {
	if salario > 50000 && salario < 150000 {
		return salario * 17 / 100
	} else {
		return salario * 27 / 100
	}

}
