package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/edsonjuniordev/client-server-challenge-fullcycle/server/database"
	getdollarexchange "github.com/edsonjuniordev/client-server-challenge-fullcycle/server/get-dollar-exchange"
)

func CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	dollar, err := getdollarexchange.GetDollarExchange()
	if err != nil {
		log.Printf("Error getting dollar exchange: %s\n", err.Error())
	}

	database, err := database.GetDatabase()
	if err != nil {
		log.Panicf("Error getting database: %s\n", err.Error())
	}

	timeout := 10 * time.Millisecond

	databaseContext, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	tx := database.WithContext(databaseContext).Create(dollar.USDBRL)
	if tx.Error != nil {
		log.Printf("Error creating dollar exchange: %s\n", tx.Error.Error())
	}

	json.NewEncoder(w).Encode(dollar.USDBRL)
}
