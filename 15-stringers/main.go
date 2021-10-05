package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB ram, %d GB disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	// personalizar el output de structs
	myPC := pc{ram: 16, disk: 100, brand: "msi"}
	fmt.Println(myPC)
}
