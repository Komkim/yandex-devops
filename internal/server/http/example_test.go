package router

//func ExampleRouter_SaveOrUpdate() {
//	data := []byte(`{
//		"id": "HeapIdle",
//		"type": "gauge",
//		"value": 5
//	}`)
//
//	r, err := http.NewRequest("POST", "http://localhost:8080/update/", bytes.NewBuffer(data))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	r.Header.Add("Content-Type", "application/json")
//	//r.Header.Add("Accept-Encoding", "gzip")
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
//	// {"id":"HeapIdle","type":"gauge","value":5,"hash":"1efc9ae7e7af8fae3397be7449c1dc389a63b2df3c29a76423870638a216909d"}
//}

//func ExampleRouter_SaveOrUpdateOld() {
//	r, err := http.NewRequest("POST", "http://localhost:8080/update/gauge/HeapIdle/5", nil)
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
//	// {"id":"HeapIdle","type":"gauge","value":5,"hash":"1efc9ae7e7af8fae3397be7449c1dc389a63b2df3c29a76423870638a216909d"}
//}
//
//func ExampleRouter_Ping() {
//	r, err := http.NewRequest("GET", "http://localhost:8080/ping", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
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
//	// "Pong"
//}
//
//func ExampleRouter_SetAll() {
//	data := []byte(`[
//    {
//        "id": "HeapSys",
//        "type": "gauge",
//        "value": 3702784,
//        "hash": "e3ec1cae0b022f109fada933959833ee75a54c58900c6fe6eca8d70195df13e5"
//    },
//    {
//        "id": "MCacheSys",
//        "type": "gauge",
//        "value": 15600,
//        "hash": "80244ed7058cae7f186824230a2c197540b33e838b71fa23f60ec72a71b83666"
//    },
//    {
//        "id": "StackSys",
//        "type": "gauge",
//        "value": 491520,
//        "hash": "aeb7b92dc149f025e97a27a0bee5997a55975a943a7171f7dc592b3d3f1c7350"
//    },
//    {
//        "id": "FreeMemory",
//        "type": "gauge",
//        "value": 11954253824,
//        "hash": "fd7fdec8f8cb7e4e44ef913f39b2c9801b12a07f80fb2872668f1c15c9aebd2f"
//    }
//]`)
//
//	r, err := http.NewRequest("POST", "http://localhost:8080/updates/", bytes.NewBuffer(data))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	r.Header.Add("Content-Type", "application/json")
//	//r.Header.Add("Accept-Encoding", "gzip")
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
//	// [{"id":"HeapSys","type":"gauge","value":3702784,"hash":"e3ec1cae0b022f109fada933959833ee75a54c58900c6fe6eca8d70195df13e5"},{"id":"MCacheSys","type":"gauge","value":15600,"hash":"80244ed7058cae7f186824230a2c197540b33e838b71fa23f60ec72a71b83666"},{"id":"StackSys","type":"gauge","value":491520,"hash":"aeb7b92dc149f025e97a27a0bee5997a55975a943a7171f7dc592b3d3f1c7350"},{"id":"FreeMemory","type":"gauge","value":11954253824,"hash":"fd7fdec8f8cb7e4e44ef913f39b2c9801b12a07f80fb2872668f1c15c9aebd2f"}]
//}
//
//func ExampleRouter_GetAll() {
//	r, err := http.NewRequest("GET", "http://localhost:8080/", nil)
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
//	// [{"id":"CounterBatchZip100","type":"counter","delta":28311582816,"hash":"281f3eef5fc34b62b227793bc38f0e765d0e369cea1cbc5f2c29202a00073f62"},{"id":"GaugeBatchZip55","type":"gauge","value":927891.479000164,"hash":"7d2496649719506de1d124680f1e88ce32fc992834c503a710a8b2bca02328c2"}]
//}
//
//func ExampleRouter_GetByKey() {
//	data := []byte(`{
//		"id": "HeapIdle",
//		"type": "gauge"
//	}`)
//	r, err := http.NewRequest("POST", "http://localhost:8080/value", bytes.NewBuffer(data))
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
//	// {"id":"HeapIdle","type":"gauge","value":1884160,"hash":"f7c297cfa2f13aaa233ce0a02bd70c626d399f53eee4ce93469717a16b8f5859"}
//}
//
//func ExampleRouter_GetByKeyOld() {
//	r, err := http.NewRequest("GET", "http://localhost:8080/value/gauge/HeapIdle", nil)
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
//	// 1884160
//}
