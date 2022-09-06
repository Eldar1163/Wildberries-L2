package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Функция проверяет является ли символ цифрой
func isDigit(r rune) bool {
	if _, err := strconv.Atoi(string(r)); err == nil {
		return true
	} else {
		return false
	}
}

// Функция проверяет является ли символ буквой
func isAlpha(r rune) bool {
	if !isDigit(r) && string(r) != `\` {
		return true
	}
	return false
}

// Функция записывает руну ch в builder cnt раз
func write(ch rune, b *strings.Builder, cnt int) {
	for i := 0; i < cnt; i++ {
		b.WriteString(string(ch))
	}
}

// Функция распоковывает строку
func extract(s string) (string, error) {
	arr := []rune(s)

	builder := strings.Builder{}

	ind := 0
	for ind < len(arr) {
		curCh := arr[ind]
		if isAlpha(curCh) {
			builder.WriteString(string(curCh))
		} else if isDigit(curCh) {
			cnt := 0

			j := ind
			prevInd := ind
			for j < len(arr) && isDigit(arr[j]) {
				buf, _ := strconv.Atoi(string(arr[j]))
				cnt = cnt*10 + buf
				j++
				ind++
			}
			if prevInd > 0 {
				write(arr[prevInd-1], &builder, cnt-1)
			} else {
				return "", errors.New("wrong string")
			}
			continue
		} else if string(curCh) == `\` {
			if ind < len(arr)-1 {
				builder.WriteString(string(arr[ind+1]))
				ind++
			} else {
				return "", errors.New("wrong string")
			}
		}
		ind++
	}

	return builder.String(), nil
}

func main() {
	fmt.Println(extract(`a\`))
	fmt.Println(extract(`a4bc2d5e`))
	fmt.Println(extract(`abcd`))
	fmt.Println(extract(`a11b`))
	fmt.Println(extract(`a12`))
	fmt.Println(extract(`45`))
	fmt.Println(extract(``))

	fmt.Println(extract(`qwe\4\5`))
	fmt.Println(extract(`qwe\45`))
	fmt.Println(extract(`qwe\\5`))
	fmt.Println(extract(`qwe\\`))
}
