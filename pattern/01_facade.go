package main

import (
	"fmt"
	"time"
)

type tv struct {
}

func (t *tv) turnOn() {
	fmt.Println("Tv is on")
}

func (t *tv) turnOff() {
	fmt.Println("Tv is off")
}

func (t *tv) selectChanel(chanel int) {
	fmt.Printf("Chanel %d is selected\n", chanel)
}

type conditioner struct {
}

func (c *conditioner) turnOn() {
	fmt.Println("Conditioner is on")
}

func (c *conditioner) turnOff() {
	fmt.Println("Conditioner is off")
}

func (c *conditioner) setTemperature(temp int) {
	fmt.Printf("Set temperature to %d celsius\n", temp)
}

type relaxFacade struct {
	c conditioner
	t tv
}

func (f *relaxFacade) relax() {
	f.t.turnOn()
	f.t.selectChanel(14)
	f.c.turnOn()
	f.c.setTemperature(23)
	fmt.Println("\nJust relax!\n")
}

func (f *relaxFacade) goToWork() {
	f.t.turnOff()
	f.c.turnOff()
	fmt.Println("\nGo to work!\n")
}

func main() {
	facade := relaxFacade{t: tv{}, c: conditioner{}}
	facade.relax()
	time.Sleep(3 * time.Second)
	facade.goToWork()
}
