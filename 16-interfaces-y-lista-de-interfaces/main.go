package main

import (
	"fmt"
)

// Una interfaz es un método en el que se pueden compartir diferentes otros métodos
// Ejemplo: un método que aplica a varios structs
// Las interfaces se usan cuando se tienen muchas estructuras que comparten una misma función

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.altura * r.base
}

func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	calcular(myCuadrado)
	calcular(myRectangulo)

	// Lista interfaces
	// En go uno no puede hacer una lista con distintos tipos de datos.
	// Ejm: no puede existir un slide ["Hola", 12, 4.90], debe ser de un mismo tipo de dato.
	// Eso se soluciona creando una lista de interfaces.
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)
}
