package transport

import (
	"log"
	"net/http"
	"net/url"
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
	u := &url.URL{
		Scheme: c.config.Scheme,
		Host:   c.config.Host + ":" + c.config.Port,
	}
	u = u.JoinPath("update")
	u = u.JoinPath(metric.TypeMetric)
	u = u.JoinPath(metric.Name)
	u = u.JoinPath(metric.Value)

	req, err := http.NewRequest(
		http.MethodPost,
		u.String(),
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Add("Content-Type", "text/plain")

	resp, err := c.client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
}

func (c MyClient) SendAll(metrics []storage.OneMetric) {
	for _, v := range metrics {
		c.SendOne(v)
	}
}
