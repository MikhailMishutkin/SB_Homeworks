// Напишите HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом:

//1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей.
// Пользователя необходимо сохранять в мапу. Пример запроса:

// POST /create HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"name":"some name","age":"24","friends":[]}
// Данный запрос должен возвращать ID пользователя и статус 201
// далее в pkg/handlers.go

package main

import (
	"log"
	"net"
	"net/http"

	"task30/pkg"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	handler := pkg.NewHandler()
	handler.Register(r)

	// открываем мапу для записи
	pkg.UStorage = make(pkg.UserStorage)

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
