// * Simple Client HTTP Request
// cara mengonsumsi API dengan HTTP Request dan menyimpannya ke json

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sbwhitecap/tqdm"
	"github.com/sbwhitecap/tqdm/iterators"
)

type Quote struct {
	Id      string   `json:"_id"`
	Content string   `json:"content"`
	Author  string   `json:"author"`
	Tags    []string `json:"tags"`
}

type Quotes []Quote

var (
	baseUrl    = "https://api.quotable.io/quotes/random?limit=50"
	timeout    = time.Second * 5
	outputFile = "quotes.json"
)

func fetchQuote() (Quotes, error) {
	client := &http.Client{Timeout: timeout}
	var data Quotes

	request, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return data, nil
}

func main() {
	quotes := Quotes{}
	quoteChan := make(chan Quotes, 3)

	now := time.Now()
	defer func() {
		totalTime := time.Since(now)
		fmt.Printf("Took time:%v\n", totalTime)
	}()

	go func() {
		for quote := range quoteChan {
			quotes = append(quotes, quote...)
		}
	}()

	tqdm.With(iterators.Interval(0, 5), "Get quote data:", func(v interface{}) (brk bool) {
		quote, err := fetchQuote()
		if err != nil {
			log.Println("error: ", err.Error())
		} else {
			quoteChan <- quote
		}
		return
	})
	close(quoteChan)

	file, err := json.MarshalIndent(quotes, "", " ")
	if err != nil {
		log.Println("error: ", err.Error())
		return
	}

	err = ioutil.WriteFile(outputFile, file, 0644)
	if err != nil {
		log.Println("error: ", err.Error())
		return
	}
	fmt.Println("File quotes.json has been saved!")
}
