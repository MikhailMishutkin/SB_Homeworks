package main

import "net/http"

func main() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {

}
