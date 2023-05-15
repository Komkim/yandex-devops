package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"time"
	"yandex-devops/config"
	"yandex-devops/internal/server/server"
	"yandex-devops/internal/server/service"
	"yandex-devops/storage"
	"yandex-devops/storage/memory"
)

type Example struct {
	Server server.Server
	Config config.Config
}

var example *Example

func getExampleServer() *Example {
	if example == nil {
		example = &Example{}
		cfg, err := config.InitFlagServer()
		if err != nil {
			log.Println(err)
		}
		example.Config = *cfg

		s := memory.NewMemStorage()
		rt := NewRouter(&cfg.Server, service.NewServices(s))
		srv := server.NewServer(&cfg.HTTP, rt.Init())
		example.Server = *srv

		go example.Server.Start()

		time.Sleep(1 * time.Second)

	}
	return example
}

func ExampleRouter_SaveOrUpdate() {
	e := getExampleServer()

	u := &url.URL{
		Scheme: "http",
		Host:   e.Config.HTTP.Address,
	}
	u = u.JoinPath("update")

	log.Println(u.String())

	data := []byte(`{
		"id": "HeapIdle",
		"type": "gauge",
		"value": 5
	}`)

	r, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
	// Output:
	// {"id":"HeapIdle","type":"gauge","value":5}
}

func ExampleRouter_SaveOrUpdateOld() {
	e := getExampleServer()

	u := &url.URL{
		Scheme: "http",
		Host:   e.Config.HTTP.Address,
	}

	u = u.JoinPath("update/gauge/HeapIdle/5")

	r, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
	// Output:
	// {"id":"HeapIdle","type":"gauge","value":5}
}

func ExampleRouter_Ping() {
	e := getExampleServer()

	u := &url.URL{
		Scheme: "http",
		Host:   e.Config.HTTP.Address,
	}
	u = u.JoinPath("ping")

	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
	// Output:
	// "Error connect database"
}

func ExampleRouter_SetAll() {
	data := []byte(`[
{
  "id": "HeapSys",
  "type": "gauge",
  "value": 3702784,
  "hash": "e3ec1cae0b022f109fada933959833ee75a54c58900c6fe6eca8d70195df13e5"
},
{
  "id": "MCacheSys",
  "type": "gauge",
  "value": 15600,
  "hash": "80244ed7058cae7f186824230a2c197540b33e838b71fa23f60ec72a71b83666"
},
{
  "id": "StackSys",
  "type": "gauge",
  "value": 491520,
  "hash": "aeb7b92dc149f025e97a27a0bee5997a55975a943a7171f7dc592b3d3f1c7350"
},
{
  "id": "FreeMemory",
  "type": "gauge",
  "value": 11954253824,
  "hash": "fd7fdec8f8cb7e4e44ef913f39b2c9801b12a07f80fb2872668f1c15c9aebd2f"
}
]`)
	e := getExampleServer()

	u := &url.URL{
		Scheme: "http",
		Host:   e.Config.HTTP.Address,
	}
	u = u.JoinPath("updates")

	r, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
	// Output:
	// [{"id":"HeapSys","type":"gauge","value":3702784,"hash":"e3ec1cae0b022f109fada933959833ee75a54c58900c6fe6eca8d70195df13e5"},{"id":"MCacheSys","type":"gauge","value":15600,"hash":"80244ed7058cae7f186824230a2c197540b33e838b71fa23f60ec72a71b83666"},{"id":"StackSys","type":"gauge","value":491520,"hash":"aeb7b92dc149f025e97a27a0bee5997a55975a943a7171f7dc592b3d3f1c7350"},{"id":"FreeMemory","type":"gauge","value":11954253824,"hash":"fd7fdec8f8cb7e4e44ef913f39b2c9801b12a07f80fb2872668f1c15c9aebd2f"}]
}

func ExampleRouter_GetAll() {
	e := getExampleServer()

	u := &url.URL{
		Scheme: "http",
		Host:   e.Config.HTTP.Address,
	}
	u = u.JoinPath("")

	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	//b, err := io.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//ExampleRouter_SetAll()

	var p []storage.Metrics

	err = json.NewDecoder(res.Body).Decode(&p)
	if err != nil {
		log.Fatal(err)
	}

	sort.SliceStable(p, func(i, j int) bool {
		return p[i].ID < p[j].ID
	})

	out, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(b))
	fmt.Println(string(out))
	// Output:
	// [{"id":"FreeMemory","type":"gauge","value":11954253824,"hash":"fd7fdec8f8cb7e4e44ef913f39b2c9801b12a07f80fb2872668f1c15c9aebd2f"},{"id":"HeapIdle","type":"gauge","value":5},{"id":"HeapSys","type":"gauge","value":3702784,"hash":"e3ec1cae0b022f109fada933959833ee75a54c58900c6fe6eca8d70195df13e5"},{"id":"MCacheSys","type":"gauge","value":15600,"hash":"80244ed7058cae7f186824230a2c197540b33e838b71fa23f60ec72a71b83666"},{"id":"StackSys","type":"gauge","value":491520,"hash":"aeb7b92dc149f025e97a27a0bee5997a55975a943a7171f7dc592b3d3f1c7350"}]
}

func ExampleRouter_GetByKey() {
	data := []byte(`{
		"id": "HeapIdle",
		"type": "gauge"
	}`)
	e := getExampleServer()

	u := &url.URL{
		Scheme: "http",
		Host:   e.Config.HTTP.Address,
	}
	u = u.JoinPath("value")

	r, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
	// Output:
	// {"id":"HeapIdle","type":"gauge","value":5,"hash":"80a5f460b1ed201c1155e1fad7ace28ff6b2482cc50e11a9fbfa56e422e09cda"}
}

//func ExampleRouter_GetByKeyOld() {
//	e := getExampleServer()
//
//	u := &url.URL{
//		Scheme: "http",
//		Host:   e.Config.HTTP.Address,
//	}
//	u = u.JoinPath("value/gauge/HeapIdle")
//
//	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	r.Header.Add("Content-Type", "application/json")
//
//	client := &http.Client{}
//	res, err := client.Do(r)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer res.Body.Close()
//
//	b, err := io.ReadAll(res.Body)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(string(b))
//	// Output:
//	// 5
//}
