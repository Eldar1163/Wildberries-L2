package main

import "fmt"

/*
В данном файле реализован паттерн Factory method на примере фабрики разработчиков, где существует 2 вида разработчиков
Golang и С++
*/

const (
	CPP    string = "cpp"
	GOLANG string = "golang"
)

// Фабрика разработчиков с методов создания разработчика
type developerFactory interface {
	createDeveloper() developer
}

// Вот фабричный метод
func createDeveloperFactoryBySpeciality(spec string) developerFactory {
	switch spec {
	case
		CPP:
		return &cppDeveloperFactory{}
	case GOLANG:
		return &golangDeveloperFactory{}
	default:
		panic(spec + " is unknown")
	}
}

// Конкретная фабрика c++ разработчиков
type cppDeveloperFactory struct {
}

func (cppDF *cppDeveloperFactory) createDeveloper() developer {
	return &cppDeveloper{}
}

// Конкретная фабрика golang разработчиков
type golangDeveloperFactory struct {
}

func (gDF *golangDeveloperFactory) createDeveloper() developer {
	return &golangDeveloper{}
}

// Обобщенный интерфейс разработчика с методом написания кода
type developer interface {
	writeCode()
}

// Конкретный c++ разработчик
type cppDeveloper struct {
}

func (cppD *cppDeveloper) writeCode() {
	fmt.Println("Cpp developer writes c++ code")
}

// Конкретный golang разработчик
type golangDeveloper struct {
}

func (gD *golangDeveloper) writeCode() {
	fmt.Println("Golang developer writes golang code")
}

// Клиент
func main() {
	// Создаем фабрику нужных разработчиков (использование фабричного метода)
	devFactory := createDeveloperFactoryBySpeciality(GOLANG)
	// Создаем разработчика
	dev := devFactory.createDeveloper()
	// Заставляем разработчика писать код
	dev.writeCode()
}
