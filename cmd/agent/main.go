package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"

	"resty.dev/v3"

	"github.com/triple-sun/metriccollector/internal/agent"
	"github.com/triple-sun/metriccollector/internal/config"
)

func main() {
	var stats runtime.MemStats
	var cfg config.AgentConfig

	var address string
	var pInterval int
	var rInterval int

	// Аргументы и переменные окружения
	config.ParseAgentFlags(&address, &pInterval, &rInterval)
	config.ParseAgentEnv(&cfg, &address, &pInterval, &rInterval)

	client := resty.New().SetBaseURL(fmt.Sprintf("http://%s", address))

	pCount := 0
	pTicker := time.NewTicker(time.Duration(pInterval) * time.Second)
	defer pTicker.Stop()

	log.Printf(`Запущен агент отправки метрик. Адрес: %s, интервал сбора: %dс, интервал отправки: %dс`, address, pInterval, rInterval)

	for range pTicker.C {
		pCount += 1
		runtime.ReadMemStats(&stats)
		log.Println(`Метрики прочитаны!`)

		if pCount%(rInterval/pInterval) == 0 {
			for _, params := range []agent.MetricUpdateParams{
				{ID: "PollCount", MType: "counter", Delta: int64(pCount)},
				{ID: "RandomValue", MType: "gauge", Value: float64(rand.Intn(1000))},
				{ID: "Alloc", MType: "gauge", Value: float64(stats.Alloc)},
				{ID: "BuckHashSys", MType: "gauge", Value: float64(stats.BuckHashSys)},
				{ID: "Frees", MType: "gauge", Value: float64(stats.Frees)},
				{ID: "GCCPUFraction", MType: "gauge", Value: stats.GCCPUFraction},
				{ID: "GCSys", MType: "gauge", Value: float64(stats.GCSys)},
				{ID: "HeapAlloc", MType: "gauge", Value: float64(stats.HeapAlloc)},
				{ID: "HeapIdle", MType: "gauge", Value: float64(stats.HeapIdle)},
				{ID: "HeapInuse", MType: "gauge", Value: float64(stats.HeapInuse)},
				{ID: "HeapObjects", MType: "gauge", Value: float64(stats.HeapObjects)},
				{ID: "HeapReleased", MType: "gauge", Value: float64(stats.HeapReleased)},
				{ID: "HeapSys", MType: "gauge", Value: float64(stats.HeapSys)},
				{ID: "LastGC", MType: "gauge", Value: float64(stats.LastGC)},
				{ID: "Lookups", MType: "gauge", Value: float64(stats.Lookups)},
				{ID: "MCacheInuse", MType: "gauge", Value: float64(stats.MCacheInuse)},
				{ID: "MCacheSys", MType: "gauge", Value: float64(stats.MCacheSys)},
				{ID: "MSpanInuse", MType: "gauge", Value: float64(stats.MSpanInuse)},
				{ID: "MSpanSys", MType: "gauge", Value: float64(stats.MSpanSys)},
				{ID: "Mallocs", MType: "gauge", Value: float64(stats.Mallocs)},
				{ID: "NextGC", MType: "gauge", Value: float64(stats.NextGC)},
				{ID: "NumForcedGC", MType: "gauge", Value: float64(stats.NumForcedGC)},
				{ID: "NumGC", MType: "gauge", Value: float64(stats.NumGC)},
				{ID: "OtherSys", MType: "gauge", Value: float64(stats.OtherSys)},
				{ID: "PauseTotalNs", MType: "gauge", Value: float64(stats.PauseTotalNs)},
				{ID: "StackInuse", MType: "gauge", Value: float64(stats.StackInuse)},
				{ID: "StackSys", MType: "gauge", Value: float64(stats.StackSys)},
				{ID: "Sys", MType: "gauge", Value: float64(stats.Sys)},
				{ID: "TotalAlloc", MType: "gauge", Value: float64(stats.TotalAlloc)},
			} {
				agent.UpdateMetric(client, params)
			}

			log.Println(`Метрики отправлены!`)
		}
	}
}
