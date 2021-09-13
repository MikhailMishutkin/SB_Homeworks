//Практика в подмножестве

package main

import "fmt"

func main() {
	students := []string{"Ali", "Mea", "Nata", "Pete", "John", "Tony"}
	studentsMansSubSlice := students[3:]

	fmt.Println(students)
	fmt.Println(studentsMansSubSlice)

	studentsMansSubSlice[0] = "David" // присваиваем новое значение первому элементу сабслайса, изменяется и первый слайс, 4-й элемент
	fmt.Println(students)
	fmt.Println(studentsMansSubSlice)

	modifySlice(studentsMansSubSlice) // присваеиваем новое значение через функцию, результат тот же
	fmt.Println(students)
	fmt.Println(studentsMansSubSlice)
}

func modifySlice(slice []string) {
	slice[0] = "Вася"
}
