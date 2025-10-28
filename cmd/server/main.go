package main

import (
	"log"
	"net/http"

	"github.com/triple-sun/metriccollector/internal/config"
	"github.com/triple-sun/metriccollector/internal/router"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func main() {
	var address string
	var cfg config.ServerConfig

	config.ParseServerFlags(&address)
	config.ParseServerEnv(&cfg, &address)

	log.Println(`Запускаю приложение...`)
	storage := storage.NewMemStorage()
	log.Println(`Создано хранилище`)
	r := router.SetupRoutes(storage)
	log.Println(`Настроен Router`)

	log.Printf(`Запускаю приложение на адресе %s...`, address)

	if err := http.ListenAndServe(address, r); err != nil {
		log.Fatal(`Ошибка запуска сервера`)
		panic(err)
	}
}
