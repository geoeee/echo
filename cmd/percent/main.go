package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	v1Count int
	v2Count int
)

type echoResponse struct {
	Version string `json:"version"`
}

func getEcho() {
	url := "http://192.168.99.100:30080/echo"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	er := &echoResponse{}
	json.Unmarshal(body, er)
	if er.Version == "1" {
		v1Count++
	}

	if er.Version == "2" {
		v2Count++
	}

	fmt.Printf("er version [%s]\n", er.Version)
	// fmt.Println(res)
	// fmt.Println(string(body))
}

func main() {
	v1Count = 0
	v2Count = 0
	for i := 0; i < 100; i++ {

		go getEcho()
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("count v1: %d, v2: %d\n", v1Count, v2Count)
}
