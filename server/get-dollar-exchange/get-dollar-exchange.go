package getdollarexchange

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type USDBRL struct {
	Bid string `json:"bid" gorm:"bid"`
}

type Dollar struct {
	USDBRL USDBRL `json:"USDBRL"`
}

func GetDollarExchange() (*Dollar, error) {
	dollarUrl := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	timeout := 200 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", dollarUrl, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Error making request with timeout: %s", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	body, error := io.ReadAll(response.Body)
	if error != nil {
		return nil, err
	}

	var dollar Dollar

	error = json.Unmarshal(body, &dollar)
	if error != nil {
		return nil, err
	}

	return &dollar, nil
}
