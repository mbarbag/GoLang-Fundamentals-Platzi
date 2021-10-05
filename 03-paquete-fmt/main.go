package main

import "fmt"

func main() {
	// Declaración de variables
	helloMessage := "Hello"
	worldMessage := "world"

	//Println -> ln es para imprimir con salto de linea
	println(helloMessage, worldMessage)

	//Printf -> para reempazar vbles. %s para strings %d para ints %v si no sabemos
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)

	//Sprintf -> genera un string pero no imprime en consola
	message := fmt.Sprintf("%v tiene más de %v cursos", nombre, cursos)
	println(message)

	//Tipo dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
