Что выведет программа? Объяснить вывод программы.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
Программа выведет [77 78 79]
Это произойдет потому что операция создания слайса [a:b]
берет элемента с индексами от a, до b-1 включая их оба.

```