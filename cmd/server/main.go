package main

import (
	"fmt"
	"net/http"

	"github.com/triple-sun/metriccollector/internal/handler"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func main() {
	fmt.Println(`Запускаю приложение...`)
	mux := http.NewServeMux()
	fmt.Println(`Создан Mux`)
	storage := server.NewMemStorage()
	fmt.Println(`Создано хранилище`)

	mux.HandleFunc(`/update/{metricType}/{metricName}/{metricValue}`, func(w http.ResponseWriter, r *http.Request) {
		handler.HandleMetric(w, r, storage)
	})

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println(`Ошибка запуска сервера`)

		panic(err)
	}

	fmt.Println(`Сервер запущен!`)
}
