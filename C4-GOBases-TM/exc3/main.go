// Golang program to illustrate the usage of
// fmt.Errorf() function

// Including the main package
package main

// Importing fmt
import (
	"fmt"
)

// Calling main
func main() {
	salary := 150000
	if salary < 150000 {
		err := fmt.Errorf("El salario %d  no alcanza el minimo imponible", salary)
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuestos")
	}

}
