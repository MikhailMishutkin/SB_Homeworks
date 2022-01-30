package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task31/internal/storage"
)

var _ Service = &service{}

type service struct {
}

type Service interface {
	IncrId()
	Unmarshal()
	JsonMarshToFile(u storage.UserStorage)
	JsonMarshToDB(u storage.UserStorage)
	SplitContent(c []byte)
	Atoi(a []string) (int, error)
}

// заполняем мапу!!!!!
u := UStorage.put(k, userID)

// взял из handler MakeFriends task 30: записывает в структуру список друзей обоим друзьям
u3 := UStorage.putFriend(u1, u2)
u3 = UStorage.putFriend(u2, u1)

//NewUser(splittedContent)

/* func (u *model.User) NewUser(ctx context.Context, s []string) model.User {
	// новый юзер структура
	//var u *model.User
	u = &model.User{s[1], s[3], storage.Friends}
	return u
} */

func (s service) JsonMarshToFile(u storage.UserStorage) {
	u1 := UStorage
	j, err := json.Marshal(u1)
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile("userstr.txt", os.O_RDWR, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	_, err = f.Write(j)
	if err != nil {
		log.Println(err)
	}
}
