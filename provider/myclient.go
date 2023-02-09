package myclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"yandex-devops/config"
)

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
	Hash  string   `json:"hash,omitempty"`  // значение хеш-функции
}

type MyClient struct {
	client *http.Client
	config *config.HTTP
}

func New(config *config.HTTP) MyClient {
	return MyClient{
		client: &http.Client{},
		config: config,
	}
}

func (c MyClient) SendOneMetric(metric Metrics) error {
	u := &url.URL{
		Scheme: c.config.Scheme,
		//Host:   c.config.Host + ":" + c.config.Port,
		Host: c.config.Address,
	}
	u = u.JoinPath("update")

	data, err := json.Marshal(metric)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		u.String(),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (c MyClient) SendAllMetric(metrics *[]Metrics) error {
	for _, m := range *metrics {
		if err := c.SendOneMetric(m); err != nil {
			return err
		}
	}

	return nil
}
