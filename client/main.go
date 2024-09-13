package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Dollar struct {
	Bid string `json:"bid"`
}

func main() {
	serverUrl := "http://localhost:8080/cotacao"

	timeout := 300 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", serverUrl, nil)
	if err != nil {
		log.Panicf("Error creating request: %s", err.Error())
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Panicf("Error making request with timeout: %s", err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Panicf("Error reading response body: %s", err.Error())
	}

	var dollar Dollar

	err = json.Unmarshal(body, &dollar)
	if err != nil {
		log.Panicf("Error decoding response: %s", err.Error())
	}

	file, err := os.Create("client/cotacao.txt")
	if err != nil {
		log.Panicf("Error creating file: %s", err.Error())
	}

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s", dollar.Bid))
	if err != nil {
		log.Panicf("Error writing file: %s", err.Error())
	}
}
