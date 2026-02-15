package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchTrades retrieves recent trades for a market
func FetchTrades(marketID string, limit int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/exchange/spot/v2/trades?marketIds=%s&limit=%d", BASE_URL, marketID, limit)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data, err
}
