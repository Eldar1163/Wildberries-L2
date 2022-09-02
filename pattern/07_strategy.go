package main

import "fmt"

/*
В данном файле реализован паттерн Strategy на примере человека и его активностей в течении дня
*/

// Человек и его активность
type human struct {
	activity activity
}

// Метод выполняет активность
func (h human) execute() {
	h.activity.justDoIt()
}

// Обобщенный интерфейс активности
type activity interface {
	justDoIt()
}

// Конкретная активность (чтение)
type reading struct {
}

func (r *reading) justDoIt() {
	fmt.Println("Reading...")
}

// Конкретная активность (сон)
type sleeping struct {
}

func (s *sleeping) justDoIt() {
	fmt.Println("Sleeping...")
}

// Конкретная активность (прием пищи)
type eating struct {
}

func (e *eating) justDoIt() {
	fmt.Println("Eating...")
}

// Конкретная активность (занятия спортом)
type training struct {
}

func (t *training) justDoIt() {
	fmt.Println("Training...")
}

func main() {
	// Создание человека
	human := human{}

	// Установка и выполнение различных активностей

	human.activity = &sleeping{}
	human.execute()

	human.activity = &eating{}
	human.execute()

	human.activity = &training{}
	human.execute()

	human.activity = &reading{}
	human.execute()
}
