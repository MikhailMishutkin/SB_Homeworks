package app

import (
	"net/http"
)

func proxy() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {

}
