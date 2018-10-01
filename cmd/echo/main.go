package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/zhangzhoujian/echo/pkg/version"
)

type response struct {
	Version string    `json:"version"`
	Time    time.Time `json:"time"`
}

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {

		resp := &response{
			Version: version.Version,
			Time:    time.Now(),
		}

		bt, _ := json.MarshalIndent(resp, "", "  ")
		io.WriteString(w, fmt.Sprintf("%s\n", bt))
	}

	http.HandleFunc("/echo", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
