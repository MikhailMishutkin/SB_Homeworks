package model

//Содержимое пакета:
//1. функция newUser(s []string) создания пользователя: принимает строку пользовательского ввода, возвращает структуру

//проверка соответствия структуры и интерфейса
//var _ User = &user{}

type User struct {
	Name    string   `json:"name,omitempty"`
	Age     string   `json:"age,omitempty"`
	Friends []string `json:"friends,omitempty"`
}

/* type User interface {
} */

//создание пользователя: принимает строку пользовательского ввода, возвращает структуру
func NewUser(s []string) User {

	u := User{
		Name:    s[1],
		Age:     s[3],
		Friends: nil,
	}
	return u
}
