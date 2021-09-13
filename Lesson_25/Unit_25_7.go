//Понимание аллокаций памяти, pointer и value семантика
// спецаильный синтаксис после go run и пути к файлу надо ввести --name John --age 23
package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	var age int

	flag.IntVar(&age, "age", 0, "set age") //пакет flag указатели на переменные
	flag.StringVar(&name, "name", "default name", "set name")

	flag.Parse()
	fmt.Println(name, age)

	/* printByPointer(&name, &age) //передаём по указателю значения нэйм и эйдж в функцию

	err := printByPointer(nil, nil)
	if err != nil {
		log.Fatalln(err)
	} */
}

/* func printByPointer(n *string, a *int) error {
	if n == nil || a == nil { // проверка на пустые значения указателей
		return fmt.Errorf("nil pointer")
	}
	fmt.Println("Hello my name is " + *n + " and my age is " + strconv.Itoa(*a))
	return nil
} */
