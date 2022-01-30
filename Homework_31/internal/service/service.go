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
		GetAll(ctx context.Context)
		GetName(ctx context.Context, user_id int) string
		PutInto(ctx context.Context, strt model.User, user_id int)
		PutFriends(ctx context.Context, user_id1, user_id2 int)
		GetFNs(ctx context.Context, user_id int) []string
		UpdateUserAge(ctx context.Context, user_id int, newAge int)
		Delete(ctx context.Context, user_id int)
	}













// взял из handler MakeFriends task 30: записывает в структуру список друзей обоим друзьям
u3 := UStorage.putFriend(u1, u2)
u3 = UStorage.putFriend(u2, u1)

//NewUser(splittedContent)




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
