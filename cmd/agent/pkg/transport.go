package transport

import (
	"Komkim/go-musthave-devops-tpl/cmd/agent/config"
	"Komkim/go-musthave-devops-tpl/cmd/agent/storage"
	"fmt"
	"net/http"
)

type MyClient struct {
	client *http.Client
}

func New() MyClient {
	c := &http.Client{}
	client := MyClient{c}
	return client
}

func (c MyClient) SendOne(metric storage.OneMetric) {
	req, err := http.NewRequest(
		http.MethodPost,
		config.DeliveryAddress+":"+config.DeliveryPort+"/update/"+metric.TypeMetric+"/"+metric.Name+"/"+metric.Value,
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
