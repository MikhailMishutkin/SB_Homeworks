package handlers

import (
	"io/ioutil"
	"net/http"
	"task31/internal/model"
	"task31/internal/storage"

	"github.com/go-chi/chi/v5"
)

const (
	userURL        = "/user" //?
	userIdURL      = "/{user_id}"
	friendsListURL = "/friends/{user_id}"
)

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
	r.Delete(userURL, h.DeleteUserProfile)  //200/204
	r.Put(userIdURL, h.UpdateAge)           // 200, 4xx, 500
}

//1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей.
// Пользователя необходимо сохранять в мапу. Пример запроса:
// POST /create HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"name":"some name","age":"24","friends":[]}
// Данный запрос должен возвращать ID пользователя и статус 201

func (h *handler) NewUserProfile(w http.ResponseWriter, r *http.Request) {

	// читаем из тела запроса
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	data := model.NewUser(content)
	DataDB := storage.InsertToDB(data)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//тело http ответа
	w.Write([]byte("User was created with ID = " + model.IdCounter))

	return

}

// 2. Сделайте обработчик, который делает друзей из двух пользователей. Например, если мы создали двух пользователей
// и нам вернулись их ID, то в запросе мы можем указать ID пользователя, который инициировал запрос на дружбу, и ID пользователя,
//  который примет инициатора в друзья. Пример запроса:
// POST /make_friends HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"source_id":"1","target_id":"2"}
// Данный запрос должен возвращать статус 200 и сообщение «username_1 и username_2 теперь друзья».

func (h *handler) MakeFriends(w http.ResponseWriter, r *http.Request) {
	/*
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		splittedContent := strings.Split(string(content), " ")
		log.Println(splittedContent[1], splittedContent[3])
		a, b := splittedContent[1], splittedContent[3]
		u1, err := strconv.Atoi(a)
		if err != nil {
			log.Println(err)
		}
		u2, err := strconv.Atoi(b)
		if err != nil {
			log.Println(err)
		}

		jsonMarshToFile(u3)

		UStorage.get() // проверка записи в слайсы friends
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(UStorage.getName(u1) + " и " + UStorage.getName(u2) + " теперь друзья"))
	*/
}

/* 3. Сделайте обработчик, который удаляет пользователя. Данный обработчик принимает ID пользователя и удаляет его из хранилища,
 а также стирает его из массива friends у всех его друзей. Пример запроса:
DELETE /user HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"target_id":"1"}
Данный запрос должен возвращать 200 и имя удалённого пользователя. */

func (h *handler) DeleteUserProfile(w http.ResponseWriter, r *http.Request) {
	/* content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	splittedContent := strings.Split(string(content), " ")
	log.Println(splittedContent[1])
	user, err := strconv.Atoi(splittedContent[1])
	if err != nil {
		log.Println(err)
	}
	name := UStorage.getName(user)
	UStorage.delete(user)

	jsonMarshToFile(UStorage)

	UStorage.get() // проверка записи в слайсы friends

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("пользователь " + name + " удалён"))
	*/
}

/* 4. Сделайте обработчик, который возвращает всех друзей пользователя. Пример запроса:
GET /friends/user_id HTTP/1.1
Host: localhost:8080
Connection: close
После /friends/ указывается id пользователя, друзей которого мы хотим увидеть. */

func (h *handler) GetFriendsList(w http.ResponseWriter, r *http.Request) {
	/*
		user_idStr := chi.URLParam(r, "user_id")

		if errs := validation.Validate(user_idStr, validation.Required, validation.Match(regexp.MustCompile(`^\d+$`))); errs != nil {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(string(errs.Error())))
			return
		}
		user_idInt, err := strconv.Atoi(user_idStr)
		if err != nil {
			log.Println(err)
		}

		a := strings.Join(UStorage.getFN(user_idInt), ", ")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Connection:", "close")
		w.Write([]byte("В друзьях: " + a))
		return
	*/
}

// 5. Сделайте обработчик, который обновляет возраст пользователя. Пример запроса:
// PUT /user_id HTTP/1.1
// Content-Type: application/json; charset=utf-8
// Host: localhost:8080
// {"new age":"28"}
// Запрос должен возвращать 200 и сообщение «возраст пользователя успешно обновлён».

func (h *handler) UpdateAge(w http.ResponseWriter, r *http.Request) {
	/*
		user_idstr := chi.URLParam(r, "user_id")

		if errs := validation.Validate(user_idstr, validation.Required, validation.Match(regexp.MustCompile(`^\d+$`))); errs != nil {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(string(errs.Error())))
			return
		}
		user_idint, err := strconv.Atoi(user_idstr)
		if err != nil {
			log.Println(err)
		}

		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		splittedContent := strings.Split(string(content), " ")
		log.Println(splittedContent[1])
		nAge, err := strconv.Atoi(splittedContent[1])
		if err != nil {
			log.Println(err)
		}

		UStorage.updateUserAge(user_idint, nAge)

		jsonMarshToFile(UStorage)

		UStorage.get() // проверка записи в слайсы friends

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte("возраст пользователя успешно обновлён: " + UStorage[user_idint].Age))
	*/
}
