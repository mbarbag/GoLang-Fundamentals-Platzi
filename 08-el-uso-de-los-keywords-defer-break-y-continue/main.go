package main

import "fmt"

func main() {
	// Defer
	// Antes de que termine el main se ejecutará el defer
	// Por ende se imprimirá de última "hola"
	// Sirve por ejemplo para cerrar archivos
	defer fmt.Println("hola")
	fmt.Println("mundo")

	// continue y break
	// se usan dentro del for
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// continue
		// sirve por ejemplo cuando puede pasar un error
		// pero yo quiero que el for continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// break
		// cuando pasa una condición determinada
		// y quiero que el código deje de ejecutarse
		if i == 8 {
			fmt.Println("break")
			break
		}
	}
}
