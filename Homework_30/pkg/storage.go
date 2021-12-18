package pkg

import (
	"fmt"
	"log"
	"strconv"
)

var friends []string

type User struct {
	Name    string
	Age     string
	Friends []string
}

type UserStorage map[int]*User

//метод для вызова содержимого мапы
func (m UserStorage) get() {
	for k, v := range m {
		fmt.Printf("id: %v %v\n", k, v)
	}
}

//метод извлечения имени пользователя по id
func (m UserStorage) getName(user_id int) string {
	v := m[user_id]

	return v.Name
	// for k, v := range m {
	// 	fmt.Printf("id: %v %v\n", k, v)
	// }
}

/* func (m UserStorage) getID(x int) int { //метод для вызова содержимого мапы
	v := m[x]

	return v.Name
	// for k, v := range m {
	// 	fmt.Printf("id: %v %v\n", k, v)
	// }
} */

// метод для записи в мапу пользователя по id
func (m UserStorage) put(strt User, user_id int) UserStorage {
	m[user_id] = &strt

	return m
}

// метод для мапы: на вход id юзера которому надо добоваить в друзья и id второго, которого надо добавить первому
func (m UserStorage) putFriend(user_id1 int, user_id2 int) UserStorage {

	var a []string
	ui2 := strconv.Itoa(user_id2)

	// записываем в промежуточный слайс значение friends по id
	a = m[user_id1].Friends

	// добавляем нового друга
	a = append(a, ui2)

	// переменная со значениями первого id
	//b := m[user_id1]
	//записываем id друга в структуру
	//c := updateUserFriends(a, *b)
	// ui2:= strconv.Itoa(user_id2)
	// a = append(a, ui1)

	// возвращаем в хранилище с изменениями
	m[user_id1].Friends = a

	return m
}

//принимает строку пользовательского ввода, возвращает структуру
func newUser(s []string) User {
	var ns User // новый юзер структура
	ns = User{s[1], s[3], friends}
	return ns
}

func (m UserStorage) delete(user_id int) {
	//friendsSplit := strings.Split(m[user_id].Friends, " ")
	for _, v := range m[user_id].Friends {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
		}
		a := m[vInt].Friends
		for i, _ := range a {
			a[i] = a[len(a)-1]
			a[len(a)-1] = ""
			a = a[:len(a)-1]
			m[vInt].Friends = a
		}
	}

	delete(m, user_id)
}

//записываем id друга в структуру, на вход срез с id друга
/* func updateUserFriends(s []string, u User) User {

	u.Friends = append(u.Friends, s)

	return u
} */
