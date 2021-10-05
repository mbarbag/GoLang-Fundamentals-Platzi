package main

import (
	"fmt"
)

func say(text string, c chan string) {
	// Tomaremos el texto que se recibe en la función
	// y se lo pasaremos al canal
	c <- text
}

// Channels
//Una manera de trabajar con las Goroutines un poco menos eficiente
//pero permite trabajar con datos a través de ella.
func main() {
	// Creamos un channel
	// El tipo de dato que pasará por ese canal será string.
	// Recibirá 1 (una) goroutine a la vez de tipo string.
	c := make(chan string, 1)

	fmt.Println("Hello")

	//Invocamos la funcion say con una goroutine
	go say("Bye", c)

	//Con esto aseguramos que la goroutine del main espere la ejecucioin de la goroutine de la función say antes de terminar.
	// Extraemos el dato que agregamos al canal anteriormente
	fmt.Println(<-c)
}
