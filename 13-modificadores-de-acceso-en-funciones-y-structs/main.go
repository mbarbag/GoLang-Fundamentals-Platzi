package main

import (
	"ejemplo/13-modificadores-de-acceso-en-funciones-y-structs/mypackage"
	"fmt"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)
}
