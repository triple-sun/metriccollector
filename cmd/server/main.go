package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/triple-sun/metriccollector/internal/router"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func main() {
	var addr string

	flag.StringVar(&addr, "a", "localhost:8080", "Адрес в формате host:port")
	flag.Parse()

	log.Println(`Запускаю приложение...`)
	storage := storage.NewMemStorage()
	log.Println(`Создано хранилище`)
	r := router.Setup(storage)
	log.Println(`Создан Router`)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(`Ошибка запуска сервера`)
		panic(err)
	}

	fmt.Printf(`Сервер запущен по адресу %s!`, addr)
}
