package main

import (
	"encoding/json"
	"fmt"
	"github.com/LeakIX/bannerid"
	"github.com/LeakIX/l9format"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var TBIClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	var pages = make([]int, 200)
	for page ,_ := range pages {
		banners, _ := getWebServerHeaders(page)
		for _ ,banner := range banners {
			software, err := bannerid.ParseWebServerBanner(banner)
			if err == nil {
				softwareJson, _ := json.MarshalIndent(software,"", " ")
				fmt.Println(string(softwareJson))
			}
		}
	}

}

func getWebServerHeaders(page int) ([]string, error) {
	// Query TBI to check if we know this range yet
	req, _ := http.NewRequest("GET",
		fmt.Sprintf(
			"https://leakix.net/search?&q=protocol:http%%20AND%%20cloud&scope=service&page=%d", page),
		nil)
	req.Header.Set("Accept", "application/json")
	resp, err := TBIClient.Do(req)
	if err != nil {
		log.Println(err)
		return []string{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return []string{}, err
	}
	var tbiResults []*l9format.L9Event
	err = json.Unmarshal(body, &tbiResults)
	if err != nil {
		log.Println(err)
		return []string{}, err
	}
	var banners []string
	for _ , service := range tbiResults {
		for header, value := range service.Http.Headers {
			if strings.ToLower(header) == "server" {
				banners = append(banners, value)
			}
		}
	}
	return banners, nil
}