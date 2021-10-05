package main

import (
	"fmt"
)

//Va a agarrar un mensaje y lo va a guardar en un canal
func mensaje(text string, c chan string) {
	c <- text
}

func main() {
	//Un channel que manejará dos datos a la vez
	c := make(chan string, 2)

	//Le insertamos dos mensajes
	c <- "Mensaje1"
	c <- "Mensaje2"

	//Si intento meter otro dato no me va a dejar porque ya lo definí solo para 2
	//c <- "Mensaje3"

	// El len me dice cuantas goroutines hay dentro del channel.
	// El cap me indica cuánta es la capacidad máxima que puede almacenar ese channel.
	fmt.Println(len(c), cap(c))

	// Range y close

	// Close lo que hace es decirle a la goroutine de go que va a cerrar el canal, es decir, el canal no va a recibir ningún otro dato adicional.
	// Lo ideal es cerrar los canales una vez que ya dejas de usarlo.
	close(c)
	//Una vez cerrado no puedo insertar valores
	// c <- "Mensaje3"

	//Range: Recorrer todos los mensajes que están en el channel
	for mensaje := range c {
		fmt.Println(mensaje)
	}

	// Cuando se empiezan a manejar multiples canales y no tenemos certeza de cuál de los canales va a responder primero, usamos:
	// Select

	email1 := make(chan string)
	email2 := make(chan string)
	go mensaje("mensaje1", email1)
	go mensaje("mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		// en caso de que el mensaje venga del email1
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
