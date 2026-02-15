package clients

import (
	"encoding/json"
	"net/http"
)

const BASE_URL = "https://k8s.mainnet.exchange.injective.network"

// Fetch all markets
func FetchMarkets() (map[string]interface{}, error) {
	resp, err := http.Get(BASE_URL + "/spot/markets")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data, err
}

// Fetch orderbook
func FetchOrderbook(marketID string) (map[string]interface{}, error) {
	url := BASE_URL + "/spot/orderbook?marketId=" + marketID

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data, err
}
