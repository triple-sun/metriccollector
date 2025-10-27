package agent

import (
	"log"

	"resty.dev/v3"

	"github.com/triple-sun/metriccollector/internal/model"
	"github.com/triple-sun/metriccollector/internal/responses"
)

type MetricUpdateParams struct {
	ID    string
	MType string
	Value float64
	Delta int64
}

func UpdateMetric(client *resty.Client, params MetricUpdateParams) {
	var response responses.MetricUpdateResponse
	var responseErr responses.ErrorResponse

	log.Printf(`Обновляю метрику %s типа %s`, params.ID, params.MType)

	var body model.Metrics

	switch params.MType {
	case model.Counter:
		{
			body = model.Metrics{ID: params.ID, MType: params.MType, Delta: &params.Delta}
		}
	case model.Gauge:
		{
			body = model.Metrics{ID: params.ID, MType: params.MType, Value: &params.Value}
		}
	}

	req := client.R().SetBody(body).SetResult(&response).SetError(&responseErr)

	res, err := req.Post(client.BaseURL() + "/update")

	if err != nil {
		log.Printf("ошибка обновления метрики %s: %s", params.ID, err.Error())
		return
	}

	log.Printf("метрика %s отправлена: %v", params.ID, res)
}
