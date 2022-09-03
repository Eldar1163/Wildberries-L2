package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"io"
	"os"
)

/*
=== Базовая задача ===
Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

var TimeHost = "0.beevik-ntp.pool.ntp.org"
var Stderr io.Writer = os.Stderr
var Stdout io.Writer = os.Stdout
var ExitFunc = os.Exit

func main() {
	time, err := ntp.Time(TimeHost)
	if err != nil {
		_, _ = fmt.Fprintf(Stderr, "Error has occured: %v", err)
		ExitFunc(-1)
	}
	_, _ = fmt.Fprintf(Stdout, "Current time: %v\n", time.Format("15:04:05"))
	_, _ = fmt.Fprintf(Stdout, "Exact time: %v", time.Format("15:04:05.000000000"))
}
