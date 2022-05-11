package main

import "fmt"

type Appliance interface {
	TurnOn()
}

type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}
func (f Fan) Oscillate() {
	fmt.Println("Rotating on base")
}

type CoffeePot string

func (c CoffeePot) TurnOn() {
	fmt.Println("Powering up")
}
func (c CoffeePot) Brew() {
	fmt.Println("Heating Up")
}

func Use(appliance Appliance) {
	fmt.Println(appliance)
	appliance.TurnOn()
	fan, ok := appliance.(Fan)
	if ok {
		fan.Oscillate()
	}
	coffeePot, ok := appliance.(CoffeePot)
	if ok {
		coffeePot.Brew()
	}
}

func main() {
	Use(Fan("Windco Breeze"))
	Use(CoffeePot("LuxBrew"))
}