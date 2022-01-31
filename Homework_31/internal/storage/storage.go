package storage

import (
	"context"
	"encoding/json"
	"log"
	"task31/internal/model"
)

//здесь будет описано взаимодействие с БД
//var Friends []string

type storage struct{}

func (m storage) InsertToDB(u model.User) []byte {
	jsonData, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	return jsonData
}

//метод для вызова содержимого хранилища
func (m storage) GetAll(ctx context.Context) {
	/* for k, v := range m.UserStorage {
		fmt.Printf("id: %v %v\n", k, v)
	} */
}

//метод извлечения имени пользователя по id
func (m storage) GetName(ctx context.Context, user_id int) string {
	//v := m.model.User[user_id]

	return "" // заглушка
}

// метод для записи в мапу пользователя по id
func (m storage) PutInto(ctx context.Context, strt model.User, user_id int) {
	//m.UserStorage[user_id] = &strt
}

// метод добавления в друзья: на вход id юзера которому надо добавить в друзья и id второго, которого надо добавить первому
func (m storage) PutFriends(ctx context.Context, user_id1 int, user_id2 int) {
	/*
		var a []string
		ui2 := strconv.Itoa(user_id2)

		// записываем в промежуточный слайс значение friends по id
		a = m.UserStorage[user_id1].Friends

		// добавляем нового друга
		a = append(a, ui2)

		// возвращаем в хранилище с изменениями
		m.UserStorage[user_id1].Friends = a */
}

//метод получения имён всех друзей
func (m storage) GetFNs(ctx context.Context, user_id int) []string {
	/* var c []string
	for i := 0; i < len(m.UserStorage[user_id].Friends); i++ {
		a := m.UserStorage[user_id].Friends[i]
		b, err := strconv.Atoi(a)
		if err != nil {
			log.Println(err)
		}

		c = append(c, m.GetName(ctx, b))
	} */
	return nil //заглушка
}

//метод для обновления возраста пользователя
func (m storage) UpdateUserAge(ctx context.Context, user_id int, newAge int) {
	/* a := strconv.Itoa(newAge)
	m.UserStorage[user_id].Age = a */
}

// метод удаления пользователя, по user_id удаляем самого пользователя и id из Friends друзей
func (m storage) Delete(ctx context.Context, user_id int) {
	// проходим по значению id друзей в самом пользователе
	/* for _, v := range m.UserStorage[user_id].Friends {
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

	delete(m.UserStorage, user_id) */
}
