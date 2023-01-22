package transport

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"yandex-devops/config"
	"yandex-devops/storage"
)

type MyClient struct {
	client *http.Client
	config *config.Config
}

func New(config *config.Config) MyClient {
	return MyClient{
		client: &http.Client{},
		config: config,
	}
}

func (c MyClient) SetOne(metric storage.Metrics) (*storage.Metrics, error) {
	u := &url.URL{
		Scheme: c.config.Scheme,
		//Host:   c.config.Host + ":" + c.config.Port,
		Host: c.config.Address,
	}
	u = u.JoinPath("update")

	data, err := json.Marshal(metric)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		u.String(),
		bytes.NewBuffer(data),
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return &metric, nil
}

func (c MyClient) SetAll(metrics []storage.Metrics) (*[]storage.Metrics, error) {
	mm := make([]storage.Metrics, 0, len(metrics))
	for _, m := range metrics {
		if ss, err := c.SetOne(m); err != nil {
			return nil, err
		} else {
			mm = append(mm, *ss)
		}
	}

	return &mm, nil
}

func (c MyClient) GetOne(key string) (storage.Metrics, error) {
	return storage.Metrics{}, nil
}

func (c MyClient) GetAll() ([]storage.Metrics, error) {
	return []storage.Metrics{}, nil
}
