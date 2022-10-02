Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

В первом случае выводится nil, так как возвращаемый интерфейсный тип
содержит os.PathError == nil. Интерфейсный тип
содержит 2 поля (itab и data) и для того, чтобы он был равен nil,
необходимо, чтобы itab == nil и data == nil. Во втором случае выводится
false, тк itab != nil.

```