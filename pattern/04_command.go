package main

import "fmt"

/*
В данном файле реализован паттерн команда, он инкапсулирует действие
в объект, который затем выполняет приемник (control). Паттерн
представлен на примере светильника, которые можно включить и выключить
*/

// Интерфейс команды с методом execute (можно добавить метод undo для отмены команды)
type command interface {
	execute()
}

// Светильник
type light struct {
}

// Метод включения светильника
func (l *light) turnOn() {
	fmt.Println("Light is on")
}

// Метод выключения светильника
func (l *light) turnOff() {
	fmt.Println("Light is off")
}

// Структура, представляющая команду включения светильника
type turnOnLightCommand struct {
	light light
}

// Реализация интерфейса command
func (onCmd *turnOnLightCommand) execute() {
	onCmd.light.turnOn()
}

// Структура, представляющая команду выключения светильника
type turnOffLightCommand struct {
	light light
}

// Реализация интерфейса command
func (offCmd *turnOffLightCommand) execute() {
	offCmd.light.turnOff()
}

// Приемник команд (пульт с двумя кнопками)
type control struct {
	flipUp   command
	flipDown command
}

// Кнопка в верхнем положении
func (c *control) flipUpExec() {
	c.flipUp.execute()
}

// Кнопка в нижнем положении
func (c *control) flipDownExec() {
	c.flipDown.execute()
}

func main() {
	// Создаем светильник
	light := light{}

	// Создаем команды
	onLight := turnOnLightCommand{light: light}
	offLight := turnOffLightCommand{light: light}

	// Создаем пульт
	ctrl := control{
		flipUp:   &onLight,
		flipDown: &offLight,
	}

	// Щечкаем кнопку
	ctrl.flipUpExec()
	ctrl.flipDownExec()

	fmt.Println()

	// Меняем местами команды
	ctrl = control{
		flipUp:   &offLight,
		flipDown: &onLight,
	}

	// Еще раз щелкаем кнопку
	ctrl.flipUpExec()
	ctrl.flipDownExec()
}
