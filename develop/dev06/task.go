package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===
Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Структура хранит флаги для работы программы
type flags struct {
	fields    []int
	delimiter string
	separated bool
}

// Функция парсит флаги из аргументов запуска программы
func parseFlags() *flags {
	f := flags{}

	var fieldList string
	flag.StringVar(&fieldList, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&f.delimiter, "d", "	", "установить разделить")
	flagS := flag.Bool("s", true, "только строки с разделителем")

	flag.Parse()

	f.separated = *flagS

	var err error
	f.fields, err = parseFieldList(fieldList)
	if err != nil {
		log.Fatal("Bad field list")
	}

	return &f
}

// Функция парсит список колонок
func parseFieldList(fieldList string) ([]int, error) {
	fieldList = strings.TrimSpace(fieldList)
	fieldList = removeSpaces(fieldList)
	intStrList := strings.Split(fieldList, ",")
	res := make([]int, len(intStrList))
	for ind, intStr := range intStrList {
		num, err := strconv.Atoi(intStr)
		if err != nil {
			return nil, errors.New("bad field list")
		}
		res[ind] = num
	}
	return res, nil
}

// Функция удаляет пробелы и табы из строки
func removeSpaces(str string) string {
	res := make([]rune, 0, 1)
	for _, r := range str {
		if string(r) != " " && string(r) != "	" {
			res = append(res, r)
		}
	}
	return string(res)
}

// Функция реализует утилиту cut
func cut(str string, f flags) {
	columns := strings.Split(str, f.delimiter)
	if len(columns) == 1 && !f.separated {
		fmt.Println(str)
	} else if len(columns) > 1 && f.separated {
		for _, columnInd := range f.fields {
			if columnInd <= len(columns) {
				fmt.Print(columns[columnInd-1] + " ")
			}
		}
		fmt.Println()
	}
}

// Цикл чтения и выполнения cut
func readAndCutCycle(f flags) {
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		if err != io.EOF {
			cut(str, f)
		} else {
			break
		}
	}
}

func main() {
	f := flags{}
	// Парсим флаги
	f = *parseFlags()
	// Запускаем бесконечный цикл выполнения cut (выход по EOF)
	readAndCutCycle(f)
}
