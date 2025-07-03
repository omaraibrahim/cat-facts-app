package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const APIKey = "12345-SECRET-API-KEY"

type Fact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type ApiResponse struct {
	Data []Fact `json:"data"`
}

func main() {
	url := "https://api.catfacts.example.com/v1/facts?limit=5"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error")
		return
	}

	req.Header.Set("Authorization", "Bearer "+APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed")
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result ApiResponse
	json.Unmarshal(body, &result)

	fmt.Println("Cat Facts:")
	for _, fact := range result.Data {
		fmt.Println(fact.Fact)
	}
}
