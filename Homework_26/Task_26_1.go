/* Задача
Цель задания
Написать программу аналог cat.
Что нужно сделать
Программа должна получать на вход имена двух файлов и конкатенировать их, используя strings.Join.
Что оценивается
При получении одного файла на входе программа должна печатать его содержимое на экран.
При получении двух файлов на входе программа соединяет их и печатает содержимое обоих файлов на экран.
Если программа запущена командой go run firstFile.txt secondFile.txt resultFile.txt, то
она должна написать два соединённых файла в результирующий. */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getFiles(fileName1 string, fileName2 string) string { // объединение содержимого двух файлов методом Join
	f1, err := os.Open(fileName1)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	contF1, err := ioutil.ReadAll(f1) // читаем содержимое первого файла в байтах
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(contF1))
	c := strings.Split(string(contF1), " ") // переводим байты в строку, а строку разделяем в слайс строк

	f2, err := os.Open(fileName2)
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	contF2, err := ioutil.ReadAll(f2) // читаем содержимое второго файла в байтах
	if err != nil {
		panic(err)
	}
	//	fmt.Println(string(contF2))
	d := strings.Split(string(contF2), " ")

	e := append(c, d...) // соединяем слайсы строк в один слайс

	result := strings.Join(e, " ") // получаем строку из слайса строк

	return result
}

func getFirstContent(fileName1 string) string { // функция, вызываемая в случае получения данных только об одном файле
	f1, err := os.Open(fileName1)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	contF1, err := ioutil.ReadAll(f1) // читаем содержимое первого файла в байтах
	if err != nil {
		panic(err)
	}
	return string(contF1)
}

func writeToThird() []byte { // реализация метода для записи в третий файл содержимого двух
	f1, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	contF1, err := ioutil.ReadAll(f1) // читаем содержимое первого файла в байтах в переменную
	if err != nil {
		panic(err)
	}

	f2, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	contF2, err := ioutil.ReadAll(f2)
	if err != nil {
		panic(err)
	}

	e := make([]byte, 0, 0)
	e = append(contF1, contF2...)

	return e
}

func main() {

	switch {
	case len(os.Args) == 1: // если в строке нет аргументов, то идём по сценарию поиска и операций с файлами 1 и 2
		firstFile := "1.txt"
		secondFile := "2.txt"
		f1, err := os.Open("1.txt")
		if err != nil {
			fmt.Println("Файл с таким именем не обнаружен") // в папке нет файлов
			return
		}
		defer f1.Close()

		f2, err := os.Open("2.txt")
		if err != nil {
			fmt.Println(getFirstContent(firstFile)) //есть только первый файл
			return
		} else {
			fmt.Println(getFiles(firstFile, secondFile)) // есть оба файла, но не через командную
		}
		defer f2.Close()
	case len(os.Args) == 2: // дублируем сценарий, если ввёдён один аргумент, но через командную строку
		fmt.Println(getFirstContent(os.Args[1]))
	case len(os.Args) == 3: //дублируем сценарий, если ввёдены два аргумент, но через командную строку
		fmt.Println(getFiles(os.Args[1], os.Args[2]))
	case len(os.Args) == 4: // реализуем сценарий с записью в третий файл при вводе трёх аргументов через командную строку
		ioutil.WriteFile(os.Args[3], writeToThird(), 0666)
	default:
		fmt.Println("Файл не обнаружен")

	}
}
