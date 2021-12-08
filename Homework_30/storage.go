package storage

import (
	"fmt"
)

func (m UserStorage) get() { //метод для вызова содержимого мапы
	//fmt.Println("Студенты из хранилища:")
	for k, v := range m {
		fmt.Printf("id: %v %v\n", k, v)
	}
}

func (m UserStorage) put(strt User, uuid int) UserStorage { // пишем метод для записи в мапу
	m[uuid] = &strt

	return m
}

func newUser(s []string) User { //принимает строку пользовательского ввода, возвращает структуру
	var ns User // новый юзер структура
	ns = User{s[1], s[3], friends}
	return ns
}
