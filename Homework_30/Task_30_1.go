// Напишите HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом:

//1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей.
// Пользователя необходимо сохранять в мапу. Пример запроса:

// POST /create HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"name":"some name","age":"24","friends":[]}
// Данный запрос должен возвращать ID пользователя и статус 201

// 2. Сделайте обработчик, который делает друзей из двух пользователей. Например, если мы создали двух пользователей
// и нам вернулись их ID, то в запросе мы можем указать ID пользователя, который инициировал запрос на дружбу, и ID пользователя,
//  который примет инициатора в друзья. Пример запроса:

// POST /make_friends HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"source_id":"1","target_id":"2"}
// Данный запрос должен возвращать статус 200 и сообщение «username_1 и username_2 теперь друзья».

package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

const (
	usersURL       = "/users" //?
	userURL        = "/users/:id"
	friendsListURL = "/friends/user_id"
)

var userID int
var usersStorage structs.UserStorage
var friends []string

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
		usersStorage.put(k, userID)
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

func (h *handler) MakeFriends(w http.ResponseWriter, r *http.Request) {

	/* if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		usersStorage[]
		splittedContent := strings.Split(string(content), " ")

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User was created  " + splittedContent[0]))

		return
	}

	w.WriteHeader(http.StatusBadRequest) */
}

func (h *handler) GetFriendsList(w http.ResponseWriter, r *http.Request) {
	usersStorage.get()
	w.Write([]byte("this is list of users"))
}

func (h *handler) DeleteUserProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is delete user"))
}
func (h *handler) UpdateAge(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is update user age"))
}

/* func IndexHandler(w http.ResponseWriter, r *http.Request, params chi.RouteParams) {
	// name := params.ByName("name")
	// w.Write([]byte(fmt.Sprintf("Hello %s", name)))
} */

func main() {
	r := chi.NewRouter()
	handler := NewHandler()
	handler.Register(r)

	// открываем мапу для записи
	usersStorage = make(UserStorage)

	lis, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	server := &http.Server{
		Handler: r,
	}
	log.Fatalln(server.Serve(lis))

}
