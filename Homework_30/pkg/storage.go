package pkg

//Содержимое пакета:
//1. функция newUser(s []string) создания пользователя: принимает строку пользовательского ввода, возвращает структуру
//2. метод для вызова содержимого мапы get()
//3. метод извлечения имени пользователя по id getName(user_id int)
//4. метод для записи в мапу пользователя по id put(strt User, user_id int)
//5. метод добавления в друзья: на вход id юзера которому надо добавить в друзья и id второго,
//которого надо добавить первому putFriend(user_id1 int, user_id2 int)
//6. метод удаления пользователя, по user_id удаляем самого пользователя и id из Friends друзей delete(user_id int)
//7. метод получения имён всех друзей getFN(user_id int), на выход []string
//8. метод для обновления возраста пользователя updateUserAge(user_id int, newAge int)

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

//создание пользователя: принимает строку пользовательского ввода, возвращает структуру
func newUser(s []string) User {
	// новый юзер структура
	var ns User
	ns = User{s[1], s[3], friends}
	return ns
}

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
}

// метод для записи в мапу пользователя по id
func (m UserStorage) put(strt User, user_id int) UserStorage {
	m[user_id] = &strt

	return m
}

// метод добавления в друзья: на вход id юзера которому надо добавить в друзья и id второго, которого надо добавить первому
func (m UserStorage) putFriend(user_id1 int, user_id2 int) UserStorage {

	var a []string
	ui2 := strconv.Itoa(user_id2)

	// записываем в промежуточный слайс значение friends по id
	a = m[user_id1].Friends

	// добавляем нового друга
	a = append(a, ui2)

	// возвращаем в хранилище с изменениями
	m[user_id1].Friends = a

	return m
}

// метод удаления пользователя, по user_id удаляем самого пользователя и id из Friends друзей
func (m UserStorage) delete(user_id int) {
	// проходим по значению id друзей в самом пользователе
	for _, v := range m[user_id].Friends {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
		}
		//записываем в переменную слайс друзей друга
		a := m[vInt].Friends
		// удаляем id из слайса друзей друга по индекс
		for i, v := range a {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				log.Println(err)
			}
			if vInt == user_id {
				a[i] = a[len(a)-1]
				a[len(a)-1] = ""
				a = a[:len(a)-1]
			}
		}
		m[vInt].Friends = a
	}

	delete(m, user_id)
}

//метод получения имён всех друзей
func (m UserStorage) getFN(user_id int) []string {
	var c []string
	for i := 0; i < len(m[user_id].Friends); i++ {
		a := m[user_id].Friends[i]
		b, err := strconv.Atoi(a)
		if err != nil {
			log.Println(err)
		}

		c = append(c, m.getName(b))
	}
	return c
}

//метод для обновления возраста пользователя
func (m UserStorage) updateUserAge(user_id int, newAge int) {
	a := strconv.Itoa(newAge)
	m[user_id].Age = a
}
