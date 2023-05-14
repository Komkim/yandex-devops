package router

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ExampleRouter_SaveOrUpdate() {
	data := []byte(`{
		"id": "HeapIdle",
		"type": "gauge",
		"value": 5
	}`)

	r, err := http.NewRequest("POST", "http://localhost:8080/update/", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept-Encoding", "gzip")

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
	// {"id":"HeapIdle","type":"gauge","value":5,"hash":"1efc9ae7e7af8fae3397be7449c1dc389a63b2df3c29a76423870638a216909d"}
}

func ExampleRouter_SaveOrUpdateOld() {
	r, err := http.NewRequest("POST", "http://localhost:8080/update/gauge/HeapIdle/5", nil)
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
	// {"id":"HeapIdle","type":"gauge","value":5,"hash":"1efc9ae7e7af8fae3397be7449c1dc389a63b2df3c29a76423870638a216909d"}
}

func ExampleRouter_Ping() {
	r, err := http.NewRequest("GET", "http://localhost:8080/ping", nil)
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
	// "Pong"
}

func ExampleRouter_SetAll() {
	data := []byte(`[
			{"id":"CounterBatchZip100","type":"counter","delta":1668386813,"hash":"9870bcb98440f80b8a47f7b34c3ea9ca17fa5825c9c2b314047c0369f46259c4"},
			{"id":"GaugeBatchZip55","type":"gauge","value":185171.31281115007,"hash":"15a62829863b2a0e91051028ebf28d9732f0e691544a0ca998f1186bf1108551"},
			{"id":"CounterBatchZip100","type":"counter","delta":3050210323,"hash":"ec63d8f9dd4796a4bd7e6eb75660773dbb69bc89e96197a5f3bb3666b909b60f"},
			{"id":"GaugeBatchZip55","type":"gauge","value":927891.479000164,"hash":"7d2496649719506de1d124680f1e88ce32fc992834c503a710a8b2bca02328c2"}
		]`)

	r, err := http.NewRequest("POST", "http://localhost:8080/updates/", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept-Encoding", "gzip")

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
	// [
	//
	//	{
	//	    "id": "CounterBatchZip100",
	//	    "type": "counter",
	//	    "delta": 1668386813,
	//	    "hash": "9870bcb98440f80b8a47f7b34c3ea9ca17fa5825c9c2b314047c0369f46259c4"
	//	},
	//	{
	//	    "id": "GaugeBatchZip55",
	//	    "type": "gauge",
	//	    "value": 185171.31281115007,
	//	    "hash": "15a62829863b2a0e91051028ebf28d9732f0e691544a0ca998f1186bf1108551"
	//	},
	//	{
	//	    "id": "CounterBatchZip100",
	//	    "type": "counter",
	//	    "delta": 4718597136,
	//	    "hash": "2ca8d0d9c8ca2ddc494f41879a7bcd1d9a790d7bee5721f38fd57471f7d2f78c"
	//	},
	//	{
	//	    "id": "GaugeBatchZip55",
	//	    "type": "gauge",
	//	    "value": 927891.479000164,
	//	    "hash": "7d2496649719506de1d124680f1e88ce32fc992834c503a710a8b2bca02328c2"
	//	}
	//
	// ]
}

func ExampleRouter_GetAll() {
	r, err := http.NewRequest("GET", "http://localhost:8080/", nil)
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
	// [{"id":"CounterBatchZip100","type":"counter","delta":28311582816,"hash":"281f3eef5fc34b62b227793bc38f0e765d0e369cea1cbc5f2c29202a00073f62"},{"id":"GaugeBatchZip55","type":"gauge","value":927891.479000164,"hash":"7d2496649719506de1d124680f1e88ce32fc992834c503a710a8b2bca02328c2"}]
}

func ExampleRouter_GetByKey() {
	data := []byte(`{
		"id": "HeapIdle",
		"type": "gauge"
	}`)
	r, err := http.NewRequest("POST", "http://localhost:8080/value", bytes.NewBuffer(data))
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
	// {"id":"HeapIdle","type":"gauge","value":1884160,"hash":"f7c297cfa2f13aaa233ce0a02bd70c626d399f53eee4ce93469717a16b8f5859"}
}

func ExampleRouter_GetByKeyOld() {
	r, err := http.NewRequest("GET", "http://localhost:8080/value/gauge/HeapIdle", nil)
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
	// 1884160
}
