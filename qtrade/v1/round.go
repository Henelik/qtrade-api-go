package v1

import (
	"math"

	"github.com/pkg/errors"
)

var (
	Places = map[string]int{
		"BTC":  8,
		"DOGE": 8,
	}
)

// RoundFloat64 rounds x to a specified number of decimal places
func RoundFloat64(x float64, places int) float64 {
	factor := math.Pow(10, float64(places))

	return math.Round(x*factor) / factor
}

func GetPlaces(currency string) (int, error) {
	if p, ok := Places[currency]; ok {
		return p, nil
	}

	return 0, errors.New(currency + " is not a supported currency")
}
