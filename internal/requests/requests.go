package requests

import (
	"net/http"

	"github.com/triple-sun/metriccollector/internal/model"
	"github.com/triple-sun/metriccollector/internal/utils"
)

type UpdateMetricRequest struct {
	ID    string   `json:"id" binding:"required"`
	MType string   `json:"type" binding:"required,oneof=counter gauge"`
	Delta *int64   `json:"delta,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

func (umr *UpdateMetricRequest) Bind(r *http.Request) error {
	return umr.Validate()
}

func (umr *UpdateMetricRequest) Validate() error {
	if err := utils.ValidateMName(umr.ID); err != nil {
		return err
	}
	if err := utils.ValidateMType(umr.MType); err != nil {
		return err
	}
	if err := utils.ValidateMValue(umr.MType, umr.Delta, umr.Value); err != nil {
		return err
	}

	return nil
}

func (umr *UpdateMetricRequest) ParseValue(mvalue string) error {
	if err := utils.ValidateMType(umr.MType); err != nil {
		return err
	}

	switch umr.MType {
	case model.Counter:
		if delta, err := utils.ParseCounterValue(mvalue); err != nil {
			return err
		} else {
			umr.Delta = &delta
		}
	case model.Gauge:
		if value, err := utils.ParseGaugeValue(mvalue); err != nil {
			return err
		} else {
			umr.Value = &value
		}
	}

	return nil
}

type GetMetricValueRequest struct {
	ID    string `json:"id"`
	MType string `json:"type,omitempty"`
}

func (gmvr *GetMetricValueRequest) Bind(r *http.Request) error {
	return gmvr.Validate()
}

func (gmvr *GetMetricValueRequest) Validate() error {
	if err := utils.ValidateMName(gmvr.ID); err != nil {
		return err
	}
	if err := utils.ValidateMType(gmvr.MType); err != nil {
		return err
	}

	return nil
}
