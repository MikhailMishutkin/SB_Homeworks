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
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	//"github.com/go-chi/chi/v5"
)

func main() {
	//r := chi.NewRouter()
	lis, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server is running")
	con, err := lis.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(con)
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("line recieved:", line)

		upperLine := strings.ToUpper(string(line))
		if _, err := con.Write([]byte(upperLine)); err != nil {
			log.Fatal(err)
		}
	}
}

//func CreateUserHandler()
