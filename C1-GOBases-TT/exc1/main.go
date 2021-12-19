package main

import "fmt"

func main() {
	var palabra string
	palabra = "Melanie"
	fmt.Println(len(palabra))
	for _, i := range palabra {
		fmt.Printf("%c\n", i)

	}
}
