/* Напишите программу, которая считывает ввод с stdin, создаёт структуру student и записывает указатель на структуру
в хранилище map[studentName] *Student.

type Student struct {

name string

age int

grade int

}
Программа должна получать строки в бесконечном цикле, создать структуру Student через функцию newStudent, далее
сохранить указатель на эту структуру в map, а после получения EOF (ctrl + d) вывести на экран имена всех
 студентов из хранилища. Также необходимо реализовать методы put, get.

Input
go run main.go
Строки
Вася 24 1
Семен 32 2
EOF

Output
Студенты из хранилища:
Вася 24 1
Семен 32 2

Зачёт:
при получении одной строки (например, «имяСтудента 24 1») программа создаёт студента и сохраняет его, далее
ожидает следующую строку или сигнал EOF (Сtrl + d);
при получении сигнала EOF программа должна вывести имена всех студентов из map. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	name  string
	age   int
	grade int
}

type studStorage map[string]*Student //задаём мапе отдельный тип

func (m studStorage) put(str Student) studStorage { // пишем метод для записи в мапу
	m[str.name] = &str

	return m
}

func (m studStorage) get() { //метод для вызова содержимого мапы
	fmt.Println("Студенты из хранилища:")
	for _, v := range m {
		fmt.Printf("%#v\n", v)
	}
}

func newStudent(s []string) Student { //принимает строку пользовательского ввода, возвращает структуру
	var ns Student // новый студент структура

	i, err := strconv.Atoi(s[1])
	if err != nil {
		fmt.Println(err)
	}
	k, err := strconv.Atoi(s[2])
	if err != nil {
		fmt.Println(err)
	}

	ns = Student{s[0], i, k}
	return ns
}

func main() {
	// переменная для мапы-хранилища
	var m studStorage
	m = make(studStorage)

	// показываем каким образом получим значение переменной
	input := bufio.NewScanner(os.Stdin)
	// переменная сплитуется сразу после получения значения
	input.Split(bufio.ScanWords)
	// создаём слайс для значения
	valueIn := make([]string, 0)

	for input.Scan() {
		valueIn = append(valueIn, input.Text()) // складываем в слайс все значения пользовательского ввода

		if len(valueIn) == 3 {
			//переменная для созданной из пользовательского ввода структуры
			p := newStudent(valueIn)
			//fmt.Println(p.name)
			//
			m.put(p)
			// обнуляем значение переменной для ввода данных нового студента
			valueIn = []string{}

		}

	}
	m.get()
}

/* 	p := newStudent(name)
   	fmt.Println(p) */
/* ns := map[string]*Student{
	p.name: &p,
}
fmt.Println(ns) */
