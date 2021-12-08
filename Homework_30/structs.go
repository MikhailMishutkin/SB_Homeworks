package structs

type User struct {
	Name    string
	Age     string
	Friends []string
}

type UserStorage map[int]*User

type handler struct {
}
