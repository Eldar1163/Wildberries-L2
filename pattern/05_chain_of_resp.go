package main

import "fmt"

/*
В данном файле реализован паттерн Chain of Responsibility. Проиллюстрирован на примере системы оповещения
менеджера проекта, о различных проблемах на нем. У оповещения есть 3 типа приоритета. Routine - обычное сообщение,
Important - важное сообщение и Asap (as soon as possible) - критически важное сообщение
*/

// Приоритеты сообщений
const (
	ROUTINE   int = 1
	IMPORTANT int = 2
	ASAP      int = 3
)

// Интерфейс с общим для всех notifier методом write
type writeInterface interface {
	write(message string)
}

// Структура оповестителя
type notifier struct {
	// Приоритет сообщения
	priority int
	// Ссылка на обертку, чтобы иметь возможность выполнять нужную реализацию метода write
	containerRef writeInterface
	// Ссылка на следующий оповеститель
	nextNotifier *notifier
}

// Метод сеттер устанавлиивает принимаемое значние полю nextNotifier
func (n *notifier) setNextNotifier(notifier *notifier) {
	n.nextNotifier = notifier
}

// Выполняет оповещения, при необходимости передает сообщение следующему оповестителю
func (n *notifier) notifyManager(message string, lvl int) {
	if lvl >= n.priority {
		n.write(message)
	}
	if n.nextNotifier != nil {
		n.nextNotifier.notifyManager(message, lvl)
	}
}

// Обобщенный метод write
func (n *notifier) write(message string) {
	n.containerRef.write(message)
}

// Простой оповеститель
type simpleReportNotifier struct {
	notifier
}

func (s simpleReportNotifier) write(message string) {
	fmt.Println("Notifying using simple report:", message)
}

// Email оповеститель
type emailNotifier struct {
	notifier
}

func (e emailNotifier) write(message string) {
	fmt.Println("Notifying using email:", message)
}

// SMS оповеститель
type SMSNotifier struct {
	notifier
}

func (s SMSNotifier) write(message string) {
	fmt.Println("Notifying using sms message:", message)
}

// Конструктор notifier
func constructNotifier(priority int, container writeInterface) *notifier {
	return &notifier{priority: priority, containerRef: container}
}

func main() {

	//Создание простого оповестиителя
	reportNotifier := simpleReportNotifier{}
	reportNotifier.notifier = *constructNotifier(ROUTINE, reportNotifier)

	//Создание email оповестиителя
	emailNotifier := emailNotifier{}
	emailNotifier.notifier = *constructNotifier(IMPORTANT, emailNotifier)

	//Создание sms оповестиителя
	smsNotifier := SMSNotifier{}
	smsNotifier.notifier = *constructNotifier(ASAP, smsNotifier)

	//Установка цепочки оповестителей
	reportNotifier.setNextNotifier(&emailNotifier.notifier)
	emailNotifier.setNextNotifier(&smsNotifier.notifier)

	// Передача оповещений различного приоритета
	reportNotifier.notifyManager("Everything is OK", ROUTINE)
	reportNotifier.notifyManager("Something goes wrong", IMPORTANT)
	reportNotifier.notifyManager("FATAL ERROR", ASAP)
}
