package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/triple-sun/metriccollector/internal/handler"
	server "github.com/triple-sun/metriccollector/internal/storage"
)

func main() {
	log.Println(`Запускаю приложение...`)
	storage := server.NewMemStorage()
	log.Println(`Создано хранилище`)

	r := chi.NewRouter()
	log.Println(`Создан Router`)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post(`/update/{mtype}/{mname}/{mvalue}`, handler.MetricHandler(storage))

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(`Ошибка запуска сервера`)
		panic(err)
	}

	fmt.Println(`Сервер запущен!`)
}
