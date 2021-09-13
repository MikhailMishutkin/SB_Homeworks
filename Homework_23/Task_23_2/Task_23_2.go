/* Задание 2. Поиск символов в нескольких строках
Что нужно сделать
Напишите функцию, которая на вход принимает массив предложений (длинных строк)
 и массив символов типа rune, а возвращает 2D-массив, где на позиции [i][j]
 стоит индекс вхождения символа j из chars в последнее слово в предложении i
(строку надо разбить на слова и взять последнее). То есть сигнатура следующая:

func parseTest(sentences []string, chars []rune)
Советы и рекомендации
Не забудьте проверить, что вы получили больше чем 0 аргументов.
Подход не важен: можно переписать сортировку пузырьком или отсортировать, а потом перевернуть. */

package main

import (
	"fmt"
	"strings"
)

const (
	size = 3
)

func parseTest(sentences [size]string, chars [size]rune) (twoDiArray [size][size]int) {
	var a [size]string //дополнительная переменная для сбора последних слов предложений в отдельный массив
	for i := 0; i < size; i++ {
		s := sentences[i]
		spaceIndex := strings.LastIndex(s, " ") //  вычленение последних слов с использованием пробела
		s = s[spaceIndex+1:]
		a[i] = s
		//fmt.Println(len(a[i]))
		//fmt.Println(a)
	}

	for j := 0; j < size; j++ { //два цикла для прохожения значений двумерного массива
		for i := 0; i < size; i++ {
			b := string(chars[j])       // переводим руну в строку и присваиваем это переменной
			k := strings.Index(a[i], b) // метод нахождения индекса буквы в строке, присваеваем значение переменной k

			if k > -1 { // принимает значение начиная с нуля, если нет символа в строке, то -1
				twoDiArray[i][j] = k
			} else {
				twoDiArray[i][j] = -1
			}
		}
	}
	return
}

func main() {
	s := [size]string{"some people have", "shining through the eyes", "and excess weight"}
	ch := [size]rune{'h', 'a', 'e'}

	fmt.Println(parseTest(s, ch))

}
