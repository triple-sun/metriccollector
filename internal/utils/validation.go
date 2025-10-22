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
