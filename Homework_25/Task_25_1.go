/* 25.8 Задача
Написать программу для нахождения подстроки в кириллической подстроке. Программа
должна запускаться с помощью команды:

 go run main.go --str "строка для поиска" --substr "поиска"

Для реализации такой работы с флагами воспользуйтесь пакетом flags, а для поиска
подстроки в строке вам понадобятся руны.
Что нужно сделать
Спроектировать алгоритм поиска подстроки.
Определить строку и подстроку, используя флаги.
Написать алгоритм реализацию для работы со строками UTF-8 (для этого необходимо
воспользоваться рунами).
Что оценивается
Алгоритм может работать с различными символами (кириллица, китайские иероглифы).
Использованы руны. */

package main

import (
	"flag"
	"fmt"
)

func compareStringsInFlags(s1 *string, s2 *string) bool {
	s1Rune := []rune(*s1)
	s2Rune := []rune(*s2)
	coincidenceStr := false // присваеваем переменной совпадения строк значение false

	for i := 0; i < len(s1Rune); i++ { // проходим всю строку, в которой ищем
		for j := 0; j < len(s2Rune); j++ { // ищем в пределах длинны искомой строки
			if s1Rune[i] == s2Rune[j] {
				coincidenceStr = true
				i++
			} else {
				coincidenceStr = false
				break
			}

		}

	}
	return coincidenceStr
}

func main() {
	var str, substr string

	flag.StringVar(&str, "str", "default str", "set str")
	flag.StringVar(&substr, "substr", "default substr", "set substr")

	flag.()
	flag.Parse()
	fmt.Println(compareStringsInFlags(&str, &substr))

}

/* 	for _, v := range *s1 { // переводим входящие строки в массивы руны
	s1Rune = append(s1Rune, v)
}

for _, v := range *s2 {
	s2Rune = append(s2Rune, v)
} */
