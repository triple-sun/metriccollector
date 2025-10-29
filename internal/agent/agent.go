package agent

import (
	"encoding/json"

	"resty.dev/v3"

	"github.com/triple-sun/metriccollector/internal/model"
	"github.com/triple-sun/metriccollector/internal/responses"
	"github.com/triple-sun/metriccollector/internal/utils"
)

type MetricUpdateParams struct {
	ID    string
	MType string
	Value float64
	Delta int64
}

func GetUpdateMetricRequest(client *resty.Client, params MetricUpdateParams) (*resty.Request, error) {
	var data model.Metrics

	switch params.MType {
	case model.Counter:
		{
			data = model.Metrics{ID: params.ID, MType: params.MType, Delta: &params.Delta}
		}
	case model.Gauge:
		{
			data = model.Metrics{ID: params.ID, MType: params.MType, Value: &params.Value}
		}
	}

	body, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	zbody, err := utils.Compress(body)

	if err != nil {
		return nil, err
	}

	return client.R().SetHeader("Content-Encoding", "gzip").SetBody(zbody).SetResult(&responses.MetricUpdateResponse{}).SetError(&responses.ErrorResponse{}), nil
}
