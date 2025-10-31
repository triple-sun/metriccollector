package responses

import (
	models "github.com/triple-sun/metriccollector/internal/model"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type MetricUpdateResponse struct {
	Success bool           `json:"success"`
	Metrics *models.Metrics `json:"metrics"`
}
