package transport

import (
	"fmt"
	"net/http"
	"yandex-devops/config"
	"yandex-devops/internal/agent/storage"
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

func (c MyClient) SendOne(metric storage.OneMetric) {
	req, err := http.NewRequest(
		http.MethodPost,
		c.config.Host+":"+c.config.Port+"/update/"+metric.TypeMetric+"/"+metric.Name+"/"+metric.Value,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "text/plain")

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
}

func (c MyClient) SendAll(metrics []storage.OneMetric) {
	for _, v := range metrics {
		c.SendOne(v)
	}
}
