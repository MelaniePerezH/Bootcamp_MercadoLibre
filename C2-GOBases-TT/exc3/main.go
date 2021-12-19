package main

import "fmt"

//const
const (
	pequeno = "pequeno"
	mediano = "mediano"
	grande  = "grande"
)

//Structs
type tienda struct {
	nombreProduct string
}
type producto struct {
	tipoProducto string
	nombre       string
	precio       float64
}

//Metodos
func (p producto) CalcularCosto() float64 {
	if p.tipoProducto == "pequeno" {
		return 0.0
	} else if p.tipoProducto == "mediano" {
		mantenimiento := p.precio * 0.03
		return mantenimiento
	} else if p.tipoProducto == "grande" {
		mantenimiento := p.precio * 0.06
		return mantenimiento + 2500
	}
	return 0.0
}
func (p producto) Total() float64 {
	if p.tipoProducto == "pequeno" {
		return p.precio
	} else if p.tipoProducto == "mediano" {
		mantenimiento := p.precio * 0.03
		return mantenimiento + p.precio
	} else if p.tipoProducto == "grande" {
		mantenimiento := p.precio * 0.06
		return mantenimiento + 2500 + p.precio
	}
	return 0.0
}

func (p producto) Agregar() tienda {
	new := tienda{p.nombre}
	return new
}

//Interfaces
type Producto interface {
	CalcularCosto() float64
}
type Ecommerce interface {
	Total() float64
}

//Funciones
func nuevoProducto(p producto) {
	//new := producto{tipoProducto, nombre, precio}
	fmt.Println("Los datos del producto ingresado son: \n", "Tipo de producto:", p.tipoProducto, "\n", "Nombre producto:", p.nombre, "\n", "Precio: ", "$", p.precio)
}

func nuevaTienda() {}

//Funciones para llamar las inter
func calcular(p Producto) {
	fmt.Println("El costo adicional a pagar por el producto es: $", p.CalcularCosto())
}
func calcularTotal(p Ecommerce) {
	fmt.Println("El costo TOTAL a pagar por el producto es: $", p.Total())
}

//Main
func main() {
	//Aqui se almacena en struct producto y se calcula el costo adicional con CalcularCosto()
	myproof := producto{"mediano", "chaqueta", 55400.0}
	calcular(myproof)
	fmt.Println("_____________________________________________________________")
	//Calcular costo total producto
	calcularTotal(myproof)
	fmt.Println("_____________________________________________________________")
	// Agregar productos a la tienda
	nuevoProducto(myproof)
	fmt.Println("_____________________________________________________________")
	fmt.Println("Nuevo producto agregado a la tienda: ", myproof.Agregar().nombreProduct)
	fmt.Println("_____________________________________________________________")

}
