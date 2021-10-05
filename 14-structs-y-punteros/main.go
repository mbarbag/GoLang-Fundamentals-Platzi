package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

//una funcion que accedera a sus valores mediante el puntero
func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {
	a := 50
	//direccion de memoria de a
	b := &a

	fmt.Println(b)
	//el * sirve para acceder al valor de esa direcciónd de memoria
	fmt.Println(*b)

	//Si modico el valor al que apunta una direccion de memoria
	//las variables que apuntan a esa dir de memoria van a cambiar de valor
	//En otras palabra a = *b
	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)
	myPC.ping()

	//Si quiero actualizar la ram una forma de hacerlo es:
	//myPC.ram = 20
	//pero lo mejor es usar otra funcion accediendo a la memoria
	myPC.duplicateRAM()
	fmt.Println(myPC)
}
