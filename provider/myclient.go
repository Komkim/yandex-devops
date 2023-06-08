// Модуль отправки метрик по указанному адресу
package myclient

import (
	"bytes"

	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"time"
	"yandex-devops/config"
)

const certFile = "certificat/certificate.crt"

// Metrics - метрики
type Metrics struct {
	//ID - имя метрики
	ID string `json:"id"` // имя метрики
	//MType - тим метрики
	MType string `json:"type"` // параметр, принимающий значение gauge или counter
	//Delta - значение метрики в случае передачи счетчика
	Delta *int64 `json:"delta,omitempty"` // значение метрики в случае передачи counter
	//Value - значение метрики в случае передачи числа с плавающей точкой
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
	//Hash - значение хэш-функции
	Hash string `json:"hash,omitempty"` // значение хеш-функции
}

// MyClient - клиент передачи метрик
type MyClient struct {
	//client - сам клиент
	client *http.Client
	//config - параметры клиента
	config *config.Agent
	url    *url.URL
}

// New - создание нового клиента
func New(config *config.Agent) MyClient {

	if len(config.CryptoKey) > 0 {
		caCert, err := os.ReadFile(certFile)
		if err != nil {
			panic(err)
		}

		caCertPool, err := x509.SystemCertPool()
		if err != nil {
			panic(err)
		}
		caCertPool.AppendCertsFromPEM(caCert)

		t := &http.Transport{
			TLSClientConfig: &tls.Config{
				//Certificates: []tls.Certificate{cert},
				RootCAs: caCertPool,
			},
		}

		client := http.Client{Transport: t, Timeout: 15 * time.Second}
		return MyClient{
			client: &client,
			config: config,
			url: &url.URL{
				Scheme: "https",
				Host:   config.Address,
			},
		}
	}
	return MyClient{
		client: &http.Client{},
		config: config,
		url: &url.URL{
			Scheme: "http",
			Host:   config.Address,
		},
	}

}

// SendOneMetric - отправка одной метрики
func (c MyClient) SendOneMetric(metric Metrics) error {

	u := c.url.JoinPath("update")

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

// SendAllMetric - отправка нескольких метрик
func (c MyClient) SendAllMetric(metrics []Metrics) error {
	u := &url.URL{
		Scheme: c.config.Scheme,
		Host:   c.config.Address,
	}
	u = u.JoinPath("updates")

	data, err := json.Marshal(metrics)
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
