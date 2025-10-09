package agent

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Agent struct {
	Client  *http.Client
	baseUrl string
}

type IAgent interface {
	UpdateSingleMetric(mtype string, name string, value string) ([]byte, error)
}

func NewAgent(client *http.Client, baseUrl string) *Agent {
	return &Agent{client, baseUrl}
}

func (a *Agent) UpdateSingleMetric(mtype string, mname string, mvalue string) ([]byte, error) {
	request, reqErr := http.NewRequest(http.MethodPost, a.baseUrl+"/update", nil)

	if reqErr != nil {
		return nil, fmt.Errorf(`не удалось создать запрос %s: %s`, mname, reqErr.Error())
	}

	request.SetPathValue("mtype", mtype)
	request.SetPathValue("mname", mname)
	request.SetPathValue("mvalue", mvalue)

	log.Printf(`Обновляю метрику %s типа %s: %s`, mname, mtype, mvalue)
	res, err := a.Client.Do(request)

	if err != nil {
		return nil, fmt.Errorf(`не удалось обновить метрику %s: %s`, mname, err.Error())
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf(`ошибка чтения тела ответа: %s`, err.Error())
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(`не удалось обновить метрику: [%d][%s]`, res.StatusCode, string(body))
	}

	log.Printf(`Обновлена метрика %s типа %s: %s`, mname, mtype, mvalue)

	return body, nil
}
