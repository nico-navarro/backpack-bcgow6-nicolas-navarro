package main

import "fmt"

const (
	PEQ = "Pequeño"
	MED = "Mediano"
	GRA = "Grande"
)

// Producto
type Producto interface {
	CalcularCosto() float64
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

func nuevoProducto(tipo string, nombre string, precio float64) (p producto) {
	p.tipo = tipo
	p.nombre = nombre
	p.precio = precio
	return
}

func (p producto) CalcularCosto() (costoTotal float64) {
	switch p.tipo {
	case PEQ:
		costoTotal = p.precio
	case MED:
		costoTotal = p.precio * 1.03
	case GRA:
		costoTotal = (p.precio * 1.06) + 2500
	}
	return
}

// Tienda
type Ecommerce interface {
	Total() float64
	Agregar(p producto)
}

type tienda struct {
	productos []producto
}

func nuevaTienda() (e Ecommerce) {
	return
}

func (t tienda) Total() (precioTotal float64) {
	for _, producto := range t.productos {
		precioTotal += producto.CalcularCosto()
	}
	return
}

func (t *tienda) Agregar(p producto) {
	t.productos = append(t.productos, p) //es una copia del producto, que pasa si quiero pasarle el producto como referencia? por si cambia el precio por ej
	return
}

func main() {
	zapatilla := nuevoProducto(PEQ, "Zapatilla Nike", 25_000)
	cama := nuevoProducto(GRA, "Cama 1 plaza", 150_000)
	// tienda := nuevaTienda() // no me funciona usando esto
	tienda := tienda{}
	tienda.Agregar(zapatilla)
	tienda.Agregar(cama)
	total := tienda.Total()
	fmt.Println(total)
}
