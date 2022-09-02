package main

import "fmt"

/*
В данном файле реализова шаблон State на примере пульта дистанционного управления телевизором
*/

// Состояние обощенное
type state interface {
	doAction()
}

// Конкретное состояние включенного телевизора
type tvOn struct {
}

func (t *tvOn) doAction() {
	fmt.Println("TV is turned ON")
}

// Конкретное состояние выключенного телевизора
type tvOff struct {
}

func (t *tvOff) doAction() {
	fmt.Println("TV is turned OFF")
}

// Структура контекста с инкапсулированным состоянием
type tvContext struct {
	state state
}

func (tvc *tvContext) doAction() {
	tvc.state.doAction()
}

func main() {
	// Создаем контекст
	context := tvContext{}

	// Меняем состояние и меняется действие метода doAction()

	context.state = &tvOn{}
	context.doAction()

	context.state = &tvOff{}
	context.doAction()
}
