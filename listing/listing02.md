Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

В первом случае выводится 2, так как x именован на уровне функции
и defer изменяет именно его значение после возврата из функции.
Во втором случае выводится 1, так как здесь defer  уже не может повлить
на возвращенное значение.

Оператор defer не будет выполняться до тех пор пока окружающая его функция не
завершит свое выполнение. Операторы defer вызываются по принципу LIFO.

```