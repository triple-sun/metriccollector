package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/triple-sun/metriccollector/internal/router"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func main() {
	log.Println(`Запускаю приложение...`)
	storage := storage.NewMemStorage()
	log.Println(`Создано хранилище`)
	r := router.Setup(storage)
	log.Println(`Создан Router`)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(`Ошибка запуска сервера`)
		panic(err)
	}

	fmt.Println(`Сервер запущен!`)
}
