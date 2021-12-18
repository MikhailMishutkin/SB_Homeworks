package pkg

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"strconv"
	"strings"
)

var userID int
var UStorage UserStorage

const (
	usersURL       = "/users" //?
	userURL        = "/users/:id"
	friendsListURL = "/friends/user_id"
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
	r.Delete(usersURL, h.DeleteUserProfile) //200/204
	r.Put(userURL, h.UpdateAge)             // 200, 4xx, 500
}

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
		UStorage.put(k, userID)
		w.WriteHeader(http.StatusCreated)
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

		UStorage.putFriend(u1, u2)
		UStorage.putFriend(u2, u1)
		UStorage.get() // проверка записи в слайсы friends
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(UStorage.getName(u1) + " и " + UStorage.getName(u2) + " теперь друзья"))

		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (h *handler) GetFriendsList(w http.ResponseWriter, r *http.Request) {
	UStorage.get()
	w.Write([]byte("this is list of users"))
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
		UStorage.get() // проверка записи в слайсы friends

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("пользователь " + name + " удалён"))

		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (h *handler) UpdateAge(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is update user age"))
}
