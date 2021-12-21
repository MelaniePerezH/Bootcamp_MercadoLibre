package main

import "fmt"

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Productos
}

type Productos struct {
	Nombre   string
	Precio   int
	Cantidad int
}

func nuevoProducto(nombre string, precio int, cantidad int) Productos {
	producto := Productos{}
	producto.Nombre = nombre
	producto.Precio = precio
	producto.Cantidad = cantidad

	return producto

}

func agregarProducto(user *Usuario, producto Productos, cantidad int) {
	producto.Cantidad = cantidad
	user.Productos = append(user.Productos, producto)
}

func borrarProductos(user *Usuario) {
	user.Productos = nil
	fmt.Printf("Se borraron los productos de %s\n\n", user.Nombre)
}

func main() {
	user := new(Usuario)
	user.Nombre = "Meli"
	user.Apellido = "Perez"
	user.Correo = "test@test.org"

	user2 := new(Usuario)
	user2.Nombre = "PEPITO"
	user2.Apellido = "Roncancio"
	user2.Correo = "don roncancio@gmail.com"

	corbata := nuevoProducto("Corbata", 15000, 2)
	zapatos := nuevoProducto("Zapatos", 400.00, 5)

	agregarProducto(user, corbata, 5)
	agregarProducto(user, zapatos, 2)
	borrarProductos(user)

}
