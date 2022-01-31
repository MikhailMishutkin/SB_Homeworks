package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//Содержимое пакета:
//1. функция newUser(s []string) создания пользователя: принимает строку пользовательского ввода, возвращает структуру

//проверка соответствия структуры и интерфейса
//var _ User = &user{}

type User struct {
	Id      int      `json:"user_id,omitempty"`
	Name    string   `json:"name"`
	Age     string   `json:"age"`
	Friends []string `json:"friends"`
}

var idCounter int

func main() {
	for {
		var name []byte
		fmt.Fscan(os.Stdin, &name)
		a := NewUser(name)
		fmt.Println(a)
	}
}

//создание пользователя: принимает строку пользовательского ввода, возвращает структуру
func NewUser(s []byte) User {
	idCounter++
	var u User
	err := json.Unmarshal(s, &u)
	if err != nil {
		log.Fatal(err)
	}
	u.Id = idCounter

	return u
}
