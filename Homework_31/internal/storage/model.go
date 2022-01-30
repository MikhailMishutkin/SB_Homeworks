package storage

//Содержимое пакета:
//
//1. метод для вызова содержимого мапы get()
//2. метод извлечения имени пользователя по id getName(user_id int)
//3. метод для записи в мапу пользователя по id put(strt User, user_id int)
//4. метод добавления в друзья: на вход id юзера которому надо добавить в друзья и id второго,
//которого надо добавить первому putFriend(user_id1 int, user_id2 int)
//5. метод удаления пользователя, по user_id удаляем самого пользователя и id из Friends друзей delete(user_id int)
//6. метод получения имён всех друзей getFN(user_id int), на выход []string
//7. метод для обновления возраста пользователя updateUserAge(user_id int, newAge int)

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"task31/internal/model"
)

//проверка соответствия структуры и интерфейса
var _ Storage = &storage{}

type storage struct {
	UserStorage map[int]*model.User
}

type Storage interface {
	GetAll(ctx context.Context)
	GetName(ctx context.Context, user_id int) string
	PutInto(ctx context.Context, strt model.User, user_id int)
	PutFriends(ctx context.Context, user_id1, user_id2 int)
	GetFNs(ctx context.Context, user_id int) []string
	UpdateUserAge(ctx context.Context, user_id int, newAge int)
	Delete(ctx context.Context, user_id int)
}

//метод для вызова содержимого мапы
func (m *storage) GetAll(ctx context.Context) {
	for k, v := range m.UserStorage {
		fmt.Printf("id: %v %v\n", k, v)
	}
}

//метод извлечения имени пользователя по id
func (m storage) GetName(ctx context.Context, user_id int) string {
	v := m.UserStorage[user_id]

	return v.Name
}

// метод для записи в мапу пользователя по id
func (m storage) PutInto(ctx context.Context, strt model.User, user_id int) {
	m.UserStorage[user_id] = &strt
}

// метод добавления в друзья: на вход id юзера которому надо добавить в друзья и id второго, которого надо добавить первому
func (m storage) PutFriends(ctx context.Context, user_id1 int, user_id2 int) {

	var a []string
	ui2 := strconv.Itoa(user_id2)

	// записываем в промежуточный слайс значение friends по id
	a = m.UserStorage[user_id1].Friends

	// добавляем нового друга
	a = append(a, ui2)

	// возвращаем в хранилище с изменениями
	m.UserStorage[user_id1].Friends = a
}

//метод получения имён всех друзей
func (m storage) GetFNs(ctx context.Context, user_id int) []string {
	var c []string
	for i := 0; i < len(m.UserStorage[user_id].Friends); i++ {
		a := m.UserStorage[user_id].Friends[i]
		b, err := strconv.Atoi(a)
		if err != nil {
			log.Println(err)
		}

		c = append(c, m.GetName(ctx, b))
	}
	return c
}

//метод для обновления возраста пользователя
func (m storage) UpdateUserAge(ctx context.Context, user_id int, newAge int) {
	a := strconv.Itoa(newAge)
	m.UserStorage[user_id].Age = a
}

// метод удаления пользователя, по user_id удаляем самого пользователя и id из Friends друзей
func (m storage) Delete(ctx context.Context, user_id int) {
	// проходим по значению id друзей в самом пользователе
	for _, v := range m.UserStorage[user_id].Friends {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
		}
		//записываем в переменную слайс друзей друга
		a := m.UserStorage[vInt].Friends
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
		m.UserStorage[vInt].Friends = a
	}

	delete(m.UserStorage, user_id)
}
