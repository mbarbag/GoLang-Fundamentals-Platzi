package main

import "fmt"

func main() {

	x := 10
	y := 50

	//suma
	result := x + y
	fmt.Println("suma:", result)

	//resta
	//podemos poner el = sin : porque la vble result ya existe
	result = y - x
	fmt.Println("resta:", result)

	//multiplicación
	result = x * y
	fmt.Println("multiplicación:", result)

	//división
	result = y / x
	println("división:", result)

	//módulo
	result = y % x
	println("modulo:", result)

	//Incremental
	x++
	println("decremental:", x)

	//Decremental
	x--
	println("decremental:", x)
}
