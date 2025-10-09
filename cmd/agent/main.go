package main

import (
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"time"

	"github.com/triple-sun/metriccollector/internal/agent"
)

func main() {
	var stats runtime.MemStats

	pollInterval := 2
	pollCount := 0
	reportInterval := 10

	agent := agent.NewAgent()

	for {
		for {
			pollCount += 1
			runtime.ReadMemStats(&stats)
			log.Println(`Метрики прочитаны!`)

			if pollCount%(reportInterval/2) == 0 {
				agent.UpdateSingleMetric("counter", "PollCount", strconv.Itoa(pollCount))
				agent.UpdateSingleMetric("gauge", "RandomValue", strconv.Itoa(rand.Intn(1000)))
				agent.UpdateSingleMetric("gauge", "Alloc", strconv.FormatUint(stats.Alloc, 10))
				agent.UpdateSingleMetric("gauge", "BuckHashSys", strconv.FormatUint(stats.BuckHashSys, 10))
				agent.UpdateSingleMetric("gauge", "Frees", strconv.FormatUint(stats.Frees, 10))
				agent.UpdateSingleMetric("gauge", "GCCPUFraction", strconv.FormatFloat(stats.GCCPUFraction, 'f', 2, 64))
				agent.UpdateSingleMetric("gauge", "GCSys", strconv.FormatUint(stats.GCSys, 10))
				agent.UpdateSingleMetric("gauge", "HeapAlloc", strconv.FormatUint(stats.HeapAlloc, 10))
				agent.UpdateSingleMetric("gauge", "HeapIdle", strconv.FormatUint(stats.HeapIdle, 10))
				agent.UpdateSingleMetric("gauge", "HeapInuse", strconv.FormatUint(stats.HeapInuse, 10))
				agent.UpdateSingleMetric("gauge", "HeapObjects", strconv.FormatUint(stats.HeapObjects, 10))
				agent.UpdateSingleMetric("gauge", "HeapReleased", strconv.FormatUint(stats.HeapReleased, 10))
				agent.UpdateSingleMetric("gauge", "HeapSys", strconv.FormatUint(stats.HeapSys, 10))
				agent.UpdateSingleMetric("gauge", "LastGC", strconv.FormatUint(stats.LastGC, 10))
				agent.UpdateSingleMetric("gauge", "Lookups", strconv.FormatUint(stats.Lookups, 10))
				agent.UpdateSingleMetric("gauge", "MCacheInuse", strconv.FormatUint(stats.MCacheInuse, 10))
				agent.UpdateSingleMetric("gauge", "MCacheSys", strconv.FormatUint(stats.MCacheSys, 10))
				agent.UpdateSingleMetric("gauge", "MSpanInuse", strconv.FormatUint(stats.MSpanInuse, 10))
				agent.UpdateSingleMetric("gauge", "MSpanSys", strconv.FormatUint(stats.MSpanSys, 10))
				agent.UpdateSingleMetric("gauge", "Mallocs", strconv.FormatUint(stats.Mallocs, 10))
				agent.UpdateSingleMetric("gauge", "NextGC", strconv.FormatUint(stats.NextGC, 10))
				agent.UpdateSingleMetric("gauge", "NumForcedGC", strconv.FormatUint(uint64(stats.NumForcedGC), 10))
				agent.UpdateSingleMetric("gauge", "NumGC", strconv.FormatUint(uint64(stats.NumGC), 10))
				agent.UpdateSingleMetric("gauge", "OtherSys", strconv.FormatUint(stats.OtherSys, 10))
				agent.UpdateSingleMetric("gauge", "PauseTotalNs", strconv.FormatUint(stats.PauseTotalNs, 10))
				agent.UpdateSingleMetric("gauge", "StackInuse", strconv.FormatUint(stats.StackInuse, 10))
				agent.UpdateSingleMetric("gauge", "StackSys", strconv.FormatUint(stats.StackSys, 10))
				agent.UpdateSingleMetric("gauge", "Sys", strconv.FormatUint(stats.Sys, 10))
				agent.UpdateSingleMetric("gauge", "TotalAlloc", strconv.FormatUint(stats.TotalAlloc, 10))
				log.Println(`Метрики отправлены!`)
			}

			time.Sleep(time.Duration(pollInterval) * time.Second)
		}

	}

}
