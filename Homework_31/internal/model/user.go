package model

import (
	"encoding/json"
	"log"
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

var IdCounter int

//создание пользователя: принимает строку пользовательского ввода, возвращает структуру
func NewUser(s []byte) User {
	IdCounter++
	var u User
	err := json.Unmarshal(s, &u)
	if err != nil {
		log.Fatal(err)
	}
	u.Id = IdCounter

	return u
}
