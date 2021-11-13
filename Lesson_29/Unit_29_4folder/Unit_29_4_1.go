//Мьютекс для избежания гонки горутин

package main

import "sync"

type server struct {
	status     string
	sync.Mutex // структура мьютекс создан для блокировки взаимного исключения горутин
}

func main() {

	s := server{}

	for i := 0; i < 1000; i++ {
		go s.Alive()
		go s.Down()
	}
}

func (s *server) Alive() {
	s.Lock() // метод мьютекс для закрытия пока исполняется другая горутина
	s.status = "Alive"
	s.Unlock()
}

func (s *server) Down() {
	s.Lock()
	s.status = "Down"
	s.Unlock()
}
