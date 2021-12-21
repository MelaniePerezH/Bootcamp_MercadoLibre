package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 300000

	if salary < 150000 {
		fmt.Println(errors.New("El salario no alcanza el minimo imponible"))
		return
	}
	fmt.Println("Debe pagar impuestos")
}
