package main

import "time"

var almacenar int

type Productos struct {
	nombre   string
	precio   int
	cantidad int
}
type Service struct {
	nombre            string
	precio            int
	minutosTrabajados int
}
type Mantenimiento struct {
	nombre string
	precio int
}

func sumarProductos(nombre string, precio int, cantidad int) Productos {
	producto := Productos{}
	producto.nombre = nombre
	producto.precio = precio
	producto.cantidad = cantidad
	almacenar = almacenar + precio

	return producto
}
func sumarServicios(nombre string, precio int, minutosTrabajados int) Service {
	servicios := Service{}
	servicios.nombre = nombre
	servicios.precio = precio
	servicios.minutosTrabajados = minutosTrabajados
	almacenar += servicios.precio
	return servicios
}
func sumarMantenimiento(nombre string, precio int) Mantenimiento {
	mantenimiento := Mantenimiento{}
	mantenimiento.nombre = nombre
	mantenimiento.precio = precio
	almacenar += mantenimiento.precio
	return mantenimiento
}

func main() {

	go sumarMantenimiento("Zapatos", 15000)
	go sumarServicios("Lavanderia", 10000, 60)
	go sumarProductos("Zapatos", 55000, 5)
	time.Sleep(500 * time.Microsecond)
	println(almacenar)
}
