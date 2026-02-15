package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BASE_URL = "https://sentry.exchange.grpc-web.injective.network:443"

// Fetch all markets
func FetchMarkets() (map[string]interface{}, error) {
	resp, err := http.Get(BASE_URL + "/api/exchange/spot/v1/markets")
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
	url := fmt.Sprintf("%s/api/exchange/spot/v2/orderbook/%s", BASE_URL, marketID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data, err
}
