package model

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

/* type User interface {
} */

//создание пользователя: принимает строку пользовательского ввода, возвращает структуру
func NewUser(s []string) User {
	idCounter++
	u := User{
		Id:      idCounter,
		Name:    s[1],
		Age:     s[3],
		Friends: nil,
	}
	return u
}
