package main

import "fmt"

type OfferBookInfo struct {
	OfferID         int    `db:"id_offer_list"`
	BookName        string `db:"book_name"`
	Note            string `db:"note"`
	AuthorFirstName string `db:"first_name"`
	AuthorLastName  string `db:"last_name"`
	ISBN            string `db:"isbn"`
	YearPublishing  int32  `db:"year_publishing"`
	CategoryID      int    `db:"id_category"`
}

func main() {
	var f []OfferBookInfo
	f1 := OfferBookInfo{
		OfferID:         1,
		BookName:        "Name",
		Note:            "Note",
		AuthorFirstName: "AfName",
		AuthorLastName:  "AlName",
		ISBN:            "12354864215",
		YearPublishing:  1980,
		CategoryID:      36,
	}
	f = append(f, f1)
	f = append(f, f1)

	fmt.Println(f)
}

func main() {
	result, err := db.ExecContext(ctx,
		"INSERT INTO users (name, age) VALUES ($1, $2)",
		"gopher",
		27,
	)
}
