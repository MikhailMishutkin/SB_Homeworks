package app

import (
	"log"
	"net"
	"net/http"
	"task31/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func StartApp() {
	r := chi.NewRouter()
	handler := handlers.NewHandler()
	handler.Register(r)

	var count int8
	var lis net.Listener
	var err error

	if count == 0 {
		lis, err = net.Listen("tcp4", "localhost:8080")
		if err != nil {
			log.Fatal(err)
		}
		defer lis.Close()
		count++
	} else {
		lis, err = net.Listen("tcp4", "localhost:8081")
		if err != nil {
			log.Fatal(err)
		}
		defer lis.Close()
		count--
	}

	server := &http.Server{
		Handler: r,
	}

	log.Fatalln(server.Serve(lis))

}
