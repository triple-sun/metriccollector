package utils

import (
	"fmt"

	models "github.com/triple-sun/metriccollector/internal/model"
)

func ValidateMType(mtype string) error {
	if mtype != models.Counter && mtype != models.Gauge {
		return fmt.Errorf(`некорректный тип метрики: %s`, mtype)
	}
	return nil
}

func ValidateMName(mname string) error {
	if mname == "" {
		return fmt.Errorf(`не найдено имя метрики`)
	}
	return nil
}

func ValidateMValue(mtype string, mdelta *int64, mvalue *float64) error {
	if mtype == models.Counter && mdelta == nil {
		return fmt.Errorf("отсутствует значение delta")
	}
	if mtype == models.Gauge && mvalue == nil {
		return fmt.Errorf("отсутствует значение value")
	}

	return nil
}
