package agent

import (
	"log"

	"resty.dev/v3"

	"github.com/triple-sun/metriccollector/internal/responses"
)

type MetricsParams struct {
	MType string
	Value string
}

func UpdateSingleMetric(client *resty.Client, mtype string, mname string, mvalue string) {
	var response responses.MetricUpdateResponse
	var responseErr responses.ErrorResponse

	log.Printf(`Обновляю метрику %s типа %s: %s`, mname, mtype, mvalue)

	req := resty.New().R().SetPathParams(map[string]string{
		"mtype": mtype, "mvalue": mvalue, "mname": mname,
	}).SetResult(&response).SetError(&responseErr)

	res, err := req.Post(client.BaseURL() + "/update/{mtype}/{mname}/{mvalue}")

	if err != nil {
		log.Printf("ошибка обновления метрики %s: %s", mname, err.Error())
		return
	}

	log.Printf("метрика %s отправлена: %v", mname, res)
}
