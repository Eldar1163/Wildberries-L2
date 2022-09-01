package main

import "fmt"

/*
В данном файле реализован паттерн Builder на примере пекарни пицц двух видов: гавайской и спайси
*/

// Структура пиццы с тестом, соусом и топпингом
type pizza struct {
	dough   string
	sauce   string
	topping string
}

// Интерфейс для готовки пиццы
type pizzaBuilder interface {
	buildDough()
	buildSauce()
	buildTopping()
	getPizza() *pizza
	createNewPizzaProduct()
}

// Структура для готовки гавайской пиццы
type hawaiianPizzaBuilder struct {
	p pizza
}

func (pb *hawaiianPizzaBuilder) getPizza() *pizza {
	return &pb.p
}

func (pb *hawaiianPizzaBuilder) buildDough() {
	pb.p.dough = "cross"
}

func (pb *hawaiianPizzaBuilder) buildSauce() {
	pb.p.sauce = "mild"
}

func (pb *hawaiianPizzaBuilder) buildTopping() {
	pb.p.topping = "ham+pineapple"
}

func (pb *hawaiianPizzaBuilder) createNewPizzaProduct() {
	pb.p = pizza{}
}

// Структура для готовки спайси пиццы
type spicyPizzaBuilder struct {
	p pizza
}

func (pb *spicyPizzaBuilder) getPizza() *pizza {
	return &pb.p
}

func (pb *spicyPizzaBuilder) buildDough() {
	pb.p.dough = "pan baked"
}

func (pb *spicyPizzaBuilder) buildSauce() {
	pb.p.sauce = "hot"
}

func (pb *spicyPizzaBuilder) buildTopping() {
	pb.p.topping = "pepperoni+salami"
}

func (pb *spicyPizzaBuilder) createNewPizzaProduct() {
	pb.p = pizza{}
}

// Оффициант (директор в понятии данного паттерна. Управляет процессом готовки и подачи пиццы)
type waiter struct {
	pizzaBuilder pizzaBuilder
}

func (w *waiter) getPizza() *pizza {
	return w.pizzaBuilder.getPizza()
}

func (w *waiter) constructPizza() {
	w.pizzaBuilder.createNewPizzaProduct()
	w.pizzaBuilder.buildDough()
	w.pizzaBuilder.buildSauce()
	w.pizzaBuilder.buildTopping()
}

func main() {
	// Создаем оффицианта и говорим, что будем гавайскую пиццу
	waiter := waiter{pizzaBuilder: &hawaiianPizzaBuilder{}}
	// Готовим пиццу
	waiter.constructPizza()
	// Получаем пиццу
	hawaiianPizza := *waiter.getPizza()
	// Делаем фото пиццы
	fmt.Println("HawaiianPizza: ", hawaiianPizza)

	// Говорим, что будем спайси пиццу
	waiter.pizzaBuilder = &spicyPizzaBuilder{}
	// Готовим пиццу
	waiter.constructPizza()
	// Получаем пиццу
	spicyPizza := *waiter.getPizza()
	// Делаем фото пиццы
	fmt.Println("SpicyPizza: ", spicyPizza)
}
