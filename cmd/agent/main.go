package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"time"

	"resty.dev/v3"

	"github.com/triple-sun/metriccollector/internal/agent"
)

func main() {
	var stats runtime.MemStats

	var addr string
	var pInterval int
	var rInterval int

	pCount := 0

	flag.StringVar(&addr, "a", "localhost:8080", "Адрес в формате host:port")
	flag.IntVar(&pInterval, "p", 2, "Интервал сбора метрик в секундах")
	flag.IntVar(&rInterval, "r", 10, "Интервал отправки метрик в секундах")

	flag.Parse()

	client := resty.New().SetBaseURL(fmt.Sprintf("http://%s", addr))

	log.Printf(`Запущен агент отправки метрик. Адрес: %s, интервал сбора: %dс, интервал отправки: %dс`, addr, pInterval, rInterval)

	for {
		for {
			pCount += 1
			runtime.ReadMemStats(&stats)
			log.Println(`Метрики прочитаны!`)

			if pCount%(rInterval/2) == 0 {
				for mname, params := range map[string]agent.MetricsParams{
					"PollCount":     {MType: "counter", Value: strconv.Itoa(pCount)},
					"RandomValue":   {MType: "gauge", Value: strconv.Itoa(rand.Intn(1000))},
					"Alloc":         {MType: "gauge", Value: strconv.FormatUint(stats.Alloc, 10)},
					"BuckHashSys":   {MType: "gauge", Value: strconv.FormatUint(stats.BuckHashSys, 10)},
					"Frees":         {MType: "gauge", Value: strconv.FormatUint(stats.Frees, 10)},
					"GCCPUFraction": {MType: "gauge", Value: strconv.FormatFloat(stats.GCCPUFraction, 'f', 2, 64)},
					"GCSys":         {MType: "gauge", Value: strconv.FormatUint(stats.GCSys, 10)},
					"HeapAlloc":     {MType: "gauge", Value: strconv.FormatUint(stats.HeapAlloc, 10)},
					"HeapIdle":      {MType: "gauge", Value: strconv.FormatUint(stats.HeapIdle, 10)},
					"HeapInuse":     {MType: "gauge", Value: strconv.FormatUint(stats.HeapInuse, 10)},
					"HeapObjects":   {MType: "gauge", Value: strconv.FormatUint(stats.HeapObjects, 10)},
					"HeapReleased":  {MType: "gauge", Value: strconv.FormatUint(stats.HeapReleased, 10)},
					"HeapSys":       {MType: "gauge", Value: strconv.FormatUint(stats.HeapSys, 10)},
					"LastGC":        {MType: "gauge", Value: strconv.FormatUint(stats.LastGC, 10)},
					"Lookups":       {MType: "gauge", Value: strconv.FormatUint(stats.Lookups, 10)},
					"MCacheInuse":   {MType: "gauge", Value: strconv.FormatUint(stats.MCacheInuse, 10)},
					"MCacheSys":     {MType: "gauge", Value: strconv.FormatUint(stats.MCacheSys, 10)},
					"MSpanInuse":    {MType: "gauge", Value: strconv.FormatUint(stats.MSpanInuse, 10)},
					"MSpanSys":      {MType: "gauge", Value: strconv.FormatUint(stats.MSpanSys, 10)},
					"Mallocs":       {MType: "gauge", Value: strconv.FormatUint(stats.Mallocs, 10)},
					"NextGC":        {MType: "gauge", Value: strconv.FormatUint(stats.NextGC, 10)},
					"NumForcedGC":   {MType: "gauge", Value: strconv.FormatUint(uint64(stats.NumForcedGC), 10)},
					"NumGC":         {MType: "gauge", Value: strconv.FormatUint(uint64(stats.NumGC), 10)},
					"OtherSys":      {MType: "gauge", Value: strconv.FormatUint(stats.OtherSys, 10)},
					"PauseTotalNs":  {MType: "gauge", Value: strconv.FormatUint(stats.PauseTotalNs, 10)},
					"StackInuse":    {MType: "gauge", Value: strconv.FormatUint(stats.StackInuse, 10)},
					"StackSys":      {MType: "gauge", Value: strconv.FormatUint(stats.StackSys, 10)},
					"Sys":           {MType: "gauge", Value: strconv.FormatUint(stats.Sys, 10)},
					"TotalAlloc":    {MType: "gauge", Value: strconv.FormatUint(stats.TotalAlloc, 10)},
				} {
					agent.UpdateSingleMetric(client, params.MType, mname, params.Value)
				}

				log.Println(`Метрики отправлены!`)
			}

			time.Sleep(time.Duration(pInterval) * time.Second)
		}

	}

}
