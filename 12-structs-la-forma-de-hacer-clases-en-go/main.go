package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//otra manera
	var otherCar car
	otherCar.brand = "ferrari"
	fmt.Println(otherCar)
}
