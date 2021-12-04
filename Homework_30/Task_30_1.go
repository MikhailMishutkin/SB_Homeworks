// Напишите HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом:

//1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей.
// Пользователя необходимо сохранять в мапу. Пример запроса:

// POST /create HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"name":"some name","age":"24","friends":[]}
// Данный запрос должен возвращать ID пользователя и статус 201

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

type User struct {
	Name    string
	Age     string
	Friends []string
}

type UserStorage map[string]*User

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
	//var newUser User

	if r.Method == "POST" {
		userID++
		id := strconv.Itoa(userID)
		m := make(UserStorage)
		friends := make([]string, 2)
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		splittedContent := strings.Split(string(content), " ")
		//m[splittedContent[0]] = &User{}
		m[string(userID)] = &User{splittedContent[0], splittedContent[1], friends}
		w.WriteHeader(http.StatusCreated)
		log.Println("userID is ", userID)
		w.Write([]byte("User was created with ID = " + id))

		return
	}

	w.WriteHeader(http.StatusBadRequest)
}
func (h *handler) GetFriendsList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is list of users"))
}
func (h *handler) MakeFriends(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is making friends"))
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

	server := &http.Server{
		Handler: r,
	}
	log.Fatalln(server.Serve(lis))
}
