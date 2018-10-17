package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	v1Count     int
	v2Count     int
	prodV1Count int
	prodV2Count int
)

type meta struct {
	Version string `json:"Version"`
	PodName string `json:"PodName"`
}
type product struct {
	Meta meta
}
type company struct {
	CompanyName string `json:"company_name"`
	Meta        meta
	Products    []*product
}
type echoResponse struct {
	Companyies []*company `json:"elements"`
}

type prodResponse struct {
	Meta     meta
	Products []*product `json:"elements"`
}

func getEcho() {
	url := "http://192.168.99.100:31380/company/api/v1/companies"

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

	for _, c := range er.Companyies {
		fmt.Printf("get company version [%s]\n", c.Meta.Version)
		for _, p := range c.Products {
			fmt.Printf("company [%s] product version [%s], pod name [%s]\n", c.CompanyName, p.Meta.Version, p.Meta.PodName)
			if p.Meta.Version == "1.0.0" {
				prodV1Count = prodV1Count + 1
			}
			if p.Meta.Version == "2.0.0" {
				prodV2Count = prodV2Count + 1
			}
		}
	}

	// fmt.Printf("er version [%s]\n", er.Version)
	// fmt.Println(res)
	// fmt.Println(string(body))
}

func getProducts() {
	url := "http://192.168.99.100:31380/product/api/v1/products"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	er := &prodResponse{}
	err = json.Unmarshal(body, er)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(spew.Sdump(er))
	for _, p := range er.Products {
		fmt.Printf("product version [%s], pod name [%s]\n", p.Meta.Version, p.Meta.PodName)
		if p.Meta.Version == "1.0.0" {
			prodV1Count = prodV1Count + 1
		}
		if p.Meta.Version == "2.0.0" {
			prodV2Count = prodV2Count + 1
		}

	}

}

func main() {
	v1Count = 0
	v2Count = 0
	prodV1Count = 0
	prodV2Count = 0
	for i := 0; i < 100; i++ {

		getEcho()
		// getProducts()
	}

	// time.Sleep(5 * time.Second)
	fmt.Printf("prod count v1: %d, v2: %d\n", prodV1Count, prodV2Count)
}
