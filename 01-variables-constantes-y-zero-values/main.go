package main

import "fmt"

func main() {
	//Declaracion de constantes (nunca cambian en el tiempo)
	const pi float64 = 3.1416
	const pi2 = 3.14

	fmt.Println(pi, pi2)

	//Declaracion de variables enteras

	//Para crearla por primera vez se pueden tres formas:
	//1) Poner los dos puntos, de esta manera Go asigna de manera autmática el tipo de la vble.
	base := 12
	//2) Declarando el "var", el tipo y asignandole el valor.
	var altura int = 14
	//3) Declarando el "var" y el tipo.
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna valores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Ejercicio: calcular el área de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)
}
