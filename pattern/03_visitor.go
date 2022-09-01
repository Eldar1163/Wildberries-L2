package main

import "fmt"

/*
В данном файле реализован паттерн Visitor на примере банка, где есть 2 типа клиентов: физ. лица и компании
и для обоих типов клиентов следует определить операцию их сериализации в html
*/

// Интерфейс Visitor. Структуры, которые его реалзиуют, определяют методы visit* для каждого типа клиентов
type visitor interface {
	visitPersonAcc(acc *person)
	visitCompanyAcc(acc *company)
}

// Интерфейс account, реализует метод accept принимающий visitor
type account interface {
	accept(vis visitor)
}

// Физическое лицо
type person struct {
	fio       string
	accNumber string
}

func (p *person) accept(vis visitor) {
	vis.visitPersonAcc(p)
}

// Компания
type company struct {
	name      string
	regNumber string
	accNumber string
}

func (c *company) accept(vis visitor) {
	vis.visitCompanyAcc(c)
}

// Пустая структура реализующая методы интерфейса visitor
type htmlVisitor struct {
}

func (h htmlVisitor) visitPersonAcc(acc *person) {
	res := "<table><tr><td>Свойство<td><td>Значение</td></tr>\n"
	res += "<tr><td>Name<td><td>" + acc.fio + "</td></tr>\n"
	res += "<tr><td>Acc Number<td><td>" + acc.accNumber + "</td></tr>\n"
	fmt.Println(res)
	fmt.Println()
}

func (h htmlVisitor) visitCompanyAcc(acc *company) {
	res := "<table><tr><td>Свойство<td><td>Значение</td></tr>\n"
	res += "<tr><td>Name<td><td>" + acc.name + "</td></tr>\n"
	res += "<tr><td>Reg number<td><td>" + acc.regNumber + "</td></tr>\n"
	res += "<tr><td>Number<td><td>" + acc.accNumber + "</td></tr>\n"
	fmt.Println(res)
	fmt.Println()
}

func main() {
	// Создаем физ лицо
	person1 := &person{
		fio:       "Петров Иван Петрович",
		accNumber: "123456789",
	}

	// Создаем компанию
	company1 := &company{
		name:      "ООО Мечта",
		regNumber: "098123",
		accNumber: "999123456",
	}

	// Добавляем их в список
	accList := []account{
		person1, company1,
	}

	// Выполняем "визит", в результате которого выводятся сериализованные учетные записи
	for _, acc := range accList {
		acc.accept(htmlVisitor{})
	}
}
