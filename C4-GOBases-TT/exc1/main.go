package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Iniciar archivo")
	_, err := os.Open("customers.txt")
	if err != nil {
		fmt.Println("El archivo indicado no fue encontrado o esta danado")
		panic(err)
	}
	fmt.Println("Fin de la ejecucion, espero que no se imprima ya que está el panic")
}
