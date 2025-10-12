package utils

import (
	"fmt"
	"strconv"
)

func ParseCounterValue(mvalue string) (int64, error) {
	parsed, err := strconv.ParseInt(mvalue, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("некорректный тип значения для типа Counter")
	}

	return parsed, nil
}

func ParseGaugeValue(mvalue string) (float64, error) {
	parsed, err := strconv.ParseFloat(mvalue, 64)

	if err != nil {
		return 0, fmt.Errorf("некорректный тип значения для типа Gauge")
	}

	return parsed, nil
}
