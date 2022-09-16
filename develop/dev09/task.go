package main

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
Реализовать утилиту wget с возможностью скачивать сайты целиком.
*/

// Функция скачивает страницу по url
func downloadPage(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Cannot download page with url = ", url)
		return nil
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Cannot read page data with url = ", url)
	}

	return data
}

// Функция записывает данные в файл
func writeToFile(data []byte, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// Функция создает рабочую директорию и выполняет wget
func wget(url string, outPath string, depth int, timeout time.Duration) {
	if depth < 1 {
		log.Fatal("Incorrect depth")
	}

	wd, _ := os.Getwd()
	err := os.Mkdir(wd+`\`+outPath, os.ModeDir)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Cannot create directory")
		}
	}
	err = os.Chdir(outPath)
	if err != nil {
		log.Fatal("Cannot change work dir")
	}
	wgetRec(url, url, depth, 0, 1, timeout)
}

// Реализация wget (рекурсивная)
func wgetRec(rootUrl string, curUrl string, depth int, curDepth int, pageInd int, timeout time.Duration) {
	if curDepth == depth {
		return
	}

	data := downloadPage(curUrl)

	if data != nil {
		links := getLinks(data)

		linksNorm := make([]string, len(links))
		copy(linksNorm, links)
		normalizeLinks(linksNorm, rootUrl, curUrl)

		for i := 0; i < len(links); i++ {
			oldLink := []byte(links[i])
			newLink := []byte(strconv.Itoa(curDepth+1) + "_" + strconv.Itoa(i) + ".html")
			copy(data, bytes.ReplaceAll(data, oldLink, newLink))
		}

		err := writeToFile(data, strconv.Itoa(curDepth)+"_"+strconv.Itoa(pageInd)+".html")
		if err != nil {
			log.Println("Cannot write file: ", err)
		}

		for i := 0; i < len(linksNorm); i++ {
			time.Sleep(timeout)
			wgetRec(rootUrl, linksNorm[i], depth, curDepth+1, i, timeout)
		}
	}
}

// Функция возвращает слайс ссылок из тела заданной страницы
func getLinks(body []byte) []string {
	var links []string
	bodyReader := bytes.NewReader(body)
	z := html.NewTokenizer(bodyReader)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}

				}
			}

		}
	}
}

// Функция приводит ссылки к полному виду (с протоколом и доменом)
func normalizeLinks(links []string, rootUrl string, parentUrl string) {
	for i, link := range links {
		if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
			if strings.HasPrefix(link, "/") {
				url, err := url2.Parse(rootUrl)
				if err != nil {
					log.Println("Cannot parse url: ", url)
				}
				links[i] = url.Scheme + "://" + urlAddSuffixIfNeeded(url.Host) + link[1:]
			} else {
				links[i] = urlAddSuffixIfNeeded(parentUrl) + link
			}
		}
	}
}

// Функция добавляет slash в конец ссылки
func urlAddSuffixIfNeeded(url string) string {
	if !strings.HasSuffix(url, "/") {
		return url + "/"
	} else {
		return url
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter full url, example: https://example.com/")
	scanner.Scan()
	url := scanner.Text()

	fmt.Println("Enter full directory path to store html files")
	scanner.Scan()
	path := scanner.Text()

	wget(url, path, 2, 3)

	fmt.Println("Download completed")
}
