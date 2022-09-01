package main

import (
	"fmt"
	"time"
)

/*
В данном файле реализован паттерн фасад на примере комбинированного пульта управления
устройствами (кондиционер и телевизор).
*/

// Телевизор
type tv struct {
}

// Функция включает телевизор
func (t *tv) turnOn() {
	fmt.Println("Tv is on")
}

// Функция выключает телевизор
func (t *tv) turnOff() {
	fmt.Println("Tv is off")
}

//Функция переключает канал телевизора на заданный
func (t *tv) selectChanel(chanel int) {
	fmt.Printf("Chanel %d is selected\n", chanel)
}

// Кондиционер
type conditioner struct {
}

// Функция включает кондиционер
func (c *conditioner) turnOn() {
	fmt.Println("Conditioner is on")
}

// Функция выключает кондиционер
func (c *conditioner) turnOff() {
	fmt.Println("Conditioner is off")
}

// Функция устанавливает температуру на заданную
func (c *conditioner) setTemperature(temp int) {
	fmt.Printf("Set temperature to %d celsius\n", temp)
}

// Структура, представляющая фасад, задающая 2 режима работы отдых и работа
type relaxFacade struct {
	c conditioner
	t tv
}

// Режим отдыха включает телевизор на 14 канал, включает кондиционер на 23 градуса
func (f *relaxFacade) relax() {
	f.t.turnOn()
	f.t.selectChanel(14)
	f.c.turnOn()
	f.c.setTemperature(23)
	fmt.Printf("\nJust relax!\n\n")
}

// Режим пора идти на работу выключает оба устройства
func (f *relaxFacade) goToWork() {
	f.t.turnOff()
	f.c.turnOff()
	fmt.Printf("\nGo to work!\n\n")
}

// Клиент фасада - функция main()
func main() {
	// Создаем структуру фасада
	facade := relaxFacade{t: tv{}, c: conditioner{}}
	// Включаем режим отдыха
	facade.relax()
	time.Sleep(3 * time.Second)
	// Включаем режим пора на работу
	facade.goToWork()
}
