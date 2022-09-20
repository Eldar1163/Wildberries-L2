package main

import (
	"errors"
	"fmt"
	flag "github.com/spf13/pflag"
	"io"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

/*
Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
1.	Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
	После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
2.	Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
3.	При нажатии Ctrl+C программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера,
	программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout
*/

// Флаги запуска программы
type Flags struct {
	timeout time.Duration
	host    string
	port    int
}

// Функция парсит флаги
func parseFlags() *Flags {
	flags := Flags{}
	var timeoutStr string
	flag.StringVar(&timeoutStr, "timeout", "10s", "connection timeout")
	flag.Parse()

	var err error
	flags.timeout, err = parseTimeout(timeoutStr)
	if err != nil {
		log.Fatal("Wrong timeout: ", timeoutStr)
	}

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("You must specify host")
	}

	hostURL, err := url.Parse(args[0])
	if err != nil {
		log.Fatal("Wrong host: ", args[0])
	}
	_, err = strconv.Atoi(args[0])
	if err == nil {
		log.Fatal("Wrong host: ", args[0])
	}

	flags.host = hostURL.String()

	if len(args) == 2 {
		portNum, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("Incorrect port: ", args[1])
		}
		flags.port = portNum
	}

	return &flags
}

// Функция парсит строку с timeout (например: 10s, 5m, 4h)
func parseTimeout(timeout string) (time.Duration, error) {
	if len(timeout) <= 1 {
		return 10 * time.Second, errors.New("bad timeout")
	}
	valueStr := timeout[:len(timeout)-1]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 10 * time.Second, err
	}

	switch timeout[len(timeout)-1] {
	case 's':
		return time.Duration(value) * time.Second, nil
	case 'm':
		return time.Duration(value) * time.Minute, nil
	case 'h':
		return time.Duration(value) * time.Hour, nil
	default:
		return 10 * time.Second, errors.New("Wrong measure: " + string(timeout[len(timeout)-1]))
	}
}

// Функция реализует telnet клиент
func telnet(flags *Flags) {
	// Создаем канал для системных сигналов
	gracefulShutdown := make(chan os.Signal, 1)
	// Принимаем в этот канал оповщения о SIGINT, SIGTERM, SIGQUIT
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Формируем строку подключения
	var fullHost string
	if flags.port != 0 {
		fullHost = flags.host + ":" + strconv.Itoa(flags.port)
	} else {
		fullHost = flags.host + ":80"
	}

	// Открываем сокет
	conn, err := net.DialTimeout("tcp", fullHost, flags.timeout)

	// Обрабатываем ошибки
	var dnsErr *net.DNSError
	switch {
	case errors.As(err, &dnsErr):
		time.Sleep(flags.timeout)
		log.Println("DNS error: ", err)
		os.Exit(0)
	case err != nil:
		log.Fatal("Cannot open connection: ", err)
	}

	// Отслеживаем graceful-shutdown
	go func(conn net.Conn) {
		<-gracefulShutdown
		err := conn.Close()
		if err != nil {
			log.Println("Cannot close socket")
			os.Exit(1)
		}
		os.Exit(0)
	}(conn)

	// Устанавливаем первоначальный таймаут на ответ в 5 секунд
	_ = conn.SetReadDeadline(time.Now().Add(time.Duration(5) * time.Second))
	for {
		fmt.Print(">")

		// Отправляем данные из консольного ввода в сокет
		_, err = io.Copy(conn, os.Stdin)

		_ = conn.SetReadDeadline(time.Now().Add(time.Duration(700) * time.Millisecond))

		// Принимаем данные из сокета в консольный вывод
		_, err := io.Copy(os.Stdout, conn)
		// Проверяем не закрылся ли сокет
		if errors.Is(err, net.ErrClosed) {
			fmt.Println("Socket is closed")
			os.Exit(0)
		}

		fmt.Println()
	}
}

func main() {
	// Парсим флаги
	flags := parseFlags()

	// Запускаем telnet
	telnet(flags)
}
