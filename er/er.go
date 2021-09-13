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

func getFiles(fileName1 string, fileName2 string) string {
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

func main() {
	firstFile := "1.txt"
	secondFile := "2.txt"
	f1, err := os.Open("1.txt")
	if err != nil {
		fmt.Println("Файл с таким именем не обнаружен")
		return
	}
	defer f1.Close()

	f2, err := os.Open("2.txt")
	if err != nil {
		fmt.Println(getFirstContent(firstFile))
		return
	} else {
		fmt.Println(getFiles(firstFile, secondFile))
	}
	defer f2.Close()
}

/* args := os.Args[1:] //смотрим,есть ли файлы
if len(args) == 0 {
	fmt.Println("Укажите директорию") // если нет, просим указать директорию с файлами
	return
}

files, err := ioutil.ReadDir(args[0]) // читаем файлИнфо в переменную Files
if err != nil {
	fmt.Println(err)
	return
}
name := make([]string, 0)    // создаём слайс для имён файлов в директории
	for _, file := range files { // записываем имена файлов в слайс из файлИнфо
		name = append(name, file.Name())
	}
		if len(name) < 2 { // если файлов меньше двух, то выводим имя и содержание первого
		fmt.Println(name[0])
		fmt.Println(getFirstContent(name[0]))
	} else { // если больше одного файла - соединяем первый и второй и печатаем содержимое на экран
		fmt.Println(getFiles(name[0], name[1]))
	}
*/
