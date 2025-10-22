package responses

import (
	models "github.com/triple-sun/metriccollector/internal/model"
)

type MetricUpdateResponse struct {
	Success bool           `json:"success"`
	Metrics models.Metrics `json:"metrics"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
