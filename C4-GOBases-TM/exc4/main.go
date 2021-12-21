package main

import (
	"errors"
	"fmt"
	"os"
)

func salarioMensual(horas int, valorHora int) (int, error) {
	salario := horas * valorHora
	if horas < 80 {
		fmt.Println(errors.New("No puede haber trabajado menos de 80 hrs mensuales"))
	} else if salario >= 150000 {
		descuento := salario * 10 / 100
		salario = salario - descuento
	}
	return salario, nil
}
func aguinaldo(salary int, mesesTrabajados int) (int, error) {
	if salary < 0 {
		err := fmt.Errorf("No puede ingresar valores nulos %d ", salary)
		fmt.Println(err.Error())
	}
	if mesesTrabajados < 0 {
		err := fmt.Errorf("No puede ingresar valores nulos %d ", mesesTrabajados)
		fmt.Println(err.Error())
	}
	calculo := salary / 12 * mesesTrabajados
	return calculo, nil
}
func main() {
	calculos, err := salarioMensual(50, 500)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("El salario mensual del trabajador es: ", "$", calculos)

	calculos2, err := aguinaldo(2000000, 5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("El valor correspondiente al aguinaldo es: ", "$", calculos2)
}
