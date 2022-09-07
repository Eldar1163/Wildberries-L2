package main

import (
	"fmt"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===
Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Функция быстрой сортировки строк среза
func stringQuickSort(data []string, start, end int) {
	if start < end {
		base := data[start]

		left := start
		right := end

		for left < right {

			for left < right && data[right] >= base {
				right--
			}

			if left < right {
				data[left] = data[right]
				left++
			}

			for left < right && data[left] <= base {
				left++
			}

			if left < right {
				data[right] = data[left]
				right--
			}
		}

		data[left] = base

		stringQuickSort(data, start, left-1)
		stringQuickSort(data, left+1, end)
	}
}

// Функция обертка быстрой сортировки строк среза
func stringSort(arr []string) []string {
	res := make([]string, len(arr))
	copy(res, arr)
	stringQuickSort(res, 0, len(res)-1)
	return res
}

// Функция приводит в нижний регистр в строки в срезе
func strToLower(data []string) []string {
	res := make([]string, len(data))
	copy(res, data)
	for i, str := range data {
		res[i] = strings.ToLower(str)
	}
	return res
}

// Функция сортирует строку по символам (пузырьковой сортировкой)
func sortStr(str string) string {
	arr := make([]rune, len([]rune(str)))
	copy(arr, []rune(str))

	for i := 1; i < len(arr); i++ {
		f := true
		for j := 0; j < len(arr)-i; j++ {
			if string(arr[j]) > string(arr[j+1]) {
				buf := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = buf
				f = false
			}
		}
		if f {
			break
		}
	}

	return string(arr)
}

// Функция решает поставленную задачу (строит мапу с анаграммами)
func findAnagrams(dict []string) map[string][]string {
	// В нижний регистр
	dict = strToLower(dict)
	// Сортировка
	dict = stringSort(dict)
	// Промежуточная мапа для результата
	bufMap := make(map[string][]string)
	// Мапа имен
	nameMap := make(map[string]string)
	// Строим промежуточную мапу
	for _, word := range dict {
		if len(word) > 1 {
			wsort := sortStr(word)
			if _, ok := bufMap[wsort]; ok {
				bufMap[wsort] = append(bufMap[wsort], word)
			} else {
				bufMap[wsort] = make([]string, 0, 1)
				nameMap[wsort] = word
			}
		}
	}
	// Строим результирующую мапу
	resMap := make(map[string][]string)
	for k, v := range bufMap {
		resMap[nameMap[k]] = v
	}

	return resMap
}

func main() {
	fmt.Println(findAnagrams([]string{"пятак", "листок", "пятка", "слиток", "тяпка", "столик"}))
}
