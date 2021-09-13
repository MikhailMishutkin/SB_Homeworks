/* Приложение, которое позволяет скачивать произвольную HTML-страницу посредством HTTP-запроса
на жесткий диск компьютера и выдает статистику по количеству уникальных слов в консоль.
Требования к приложению
1 В качестве входных данных в приложение принимает строку с адресом web-страницы.
Пример входной строки: https://www.simbirsoft.com/
2 Приложение разбивает текст страницы на отдельные слова с помощью списка разделителей.
Пример списка:
{' ', ',', '.', '! ', '?','"', ';', ':', '[', ']', '(', ')', '\n', '\r', '\t'}
3 В качестве результата работы пользователь должен получить статистику по
количеству уникальных слов в тексте.
Пример:
РАЗРАБОТКА -1
ПРОГРАММНОГО - 2
ОБЕСПЕЧЕНИЯ - 4 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func onlyLetters(s string) string { // возможно нет необходимости в функции
	rx := regexp.MustCompile(`[^А-Яа-я ]`)
	s = rx.ReplaceAllString(s, " ")
	return s
}
func MakeRequest() []byte {
	resp, err := http.Get("https://www.simbirsoft.com/")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body

}

func main() {
	MakeRequest()
	// при отсутствии интеренента
	/* f, err := os.ReadFile("example_after_regexp.txt")
	if err != nil {
		fmt.Println("Файл не найден")
	} */
	body := string(MakeRequest()) // закоментировал пока работаю с файлом
	//body1 := (string(f)) // при работе с файлом

	rx := regexp.MustCompile(`[^А-Яа-я' '',''.''!''?''"'';'':''[''('')''\n''\r''\t']`)
	body = rx.ReplaceAllString(body, "") //????
	//log.Println(string(body1))
	body = onlyLetters(body)
	//bodySlice := strings.SplitAfter(body, " ")
	//fmt.Println(onlyLetters(body1))
	//bodyRune := []rune(body1) попробовал в руну перевести
	//fmt.Println(bodyRune)
	s := strings.Fields(body)
	/* var s []string
	for strings.Contains(body, " ") {
		spaceIndex := strings.Index(body, " ")
		word := body[:spaceIndex] // переписать

		if len(word) > 4 {
			d := strings.Split(string(word), " ")
			s = append(s, d...)

		}
		body = body[spaceIndex+1:]
		body = strings.Trim(body, " ")
	} */
	fmt.Println(s)
	body = strings.Join(s, " ")
	bodyRune := []rune(body)
	fmt.Println(bodyRune)

}

//'\n''\r''\t' ']'
