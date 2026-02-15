package utils

import (
	"fmt"
	"strconv"
)

// OrderLevel represents a single price level in the orderbook
type OrderLevel struct {
	Price    float64
	Quantity float64
}

// ParseFloat safely converts interface{} to float64
func ParseFloat(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("cannot convert %T to float64", value)
	}
}

// ParseInt safely converts interface{} to int
func ParseInt(value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case int64:
		return int(v), nil
	case float64:
		return int(v), nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, fmt.Errorf("cannot convert %T to int", value)
	}
}

// ParseString safely converts interface{} to string
func ParseString(value interface{}) string {
	if value == nil {
		return ""
	}

	if str, ok := value.(string); ok {
		return str
	}

	return fmt.Sprintf("%v", value)
}

// ExtractOrderbookLevels parses orderbook side (buys/sells) into OrderLevel slice
func ExtractOrderbookLevels(orderbook map[string]interface{}, side string) ([]OrderLevel, error) {
	levels := []OrderLevel{}

	sideData, exists := orderbook[side]
	if !exists {
		return levels, nil
	}

	sideSlice, ok := sideData.([]interface{})
	if !ok {
		return levels, fmt.Errorf("invalid orderbook side format")
	}

	for _, item := range sideSlice {
		level, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		price, err := ParseFloat(level["price"])
		if err != nil {
			continue
		}

		quantity, err := ParseFloat(level["quantity"])
		if err != nil {
			continue
		}

		levels = append(levels, OrderLevel{
			Price:    price,
			Quantity: quantity,
		})
	}

	return levels, nil
}

// SafeGetMap safely extracts a map from interface{}
func SafeGetMap(data interface{}) (map[string]interface{}, bool) {
	if data == nil {
		return nil, false
	}

	m, ok := data.(map[string]interface{})
	return m, ok
}

// SafeGetSlice safely extracts a slice from interface{}
func SafeGetSlice(data interface{}) ([]interface{}, bool) {
	if data == nil {
		return nil, false
	}

	s, ok := data.([]interface{})
	return s, ok
}
