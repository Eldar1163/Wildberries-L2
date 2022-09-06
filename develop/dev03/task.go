package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное
Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Структура хранит флаги программы и имя сортируемого файла
type flags struct {
	filename     string
	sortColumn   int
	sortByNum    bool
	reversedSort bool
	uniqueValues bool
}

// Функция исправляет порядок слов в отсортированном лексикографически срезе так, чтобы
// слова с одинкаовой первой буквой, но с разным регистром находились согласно логике выполнение sort в linux
func fixUpperCaseOrder(data []string) {
	for i := 0; i < len(data)-1; i++ {
		str1 := []rune(data[i])
		str2 := []rune(data[i+1])

		if unicode.ToLower(str1[0]) == unicode.ToLower(str2[0]) &&
			unicode.IsUpper(str1[0]) &&
			unicode.IsLower(str2[0]) {
			buf := data[i]
			data[i] = data[i+1]
			data[i+1] = buf
		}
	}
}

// Быстрая сортировка строк, с возможностью воспринимать слово в строке как число (если оно таковым является)

func StringQuickSort(data []string, start, end int, byNum bool) {
	if start < end {
		base := data[start]

		left := start
		right := end

		for left < right {

			if !byNum {
				for left < right && strings.ToLower(data[right]) >= strings.ToLower(base) {
					right--
				}
			} else {
				r, err := strconv.Atoi(data[right])
				b, err := strconv.Atoi(base)
				for left < right && r >= b {
					if err != nil {
						log.Fatal("Not number:", data[right])
					}

					right--

					r, err = strconv.Atoi(data[right])
					b, err = strconv.Atoi(base)
				}
			}

			if left < right {
				data[left] = data[right]
				left++
			}

			if !byNum {
				for left < right && strings.ToLower(data[left]) <= strings.ToLower(base) {
					left++
				}
			} else {
				l, err := strconv.Atoi(data[left])
				b, err := strconv.Atoi(base)
				for left < right && l <= b {
					if err != nil {
						log.Fatal("Not number:", data[left])
					}
					left++
					l, err = strconv.Atoi(data[left])
					b, err = strconv.Atoi(base)
				}
			}

			if left < right {
				data[right] = data[left]
				right--
			}
		}

		data[left] = base

		StringQuickSort(data, start, left-1, byNum)
		StringQuickSort(data, left+1, end, byNum)
	}
}

// Функция оставляет только уникальные строки в срезе
func onlyUnique(data []string) []string {
	res := make([]string, 0, len(data))
	m := make(map[string]bool)
	for _, str := range data {
		if _, ok := m[str]; !ok {
			m[str] = true
			res = append(res, str)
		}
	}
	return res
}

// Функция обращает срез в обратном порядке
func reversed(data []string) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// Функция возвращает срез ключей отображения
func getKeysOfMap(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Функция сортирует срез строк по колонкам в строке (колонки по умолчанию делятся пробелом)
func sortByColumn(data []string, fgs flags) {
	srcMap := make(map[string]string)
	for _, str := range data {
		columns := strings.Split(str, " ")
		if len(columns) > fgs.sortColumn {
			srcMap[columns[fgs.sortColumn]] = str
		} else {
			srcMap[columns[0]] = str
		}
	}
	keysToBeSorted := getKeysOfMap(srcMap)
	StringQuickSort(keysToBeSorted, 0, len(keysToBeSorted)-1, fgs.sortByNum)
	for ind, key := range keysToBeSorted {
		data[ind] = srcMap[key]
	}
}

// Функция сортирует срез в соответствии с флагами
func sort(data []string, fgs flags) []string {
	res := make([]string, len(data))
	copy(res, data)

	if fgs.sortColumn >= 0 {
		sortByColumn(res, fgs)
		fixUpperCaseOrder(res)
	} else {
		StringQuickSort(res, 0, len(res)-1, false)
		fixUpperCaseOrder(res)
	}

	if fgs.uniqueValues {
		StringQuickSort(res, 0, len(res)-1, false)
		fixUpperCaseOrder(res)
		res = onlyUnique(res)
	}

	if fgs.reversedSort {
		reversed(res)
	}

	return res
}

// Функция парсит флаги
func parseFlags() *flags {
	s := flags{}

	flag.StringVar(&s.filename, "f", "", "указываем имя файла для сортировки")
	flag.IntVar(&s.sortColumn, "k", -1, "указываем колонку для сортировки")
	flagN := flag.Bool("n", false, "сортируем по числовому значению")
	flagR := flag.Bool("r", false, "сортируем в обратном порядке")
	flagU := flag.Bool("u", false, "выводим только уникальные значения")

	flag.Parse()

	s.sortByNum = *flagN
	s.uniqueValues = *flagU
	s.reversedSort = *flagR

	return &s
}

// Функция чтениия файла в срез строк
func readFile(filename string) []string {
	var rows []string

	file, err := os.Open(filename)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Не могу закрыть файл")
		}
	}(file)

	if err != nil {
		log.Fatal("Не могу открыть файл ", filename)
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}

	return rows
}

// Функция записи среза строк в файл
func writeToFile(filename string, data []string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("Не могу создать файл")
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Не могу закрыть файл")
		}
	}(f)

	for i := 0; i < len(data); i++ {
		_, err := fmt.Fprintln(f, data[i])
		if err != nil {
			log.Fatal("Не могу записать в файл")
		}
	}
}

func main() {
	// Парсим флаги
	flg := parseFlags()
	// Cчитываем файл
	data := readFile(flg.filename)
	// Сортируем срез
	dataSorted := sort(data, *flg)
	// Записываем отсортированный срез в файл
	writeToFile("sorted.txt", dataSorted)
	// Выводим в консоль отсортированный срез
	fmt.Println(dataSorted)
}
