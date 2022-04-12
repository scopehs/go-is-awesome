package esi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/scopehs/tutorial/database"
	"github.com/scopehs/tutorial/models"
)

func MarketPrices() {
	resp, err := http.Get("https://esi.evetech.net/latest/markets/prices/?datasource=tranquility")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var marketPrices []models.MarkerPrices
	err = json.Unmarshal(body, &marketPrices)

	if err != nil {
		fmt.Println(err)
	}

	log.Println(database.DB.CreateInBatches(&marketPrices, 1000).Error)

}
