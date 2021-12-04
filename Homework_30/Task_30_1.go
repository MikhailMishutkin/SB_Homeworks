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
	"fmt"
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
var usersStorage UserStorage
var friends []string

type User struct {
	Name    string
	Age     string
	Friends []string
}

type UserStorage map[int]*User

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
func (h *handler) NewUserProfile(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		// инкремент ID
		userID++
		// переменная для вывода в теле ответа
		id := strconv.Itoa(userID)
		// мапа для записи нового пользователя
		m := make(UserStorage)
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
		m.put(k, userID)
		w.WriteHeader(http.StatusCreated)
		// лог в консоль
		log.Println("userID is ", userID)
		//тело http ответа
		w.Write([]byte("User was created with ID = " + id))

		usersStorage = m

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
	//r.Get("/", IndexHandler) //func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hi"))
	// })

	lis, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	server := &http.Server{
		Handler: r,
	}
	log.Fatalln(server.Serve(lis))

	usersStorage.get()

}
