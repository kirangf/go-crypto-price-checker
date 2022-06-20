package client

import (
	"encoding/json"
	"go-api/goexamples/cryptopricechecker/model"
	"io/ioutil"
	"log"
	"net/http"
)

//Fetch is exported crypto data from API
func FetchCrypto(fiat string, crypto string) (string, error) {
	URL := "https://api.nomics.com/v1/currencies/ticker?key=4990ec554a414b59dd85d29b2286dd85&interval=1d&ids=" + crypto + "&convert=" + fiat
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	jsonData := string(body)

	var cResponse model.CryptoResponse
	err = json.Unmarshal([]byte(jsonData), &cResponse)
	if err != nil {
		log.Fatalln(err)
	}

	return cResponse.Output(), nil
}
