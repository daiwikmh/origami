package services

import (
	"github.com/daiwikmh/origami/clients"
)

func GetMarkets() (map[string]interface{}, error) {
	return clients.FetchMarkets()
}

func GetOrderbook(marketID string) (map[string]interface{}, error) {
	return clients.FetchOrderbook(marketID)
}
