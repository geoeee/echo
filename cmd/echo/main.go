package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/zhangzhoujian/echo/pkg/version"
)

type response struct {
	PodName string      `json:"pod_name"`
	Version string      `json:"version"`
	Time    time.Time   `json:"time"`
	Headers interface{} `json:"headers"`
}

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {

		log.Println("echo ...")
		resp := &response{
			PodName: os.Getenv("POD_NAME"),
			Version: version.Version,
			Time:    time.Now(),
			Headers: req.Header,
		}

		bt, _ := json.MarshalIndent(resp, "", "  ")
		io.WriteString(w, fmt.Sprintf("%s\n", bt))
	}

	http.HandleFunc("/echo", helloHandler)
	log.Println("starting echo service listen :8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
