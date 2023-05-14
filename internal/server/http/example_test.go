package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ExampleRouter_SaveOrUpdate() {
	values := map[string]string{"id": "HeapIdle", "type": "gauge", "delta": "5"}
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/update/", "application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
	// Output:
	// {
	//    "id": "HeapIdle",
	//    "type": "gauge",
	//    "delta": 5,
	//    "hash": "6d6cd63284be4a47ba7aec4a3458939a95dcbdd5cd0438f23d7457099b4b917c"
	// }
}

//func ExampleRouter_SaveOrUpdateOld() {
//	`   url : localhost:8080/update/gauge/HeapIdle/5
//	method : POST
//	header :
//		Content-Type : application/json
//		Accept-Encoding : gzip
//	body :`
//	// Output:
//	//
//	//	{
//	//	   "id": "HeapIdle",
//	//	   "type": "gauge",
//	//	   "value": 5,
//	//	   "hash": "1efc9ae7e7af8fae3397be7449c1dc389a63b2df3c29a76423870638a216909d"
//	//	}
//}
//
//func ExampleRouter_Ping() {
//	`   url : localhost:8080/updates/
//	method : GET
//	header :
//		Content-Type : application/json
//		Accept-Encoding : gzip
//	`
//	// Output:
//	// "Pong"
//}
//
//func ExampleRouter_SetAll() {
//	`   url : localhost:8080/updates/
//	method : POST
//	header :
//		Content-Type : application/json
//		Accept-Encoding : gzip
//	body :
//		[
//			{"id":"CounterBatchZip100","type":"counter","delta":1668386813,"hash":"9870bcb98440f80b8a47f7b34c3ea9ca17fa5825c9c2b314047c0369f46259c4"},
//			{"id":"GaugeBatchZip55","type":"gauge","value":185171.31281115007,"hash":"15a62829863b2a0e91051028ebf28d9732f0e691544a0ca998f1186bf1108551"},
//			{"id":"CounterBatchZip100","type":"counter","delta":3050210323,"hash":"ec63d8f9dd4796a4bd7e6eb75660773dbb69bc89e96197a5f3bb3666b909b60f"},
//			{"id":"GaugeBatchZip55","type":"gauge","value":927891.479000164,"hash":"7d2496649719506de1d124680f1e88ce32fc992834c503a710a8b2bca02328c2"}
//		]`
//	// Output:
//	// [
//	//
//	//	{
//	//	    "id": "CounterBatchZip100",
//	//	    "type": "counter",
//	//	    "delta": 1668386813,
//	//	    "hash": "9870bcb98440f80b8a47f7b34c3ea9ca17fa5825c9c2b314047c0369f46259c4"
//	//	},
//	//	{
//	//	    "id": "GaugeBatchZip55",
//	//	    "type": "gauge",
//	//	    "value": 185171.31281115007,
//	//	    "hash": "15a62829863b2a0e91051028ebf28d9732f0e691544a0ca998f1186bf1108551"
//	//	},
//	//	{
//	//	    "id": "CounterBatchZip100",
//	//	    "type": "counter",
//	//	    "delta": 4718597136,
//	//	    "hash": "2ca8d0d9c8ca2ddc494f41879a7bcd1d9a790d7bee5721f38fd57471f7d2f78c"
//	//	},
//	//	{
//	//	    "id": "GaugeBatchZip55",
//	//	    "type": "gauge",
//	//	    "value": 927891.479000164,
//	//	    "hash": "7d2496649719506de1d124680f1e88ce32fc992834c503a710a8b2bca02328c2"
//	//	}
//	//
//	// ]
//}
//
//func ExampleRouter_GetAll() {
//	`   url : localhost:8080/
//	method : GET
//	header :
//		Content-Type : application/json
//		Accept-Encoding : gzip
//	`
//	// Output:
//	// [
//	//
//	//	{
//	//	    "id": "HeapIdle",
//	//	    "type": "gauge",
//	//	    "value": 1310720,
//	//	    "hash": "4a6b966657984c0bb78e2b08d9b42b92de169c47a47f2f2450ec0db963906d59"
//	//	},
//	//	{
//	//	    "id": "LastGC",
//	//	    "type": "gauge",
//	//	    "value": 0,
//	//	    "hash": "bbbc148b05e0066845021068f811695aefe3a379d898035e50319d1c51e0f8dd"
//	//	},
//	//	{
//	//	    "id": "MCacheSys",
//	//	    "type": "gauge",
//	//	    "value": 15600,
//	//	    "hash": "80244ed7058cae7f186824230a2c197540b33e838b71fa23f60ec72a71b83666"
//	//	},
//	//	{
//	//	    "id": "OtherSys",
//	//	    "type": "gauge",
//	//	    "value": 1594890,
//	//	    "hash": "a0040a176e2a5ea19e06a2adcf322a87725ddf6047ca9252c7dd412055f3227c"
//	//	},
//	//	{
//	//	    "id": "HeapSys",
//	//	    "type": "gauge",
//	//	    "value": 3670016,
//	//	    "hash": "28c48d0eb9d75529d3c197b80e9635166c5049fa2d499c5580e3596f7e0bf21e"
//	//	},
//	//	{
//	//	    "id": "MCacheInuse",
//	//	    "type": "gauge",
//	//	    "value": 9600,
//	//	    "hash": "0b3db73490a6e0f5130f65b53cbddde177a9d16831982a898b1148604c179495"
//	//	},
//	//	{
//	//	    "id": "MSpanSys",
//	//	    "type": "gauge",
//	//	    "value": 81360,
//	//	    "hash": "2bb4b9a31e7e14973d29dc633de94cd39104cc3a40d2f3219c09ccb77ff52788"
//	//	},
//	//	{
//	//	    "id": "NextGC",
//	//	    "type": "gauge",
//	//	    "value": 4194304,
//	//	    "hash": "18025c8b9f9d46815126da80ae274dbfd69bcd8c1c1fbb1f6c7355b004460840"
//	//	},
//	//	{
//	//	    "id": "CPUutilization1",
//	//	    "type": "gauge",
//	//	    "value": 19.6402920414814,
//	//	    "hash": "e0b0967f78ab8c9412a954d4d6acd3715c911c1d387cb6328b914bd540320ffd"
//	//	},
//	//	{
//	//	    "id": "TotalMemory",
//	//	    "type": "gauge",
//	//	    "value": 25109962752,
//	//	    "hash": "b9e085ecd0de555985bf2fb856750e8d222bc6a0b9788682b0c73e533d70e5f8"
//	//	},
//	//	{
//	//	    "id": "GCCPUFraction",
//	//	    "type": "gauge",
//	//	    "value": 0,
//	//	    "hash": "cf23334951eb3200bc6bfcc503d9d5cd00f763e34b45a02f56610463d442783a"
//	//	},
//	//	{
//	//	    "id": "HeapInuse",
//	//	    "type": "gauge",
//	//	    "value": 2359296,
//	//	    "hash": "c601c852df5f8c25402539ca6002332392dfb3cefa6198e69f9631f5e9b3aeea"
//	//	},
//	//	{
//	//	    "id": "HeapReleased",
//	//	    "type": "gauge",
//	//	    "value": 1015808,
//	//	    "hash": "82040d60b630e9ea785f7c777f400204565c41b1848c66da9e75932e045e98cd"
//	//	},
//	//	{
//	//	    "id": "PollCount",
//	//	    "type": "counter",
//	//	    "delta": 13,
//	//	    "hash": "1181a3f82a7c7cfc1d8e754afcefb30e57475053c3c660b0680d21e22d558df7"
//	//	},
//	//	{
//	//	    "id": "BuckHashSys",
//	//	    "type": "gauge",
//	//	    "value": 3886,
//	//	    "hash": "5d309db0134a1a68488307c11290241f6c02f4accd9aa890fb2eb170aad40839"
//	//	},
//	//	{
//	//	    "id": "Alloc",
//	//	    "type": "gauge",
//	//	    "value": 1181136,
//	//	    "hash": "9d2c15d81fef869820a30ea11d8d8d6cce8f360256002b9bc03a23674a88d5eb"
//	//	},
//	//	{
//	//	    "id": "MSpanInuse",
//	//	    "type": "gauge",
//	//	    "value": 67392,
//	//	    "hash": "4533c3bb2b4a4fc1cf80ab5050f430e4edf0da05ecca22c61872114764982caf"
//	//	},
//	//	{
//	//	    "id": "PauseTotalNs",
//	//	    "type": "gauge",
//	//	    "value": 0,
//	//	    "hash": "0865fe95543fe0a2873bb6b6de7e0fdd40dbf8b12f70f731340bc393dcce9ef7"
//	//	},
//	//	{
//	//	    "id": "StackInuse",
//	//	    "type": "gauge",
//	//	    "value": 524288,
//	//	    "hash": "59e5e956f73f2e4566acea8b985c6a3d0d731789284577033244c8488723f569"
//	//	},
//	//	{
//	//	    "id": "TotalAlloc",
//	//	    "type": "gauge",
//	//	    "value": 1181136,
//	//	    "hash": "355f46b3898b729718e628a89a4d9c35ecfd42a82fd08cb2a0a8f7193ded8599"
//	//	},
//	//	{
//	//	    "id": "RandomValue",
//	//	    "type": "gauge",
//	//	    "value": 0.3874572996788511,
//	//	    "hash": "685ca8b9b88ddc17915d762ba08592cedc33e08d120f6c9a625c42864ea4091a"
//	//	},
//	//	{
//	//	    "id": "GCSys",
//	//	    "type": "gauge",
//	//	    "value": 8117272,
//	//	    "hash": "f6d60f1d5b815f0d5f1e5a2e606b635a47920f5a316ec5001c8b1789a492b217"
//	//	},
//	//	{
//	//	    "id": "Mallocs",
//	//	    "type": "gauge",
//	//	    "value": 11257,
//	//	    "hash": "4ad7475647fa12f25b5fe4148c05d912539e592094e8bbdf71b9b5888f9367ac"
//	//	},
//	//	{
//	//	    "id": "NumForcedGC",
//	//	    "type": "gauge",
//	//	    "value": 0,
//	//	    "hash": "62bc28d96754dad3263afb8cdeb32300a46768e992b0b5a9e7bf5c4f46aa7055"
//	//	},
//	//	{
//	//	    "id": "NumGC",
//	//	    "type": "gauge",
//	//	    "value": 0,
//	//	    "hash": "d9f660c39ea0f15a9b00975775f251f37bc4e80feb01bdcd546688f2a31a14b1"
//	//	},
//	//	{
//	//	    "id": "StackSys",
//	//	    "type": "gauge",
//	//	    "value": 524288,
//	//	    "hash": "6334b00893032aad38af179ed0afda10df4d18c3ccde11334bf2c5a6c665fb2a"
//	//	},
//	//	{
//	//	    "id": "Sys",
//	//	    "type": "gauge",
//	//	    "value": 14007312,
//	//	    "hash": "7aae9f72a482a36a156efaea26bd9a9873b5ee9f3e000d6d707def18f07d74f1"
//	//	},
//	//	{
//	//	    "id": "Frees",
//	//	    "type": "gauge",
//	//	    "value": 922,
//	//	    "hash": "21441dc37cbfd0856f71c02a7ecc361fb99ec518e79e5eb92a80b0c9586f3861"
//	//	},
//	//	{
//	//	    "id": "HeapObjects",
//	//	    "type": "gauge",
//	//	    "value": 10335,
//	//	    "hash": "8b9c113e55c5dae00e2c5d0139dadfabe0143494d3768f46296c845b63363f10"
//	//	},
//	//	{
//	//	    "id": "Lookups",
//	//	    "type": "gauge",
//	//	    "value": 0,
//	//	    "hash": "b9ed5e519bf7b32046f9d44e3079884e2492e923adec6e627008d821e72f46f8"
//	//	},
//	//	{
//	//	    "id": "HeapAlloc",
//	//	    "type": "gauge",
//	//	    "value": 1181136,
//	//	    "hash": "f52ce408c9ed264f2125a04696b935b00d5cc3a9aa1b1fbfa162f7d08408d992"
//	//	},
//	//	{
//	//	    "id": "FreeMemory",
//	//	    "type": "gauge",
//	//	    "value": 14093324288,
//	//	    "hash": "52aa82bddc140dad5af92269473c7c1338b04ee4ddca211a0605e879a5dd085e"
//	//	}
//	//
//	// ]
//}
//
//func ExampleRouter_GetByKey() {
//	`   url : localhost:8080/value
//	method : POST
//	header :
//		Content-Type : application/json
//		Accept-Encoding : gzip
//	body :
//		{
//			"id": "HeapIdle",
//			"type":"gauge"
//		}`
//	// Output:
//	//
//	//	{
//	//	   "id": "HeapIdle",
//	//	   "type": "gauge",
//	//	   "value": 1310720,
//	//	   "hash": "4a6b966657984c0bb78e2b08d9b42b92de169c47a47f2f2450ec0db963906d59"
//	//	}
//}
//
//func ExampleRouter_GetByKeyOld() {
//	`   url : localhost:8080/value/gauge/HeapIdle
//	method : GET
//	header :
//		Content-Type : application/json
//		Accept-Encoding : gzip`
//	// Output:
//	// 1310720
//}
