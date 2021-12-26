package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation"

	"strconv"
	"strings"
)

var userID int
var UStorage UserStorage

const (
	userURL        = "/user" //?
	userIdURL      = "/{user_id}"
	friendsListURL = "/friends/{user_id}"
)

type handler struct {
}

type Handler interface {
	Register(r chi.Router)
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Register(r chi.Router) {
	r.Post("/create", h.NewUserProfile)     // 201, 4хх, Header Location: url
	r.Get(friendsListURL, h.GetFriendsList) // 200, 404, 500
	r.Post("/make_friends", h.MakeFriends)  // 200
	r.Delete(userURL, h.DeleteUserProfile)  //200/204
	r.Put(userIdURL, h.UpdateAge)           // 200, 4xx, 500
}

func jsonMarshToFile(u UserStorage) {
	u1 := UStorage
	j, err := json.Marshal(u1)
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile("userstr.txt", os.O_RDWR, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	_, err = f.Write(j)
	if err != nil {
		log.Println(err)
	}
}

//1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей.
// Пользователя необходимо сохранять в мапу. Пример запроса:
// POST /create HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"name":"some name","age":"24","friends":[]}
// Данный запрос должен возвращать ID пользователя и статус 201

func (h *handler) NewUserProfile(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		// инкремент ID
		userID++
		// переменная для вывода в теле ответа
		id := strconv.Itoa(userID)

		//массив для создания мапы
		friends = make([]string, 0)

		// читаем из тела запроса
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		// нарезаем в слайс прочитанный контент
		splittedContent := strings.Split(string(content), " ")
		k := newUser(splittedContent)
		// заполняем мапу
		u := UStorage.put(k, userID)
		j, err := json.Marshal(u)
		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.Create("userstr.txt")
		if err != nil {
			log.Println(err)
		}
		defer f.Close()

		_, err = f.Write(j)
		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// лог в консоль
		log.Println("userID is ", userID)
		//тело http ответа
		w.Write([]byte("User was created with ID = " + id))
		//log.Println()
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

// 2. Сделайте обработчик, который делает друзей из двух пользователей. Например, если мы создали двух пользователей
// и нам вернулись их ID, то в запросе мы можем указать ID пользователя, который инициировал запрос на дружбу, и ID пользователя,
//  который примет инициатора в друзья. Пример запроса:
// POST /make_friends HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"source_id":"1","target_id":"2"}
// Данный запрос должен возвращать статус 200 и сообщение «username_1 и username_2 теперь друзья».

func (h *handler) MakeFriends(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		splittedContent := strings.Split(string(content), " ")
		log.Println(splittedContent[1], splittedContent[3])
		a, b := splittedContent[1], splittedContent[3]
		u1, err := strconv.Atoi(a)
		if err != nil {
			log.Println(err)
		}
		u2, err := strconv.Atoi(b)
		if err != nil {
			log.Println(err)
		}

		u3 := UStorage.putFriend(u1, u2)
		u3 = UStorage.putFriend(u2, u1)

		jsonMarshToFile(u3)

		UStorage.get() // проверка записи в слайсы friends
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(UStorage.getName(u1) + " и " + UStorage.getName(u2) + " теперь друзья"))

		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

/* 3. Сделайте обработчик, который удаляет пользователя. Данный обработчик принимает ID пользователя и удаляет его из хранилища,
 а также стирает его из массива friends у всех его друзей. Пример запроса:
DELETE /user HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"target_id":"1"}
Данный запрос должен возвращать 200 и имя удалённого пользователя. */

func (h *handler) DeleteUserProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		splittedContent := strings.Split(string(content), " ")
		log.Println(splittedContent[1])
		user, err := strconv.Atoi(splittedContent[1])
		if err != nil {
			log.Println(err)
		}
		name := UStorage.getName(user)
		UStorage.delete(user)

		jsonMarshToFile(UStorage)

		UStorage.get() // проверка записи в слайсы friends

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte("пользователь " + name + " удалён"))

		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

/* 4. Сделайте обработчик, который возвращает всех друзей пользователя. Пример запроса:
GET /friends/user_id HTTP/1.1
Host: localhost:8080
Connection: close
После /friends/ указывается id пользователя, друзей которого мы хотим увидеть. */

func (h *handler) GetFriendsList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		user_idstr := chi.URLParam(r, "user_id")

		if errs := validation.Validate(user_idstr, validation.Required, validation.Match(regexp.MustCompile(`^\d+$`))); errs != nil {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(string(errs.Error())))
			return
		}
		user_idint, err := strconv.Atoi(user_idstr)
		if err != nil {
			log.Println(err)
		}

		a := strings.Join(UStorage.getFN(user_idint), ", ")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Connection:", "close")
		w.Write([]byte("В друзьях: " + a))
		return

	}
	w.WriteHeader(http.StatusBadRequest)
}

// 5. Сделайте обработчик, который обновляет возраст пользователя. Пример запроса:
// PUT /user_id HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"new age":"28"}
// Запрос должен возвращать 200 и сообщение «возраст пользователя успешно обновлён».

func (h *handler) UpdateAge(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		user_idstr := chi.URLParam(r, "user_id")

		if errs := validation.Validate(user_idstr, validation.Required, validation.Match(regexp.MustCompile(`^\d+$`))); errs != nil {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(string(errs.Error())))
			return
		}
		user_idint, err := strconv.Atoi(user_idstr)
		if err != nil {
			log.Println(err)
		}

		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		splittedContent := strings.Split(string(content), " ")
		log.Println(splittedContent[1])
		nAge, err := strconv.Atoi(splittedContent[1])
		if err != nil {
			log.Println(err)
		}

		UStorage.updateUserAge(user_idint, nAge)

		jsonMarshToFile(UStorage)

		UStorage.get() // проверка записи в слайсы friends

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte("возраст пользователя успешно обновлён: " + UStorage[user_idint].Age))
		return

	}
	w.WriteHeader(http.StatusBadRequest)

}
