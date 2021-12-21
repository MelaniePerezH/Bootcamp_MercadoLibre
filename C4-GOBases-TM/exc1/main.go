package main

import (
	"fmt"
	"os"
)

type myCustomError struct {
	salary int
	msg    string
}

//hacemos que nuestro tipo struct implemente el método Error()
func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.salary, e.msg)
}

func myCustomErrorTest(salary int) (int, error) {
	if salary < 150000 {
		return 150000, &myCustomError{
			salary: salary,
			msg:    "No alcanza el minimo imponible",
		}
	}
	return 150000, nil
}
func main() {
	salary, err := myCustomErrorTest(100000) //llamamos a nuestra func
	if err != nil {                          //hacemos una validación del valor de err
		fmt.Println(err) //si err no es nil, imprimimos el error y...
		os.Exit(1)       //utilizamos este método para salir del programa
	}
	fmt.Printf("Debe pagar impuesto ya que el salario es suepr a :  %d \n ", salary)
}
