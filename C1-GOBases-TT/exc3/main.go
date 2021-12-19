package main

import "fmt"

func main() {
	cumplimiento := condicionesBanco(50, true, 5)
	fmt.Println(cumplimiento)
}
func condicionesBanco(edad int, empleado bool, antiguedad int) string {
	if edad > 22 && empleado == true && antiguedad > 1 {
		requisitos := "Usted cumple con los requisitos necesarios para acceder a un credito con el Banco"
		return requisitos
	} else {
		requisitos := "Usted  NO cumple con los requisitos para acceder a un credito con el Banco"
		return requisitos
	}
}
