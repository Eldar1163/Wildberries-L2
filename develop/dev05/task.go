package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

/*
=== Утилита grep ===
Реализовать утилиту фильтрации (man grep)
Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Структура флагов программы
type flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	rExp       string
	filename   string
}

// Функция считывает флаги из аргументов запуска программы
func parseFlags() *flags {
	f := flags{}

	flag.IntVar(&f.after, "A", 0, `"after" печатать +N строк после совпадения`)
	flag.IntVar(&f.before, "B", 0, `"before" печатать +N строк до совпадения`)
	flag.IntVar(&f.context, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.StringVar(&f.rExp, "r", "", "регулярное выражение")
	flag.StringVar(&f.filename, "f", "", "имя файла")

	flagC := flag.Bool("c", false, "количество строк")
	flagI := flag.Bool("i", false, "игнорировать регистр")
	flagV := flag.Bool("v", false, "вместо совпадения, исключать")
	flagF := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	flagN := flag.Bool("n", false, "напечатать номер строки")

	flag.Parse()

	f.count = *flagC
	f.invert = *flagV
	f.ignoreCase = *flagI
	f.fixed = *flagF
	f.lineNum = *flagN

	return &f
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

// Функция печатает результат вывода grep на экран
func printResults(res interface{}) {
	switch result := res.(type) {
	case []string:
		for _, str := range result {
			fmt.Println(str)
		}
	case []int:
		for _, num := range result {
			fmt.Println(num)
		}
	case int:
		fmt.Println(result)
	}
}

// Функция копирует из слайса в мапу данные из заданного интервала в слайсе
func copyToMapByInterval(m map[int]string, data []string, first int, last int) {
	for k := first; k < last; k++ {
		m[k] = data[k]
	}
}

// Функция возвращает слайс чисел (отсортированных ключей мапы)
func getSortedMapKeys(m map[int]string) []int {
	keys := make([]int, 0, 1)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

// Функция, на основе промежуточной мапы и ее отсортированных ключей, формирует результат
func evalResult(m map[int]string, keys []int) []string {
	res := make([]string, 0, len(m))
	for _, key := range keys {
		res = append(res, m[key])
	}

	return res
}

// Реализует ключ -B
func grepBefore(rExp *regexp.Regexp, f flags, data []string) map[int]string {
	mapBuf := make(map[int]string)
	for ind, str := range data {
		if rExp.MatchString(str) && !f.invert || !rExp.MatchString(str) && f.invert {
			mapBuf[ind] = str
			if ind-f.before >= 0 {
				copyToMapByInterval(mapBuf, data, ind-f.before, ind)
			} else {
				copyToMapByInterval(mapBuf, data, 0, ind)
			}
		}
	}
	return mapBuf
}

// Реализует ключ -A
func grepAfter(rExp *regexp.Regexp, f flags, data []string) map[int]string {
	mapBuf := make(map[int]string)
	for ind, str := range data {
		if rExp.MatchString(str) && !f.invert || !rExp.MatchString(str) && f.invert {
			mapBuf[ind] = str
			if ind+f.after < len(data) {
				copyToMapByInterval(mapBuf, data, ind+1, ind+f.after+1)
			} else {
				copyToMapByInterval(mapBuf, data, ind+1, len(data))
			}
		}
	}
	return mapBuf
}

// Реализует ключ -C
func grepContext(rExp *regexp.Regexp, f flags, data []string) map[int]string {
	mapBuf := make(map[int]string)
	for ind, str := range data {
		if rExp.MatchString(str) && !f.invert || !rExp.MatchString(str) && f.invert {
			mapBuf[ind] = str

			if ind-f.context >= 0 {
				copyToMapByInterval(mapBuf, data, ind-f.context, ind)
			} else {
				copyToMapByInterval(mapBuf, data, 0, ind)
			}

			if ind+f.context < len(data) {
				copyToMapByInterval(mapBuf, data, ind+1, ind+f.context+1)
			} else {
				copyToMapByInterval(mapBuf, data, ind+1, len(data))
			}
		}
	}

	return mapBuf
}

// Реализует простейший grep
func grepSimple(rExp *regexp.Regexp, f flags, data []string) map[int]string {
	mapBuf := make(map[int]string)
	for ind, str := range data {
		if rExp.MatchString(str) && !f.invert || !rExp.MatchString(str) && f.invert {
			mapBuf[ind] = str
		}
	}

	return mapBuf
}

// Решает поставленную задачу (выполняетя grep с различными ключами)
func grep(data []string, f flags) (interface{}, error) {
	var (
		prefix  string
		postfix string
	)

	if f.ignoreCase {
		prefix = "(?i)"
	}

	if f.fixed {
		prefix += "^"
		postfix += "$"
	}

	rExp, err := regexp.Compile(prefix + f.rExp + postfix)
	if err != nil {
		log.Fatal("Bad regexp")
	}

	if f.count {
		cnt := 0
		if f.before > 0 {
			beforeRes := grepBefore(rExp, f, data)
			cnt = len(beforeRes)
		} else if f.after > 0 {
			afterRes := grepAfter(rExp, f, data)
			cnt = len(afterRes)
		} else if f.context > 0 {
			contextRes := grepContext(rExp, f, data)
			cnt = len(contextRes)
		} else {
			simpleRes := grepSimple(rExp, f, data)
			cnt = len(simpleRes)
		}

		return cnt, nil
	} else if f.lineNum {
		var lineNumList []int

		if f.before > 0 {
			beforeRes := grepBefore(rExp, f, data)
			lineNumList = getSortedMapKeys(beforeRes)
		} else if f.after > 0 {
			afterRes := grepAfter(rExp, f, data)
			lineNumList = getSortedMapKeys(afterRes)
		} else if f.context > 0 {
			contextRes := grepContext(rExp, f, data)
			lineNumList = getSortedMapKeys(contextRes)
		} else {
			simpleRes := grepSimple(rExp, f, data)
			lineNumList = getSortedMapKeys(simpleRes)
		}

		return lineNumList, nil
	} else if f.before > 0 {
		mapBuf := grepBefore(rExp, f, data)
		keys := getSortedMapKeys(mapBuf)
		res := evalResult(mapBuf, keys)

		return res, nil
	} else if f.after > 0 {
		mapBuf := grepAfter(rExp, f, data)
		keys := getSortedMapKeys(mapBuf)
		res := evalResult(mapBuf, keys)

		return res, nil
	} else if f.context > 0 {
		mapBuf := grepContext(rExp, f, data)
		keys := getSortedMapKeys(mapBuf)
		res := evalResult(mapBuf, keys)

		return res, nil
	} else {
		mapBuf := grepSimple(rExp, f, data)
		keys := getSortedMapKeys(mapBuf)
		res := evalResult(mapBuf, keys)

		return res, nil
	}
}

func main() {
	// Парсим флаги
	flgs := parseFlags()
	// Cчитываем файл
	data := readFile(flgs.filename)
	// Выполняем операцию grep
	res, _ := grep(data, *flgs)
	// Выводим результат
	printResults(res)
}
