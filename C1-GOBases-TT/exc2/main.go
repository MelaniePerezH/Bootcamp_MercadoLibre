package main

import "fmt"

func main() {

	invocar := preciofinal(100000, 50)
	fmt.Println("Precio del producto con descuento $", invocar)

}
func preciofinal(precioProducto, descuento int) int {

	descuentoProducto := precioProducto * descuento / 100
	final := precioProducto - descuentoProducto
	return final
}
