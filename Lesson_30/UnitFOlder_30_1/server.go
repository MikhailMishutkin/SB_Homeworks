package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	//"github.com/go-chi/chi/v5"
)

func main() {
	//r := chi.NewRouter()
	lis, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server is running")
	con, err := lis.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		line, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("line recieved:", string(line))

		upperLine := strings.ToUpper(string(line))
		if _, err := con.Write([]byte(upperLine)); err != nil {
			log.Fatal(err)
		}
	}
}
