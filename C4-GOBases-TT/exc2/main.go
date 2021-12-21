package main

import (
	"errors"
	"fmt"
	"os"
)

type cliente struct {
	legajo    []Generador
	nombre    string
	apellido  string
	cedula    int
	telefono  int
	domicilio string
}
type Generador struct {
	id int
}

func generarId() Generador {
	newid := Generador{}
	for i := 0; i < 100; i++ {
		newid.id = i
	}
	return newid
}

func crearCliente(nombre string, apellido string, cedula int, telefono int, domicilio string) (cliente, error) {
	defer func() {
		err := recover()
		opening, err := os.Open("customers.txt")

		if err != nil {
			fmt.Println("El archivo indicado no fue encontrado o esta danado")
			panic(err)
		}
		defer opening.Close()

	}()
	defer fmt.Println("No han quedado archvios abiertos")
	newcliente := cliente{}
	if nombre == "" || apellido == "" || cedula == 0 || telefono == 0 || domicilio == "" {
		fmt.Println(errors.New("No puede ingresar valores nulos"))
		defer fmt.Println("Se detectaron varios errores en tiempo de ejecucion")
	} else {
		newcliente.nombre = nombre
		newcliente.apellido = apellido
		newcliente.cedula = cedula
		newcliente.telefono = telefono
		newcliente.domicilio = domicilio
	}
	fmt.Println("Fin de la ejecucion")
	return newcliente, nil

}

func main() {
	generarId()
	crearCliente("", "Perez", 10096845, 3106894895, "Calle 13# 15 68")

}
