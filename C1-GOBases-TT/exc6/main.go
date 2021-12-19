package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	//imprimir edad de pablo
	counter := 0
	fmt.Println(employees["Benjamin"])
	for _, employee := range employees {

		if employee > 21 {
			counter++
		}
	}
	fmt.Println(counter)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}
