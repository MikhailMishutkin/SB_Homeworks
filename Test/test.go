package main

import (
	"fmt"
	"regexp"

	//validation "github.com/go-ozzo/ozzo-validation/v4"
	//"github.com/go-ozzo/ozzo-validation/v4/is"
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ExchangeMessage struct {
	AuthorLastName  string `json:"last_name"`       // Имя автора
	AuthorFirstName string `json:"first_name"`      // Фамилия автора
	BookName        string `json:"book_name"`       // Название книги
	ISBN            string `json:"isbn"`            // ISBN книги
	YearPublishing  string `json:"year_publishing"` // Год публикации книги
	CategoryOffer   string `json:"category_offer"`  // ID категрорий книги
	CategoryWish    string `json:"category_wish"`   // ID желаемых категорий

	UserFirstName  string `json:"first_name_user"`  // Имя пользователя
	UserLastName   string `json:"last_name_user"`   // Фамилия пользователя
	UserSecondName string `json:"second_name_user"` // Отчество пользователя

	AddrIndex     string `json:"addr_index"`     // Индекс почтовый
	AddrCity      string `json:"addr_city"`      // Город
	AddrStreet    string `json:"addr_street"`    // Улица
	AddrHouse     string `json:"addr_house"`     // Дом
	AddrStructure string `json:"addr_structure"` // Строение, корпус
	AddrApart     string `json:"addr_appart"`    // Квартира
	IsDefault     bool   `json:"is_default"`     // Адрес по умолчанию
}

type UserProfile struct {
	IdUser     int    `db:"id_name" json:"IdUser"`
	LastName   string `db:"last_name" json:"LastName" validate:"required"`
	FirstName  string `db:"first_name" json:"FirstName" validate:"required"`
	SecondName string `db:"second_name" json:"SecondName"`
	Email      string `db:"e_mail" json:"Email" validate:"required"`
	Username   string `db:"user_name" json:"Username" validate:"required"`
	Password   string `db:"password_user" json:"Password" validate:"required"`
	RatingUser int
	//CreatedAt    time
	Enabled      bool
	IsStaff      bool
	IsSuperAdmin bool
	// }

	// type UserAddress struct {
	AddrIndex     string
	AddrCity      string
	AddrStreet    string
	AddrHouse     string
	AddrStructure string
	AddrAppart    string
}

// func (up UserProfile) Validate() error {
// 	return validation.ValidateStruct(&up,
// 		validation.Field(&up.LastName, validation.Required, validation.Length(2, 50), validation.Match(regexp.MustCompile("^[А-Яа-я]{2,}$"))),
// 		validation.Field(&up.FirstName, validation.Required, validation.Length(2, 25), validation.Match(regexp.MustCompile("^[А-Яа-я]{2,}$"))),
// 		validation.Field(&up.Email, validation.Required, is.Email),
// 		validation.Field(&up.Password, validation.Required, validation.Length(8, 50), is.Alphanumeric, validation.By(Check)),
// 		// 	)
// 		// }
// 		// func (up UserAddress) Validate() error {
// 		// 	return validation.ValidateStruct(&up,
// 		validation.Field(&up.AddrIndex, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{6}$"))),
// 		validation.Field(&up.AddrCity, validation.Required, validation.Length(2, 25), validation.Match(regexp.MustCompile("^[А-Я-а-я]{2,}$"))),
// 		validation.Field(&up.AddrStreet, validation.Required, validation.Length(3, 50), validation.Match(regexp.MustCompile("^[А-Яа-я0-9 -]{2,}$"))),
// 		validation.Field(&up.AddrHouse, validation.Required, validation.Length(1, 5), validation.Match(regexp.MustCompile("^[0-9А-Яа-я]{1,}$"))),
// 		validation.Field(&up.AddrAppart, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{1,4}$"))),
// 	)
// }

// AddRouteOrder Роутер вкладки создания и редактирования запросов.
// func AddRouteOrder(r *chi.Mux) {

// 	r.Route("/api/order", func(r chi.Router) {

// 		r.Post("/{UserID}", NewExchange)
// 		r.Get("/{UserID}", SendUserDataToFront)

// 		r.Get("/test", TestDBRequest)

// 	})

// 	r.Route("/api/authors", func(r chi.Router) {

// 		r.Get("/", SendAllAuthorsNames)
// 	})
// }

//SendAllAuthorsNames Функция Возвращает имя, фамилию и ID всех авторов в виде массива JSON
// func SendAllAuthorsNames(w http.ResponseWriter, _ *http.Request) {

// 	jtext, err := postgres.GetAllAuthors()

// 	if err != nil {
// 		returnErrorInfo(w, err)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	_, err = w.Write([]byte(jtext))

// 	if err != nil {
// 		return
// 	}

// }

// //SendUserDataToFront Передает в ответе почтовые данные пользователя.
// func SendUserDataToFront(w http.ResponseWriter, req *http.Request) {

// 	strUserID := chi.URLParam(req, "UserID")
// 	userID, err := strconv.Atoi(strUserID)

// 	if err != nil {
// 		returnErrorInfo(w, err)
// 		return
// 	}

// 	exists := false
// 	exists, err = postgres.UserIsExist(int64(userID))

// 	if err != nil {
// 		returnErrorInfo(w, err)
// 		return
// 	}

// 	if exists == false {
// 		w.WriteHeader(http.StatusBadRequest)
// 		_, err = w.Write([]byte("user does not exist"))
// 		return
// 	}

// 	jtext, jerr := postgres.GetUserPostalData(int64(userID))

// 	if jerr != nil {
// 		returnErrorInfo(w, jerr)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	_, err = w.Write([]byte(jtext))

// 	if err != nil {
// 		return
// 	}
// }

// NewExchange Создание новой заявки на обмен.
func NewExchange(w http.ResponseWriter, req *http.Request) {

	strUserID := chi.URLParam(req, "UserID")
	userID, err := strconv.Atoi(strUserID)
	fmt.Print(userID)
	if err != nil {
		fmt.Print(w, err)
		return
	}

	exists := false
	//exists, err = postgres.UserIsExist(int64(userID))

	if err != nil {
		fmt.Print(w, err)
		return
	}

	if exists == false {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("user does not exist"))
		return
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(req.Body)
	if err != nil {
		fmt.Print(w, err)
		return
	}

	jsont := buf.String()
	m := &ExchangeMessage{}
	err = json.Unmarshal([]byte(jsont), &m)

	if err != nil {
		fmt.Print(w, err)
		return
	}

	errs := m.Validate()
	if errs != nil {
		fmt.Print(w, errs)
		return
	}

	// var authorId, bookId, offerId, addressId, wishId int64
	// var catWishIds, catOfferIds []int64

	// //authorId, err = postgres.NewAuthor(m.AuthorLastName, m.AuthorFirstName)
	// if err != nil {
	// 	fmt.Print(w, err)
	// 	return
	// }
	// bookId, err = postgres.NewBookLiterary(authorId, m.BookName, "no note")
	// if err != nil {
	// 	fmt.Print(w, err)
	// 	return
	// }
	// offerId, err = postgres.NewOfferList(bookId, int64(userID), m.ISBN, m.YearPublishing)
	// if err != nil {
	// 	fmt.Print(w, err)
	// 	return
	// }
	// addressId, err = postgres.NewUserAddress(int64(userID), jsont)
	// if err != nil {
	// 	fmt.Print(w, err)
	// 	return
	// }
	// wishId, err = postgres.NewWishList(int64(userID), addressId)
	// if err != nil {
	// 	fmt.Print(w, err)
	// 	return
	// }
	// catOfferIds, err = postgres.NewCategoriesOffer(offerId, m.CategoryOffer)
	// if err != nil {
	// 	fmt.Print(w, err)
	// 	return
	// }
	// catWishIds, err = postgres.NewCategoriesWish(wishId, m.CategoryWish)
	// if err != nil {
	// 	fmt.Print(w, err)
	// 	return
	// }

	//err = postgres.UpdateUserName(int64(userID), m.UserFirstName, m.UserLastName, m.UserSecondName)

	mes := "OK"

	w.WriteHeader(http.StatusOK)

	_, err = w.Write([]byte(mes))

	if err != nil {
		return
	}

	//	fmt.Println(catOfferIds)
	//	fmt.Println(catWishIds)

}

// Validate Проверка JSON полученного в запросе на корректность данных.
func (em ExchangeMessage) Validate() error {
	return validation.ValidateStruct(&em,
		validation.Field(&em.UserFirstName, validation.Required, validation.RuneLength(2, 25), validation.Match(regexp.MustCompile("^[А-Яа-яЁё]+$"))),
		validation.Field(&em.UserLastName, validation.Required, validation.RuneLength(2, 50), validation.Match(regexp.MustCompile("^[А-Яа-яЁё]+$"))),
		validation.Field(&em.UserSecondName, validation.RuneLength(0, 25), validation.Match(regexp.MustCompile("^[А-Яа-яЁё]+$"))),
	)
}

func main() {
	var s ExchangeMessage = ExchangeMessage{
		AuthorLastName:  " ",
		AuthorFirstName: " ",
		BookName:        "",
		ISBN:            "",
		YearPublishing:  "",         // Год публикации книги
		CategoryOffer:   "",         // ID категрорий книги
		CategoryWish:    "14546546", // ID желаемых категорий

		UserFirstName:  "Алекса",      // Имя пользователя
		UserLastName:   "Алекса",      // Фамилия пользователя
		UserSecondName: "Собакович5!", // Отчество пользователя

		AddrIndex:     "",   // Индекс почтовый
		AddrCity:      "",   // Город
		AddrStreet:    "",   // Улица
		AddrHouse:     "",   // Дом
		AddrStructure: "",   // Строение, корпус
		AddrApart:     "",   // Квартира
		IsDefault:     true, // Адрес по умолчанию

		// LastName:   "Юзя",
		// FirstName:  "Ося",
		// SecondName: "",
		// Email:      "coolm16@gmail.com",
		// Username:   "SQence",
		// Password:   "Addlwerty3",
		// // }
		// // var t UserAddress = UserAddress{
		// AddrIndex:     "123456",
		// AddrCity:      "Москва",
		// AddrStreet:    "Стромынка-4",
		// AddrHouse:     "104",
		// AddrStructure: "1ф",
		// AddrAppart:    "124",
	}
	err := s.Validate()
	fmt.Println(err)
	// err = t.Validate()
	// fmt.Println(err)
	//	s.Validate()
	//fmt.Println(Check(s))

}

// 	if unicode.IsUpper(v) {
// 		a = true
// 	}
// 	if unicode.IsDigit(v) {
// 		b = true
// 	}
// 	if unicode.IsLower(v) {
// 		c = true
// 	}
// }
