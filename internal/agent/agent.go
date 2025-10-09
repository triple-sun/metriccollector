package agent

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
)

type Agent struct{}

type IAgent interface {
	UpdateSingleMetric(mtype string, name string, value string)
	UpdateMetrics(stats *runtime.MemStats)
}

func NewAgent() *Agent {
	return &Agent{}
}

func (a *Agent) UpdateSingleMetric(mtype string, name string, value string) {
	url := fmt.Sprintf(`http://localhost:8080/update/%s/%s/%s`, mtype, name, value)

	log.Printf(`Обновляю метрику %s типа %s: %s`, name, mtype, value)
	res, err := http.Post(url, "application/json", nil)

	if err != nil {
		log.Printf(`Не удалось обновить метрику %s: %s`, name, err.Error())
	}

	if res.StatusCode != 200 {
		log.Printf(`Не удалось обновить метрику: [%d]`, res.StatusCode)
	}
}

func (a *Agent) UpdateMetrics(stats *runtime.MemStats, pollCount int) {
	a.UpdateSingleMetric("counter", "PollCount", strconv.Itoa(pollCount))
	a.UpdateSingleMetric("gauge", "RandomValue", strconv.Itoa(rand.Intn(1000)))
	a.UpdateSingleMetric("gauge", "Alloc", strconv.FormatUint(stats.Alloc, 10))
	a.UpdateSingleMetric("gauge", "BuckHashSys", strconv.FormatUint(stats.BuckHashSys, 10))
	a.UpdateSingleMetric("gauge", "Frees", strconv.FormatUint(stats.Frees, 10))
	a.UpdateSingleMetric("gauge", "GCCPUFraction", strconv.FormatFloat(stats.GCCPUFraction, 'f', 2, 64))
	a.UpdateSingleMetric("gauge", "GCSys", strconv.FormatUint(stats.GCSys, 10))
	a.UpdateSingleMetric("gauge", "HeapAlloc", strconv.FormatUint(stats.HeapAlloc, 10))
	a.UpdateSingleMetric("gauge", "HeapIdle", strconv.FormatUint(stats.HeapIdle, 10))
	a.UpdateSingleMetric("gauge", "HeapInuse", strconv.FormatUint(stats.HeapInuse, 10))
	a.UpdateSingleMetric("gauge", "HeapObjects", strconv.FormatUint(stats.HeapObjects, 10))
	a.UpdateSingleMetric("gauge", "HeapReleased", strconv.FormatUint(stats.HeapReleased, 10))
	a.UpdateSingleMetric("gauge", "HeapSys", strconv.FormatUint(stats.HeapSys, 10))
	a.UpdateSingleMetric("gauge", "LastGC", strconv.FormatUint(stats.LastGC, 10))
	a.UpdateSingleMetric("gauge", "Lookups", strconv.FormatUint(stats.Lookups, 10))
	a.UpdateSingleMetric("gauge", "MCacheInuse", strconv.FormatUint(stats.MCacheInuse, 10))
	a.UpdateSingleMetric("gauge", "MCacheSys", strconv.FormatUint(stats.MCacheSys, 10))
	a.UpdateSingleMetric("gauge", "MSpanInuse", strconv.FormatUint(stats.MSpanInuse, 10))
	a.UpdateSingleMetric("gauge", "MSpanSys", strconv.FormatUint(stats.MSpanSys, 10))
	a.UpdateSingleMetric("gauge", "Mallocs", strconv.FormatUint(stats.Mallocs, 10))
	a.UpdateSingleMetric("gauge", "NextGC", strconv.FormatUint(stats.NextGC, 10))
	a.UpdateSingleMetric("gauge", "NumForcedGC", strconv.FormatUint(uint64(stats.NumForcedGC), 10))
	a.UpdateSingleMetric("gauge", "NumGC", strconv.FormatUint(uint64(stats.NumGC), 10))
	a.UpdateSingleMetric("gauge", "OtherSys", strconv.FormatUint(stats.OtherSys, 10))
	a.UpdateSingleMetric("gauge", "PauseTotalNs", strconv.FormatUint(stats.PauseTotalNs, 10))
	a.UpdateSingleMetric("gauge", "StackInuse", strconv.FormatUint(stats.StackInuse, 10))
	a.UpdateSingleMetric("gauge", "StackSys", strconv.FormatUint(stats.StackSys, 10))
	a.UpdateSingleMetric("gauge", "Sys", strconv.FormatUint(stats.Sys, 10))
	a.UpdateSingleMetric("gauge", "TotalAlloc", strconv.FormatUint(stats.TotalAlloc, 10))
}
