package main

import "fmt"

type alumnos struct {
	nombre   string
	apellido string
	cedula   string
	fecha    string
}

func (a alumnos) detalle() {
	fmt.Println("Bienvenid@ el estudiante consultado es " + " " + a.nombre + " " + a.apellido + " " + a.cedula + a.fecha)
}

func main() {
	alumno := alumnos{nombre: "pepito", apellido: "perez", cedula: "125456", fecha: "20/12/2021"}
	fmt.Println(alumno)

	alumno.detalle()
}
