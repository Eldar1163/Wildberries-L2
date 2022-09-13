package main

import (
	"bufio"
	"fmt"
	gops "github.com/mitchellh/go-ps"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


Так же требуется поддерживать функционал fork/exec-команд
*/

// Выводит аргумент в shell
func echo(arg string) {
	fmt.Println(arg)
}

// Выводит рабочую директорию в shell
func pwd() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(path)
	}
}

// Меняет рабочую директорию
func cd(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("Error: Bad directory")
	}
}

// Запускает новый процесс, с соответствующими аргументами
func forkExec(arg string) {
	args := strings.Split(arg, " ")

	if len(args) < 1 {
		fmt.Println("Error: Bad arguments")
	} else {
		cmd := exec.Command(args[0], args[1:]...)
		go func() {
			_ = cmd.Run()
		}()
	}
}

// Выводит список процессов в shell
func ps() {
	pcs, err := gops.Processes()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("№ PID PPID EXECUTABLE")
	for i, proc := range pcs {
		fmt.Println(i, proc.Pid(), proc.PPid(), proc.Executable())
	}
}

// Убивает процесс по PID
func kill(pidStr string) {
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	prc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	err = prc.Kill()
	if err != nil {
		fmt.Println("Error: cannot kill a process \n", err)
	}
}

// Обрабатывает команду, введенную в shell
func execCmd(cmd string) {
	switch strings.Split(cmd, " ")[0] {
	case "echo":
		echo(strings.Replace(cmd, "echo ", "", 1))
	case "pwd":
		pwd()
	case "cd":
		cd(strings.Replace(cmd, "cd ", "", 1))
	case "fork-exec":
		forkExec(strings.Replace(cmd, "fork-exec ", "", 1))
	case "ps":
		ps()
	case "kill":
		kill(strings.Replace(cmd, "kill ", "", 1))
	default:
		fmt.Println("Error: Unknown command")
	}
}

// Бесконечный цикл обработки shell
func shell() {
	sc := bufio.NewScanner(os.Stdin)
	for fmt.Print(">"); sc.Scan(); fmt.Print(">") {
		cmd := sc.Text()
		if cmd == "quit" {
			break
		} else {
			execCmd(cmd)
		}
	}
}

func main() {
	shell()
}
