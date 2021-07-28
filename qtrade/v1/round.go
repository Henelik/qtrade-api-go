package v1

import (
	"math"

	"github.com/pkg/errors"
)

var (
	Places = map[string]int{
		"USDT":  6,
		"BAN":   29,
		"BTM":   8,
		"DOGE":  8,
		"MCM":   9,
		"WEBD":  4,
		"QUAN":  8,
		"PEG":   8,
		"RVN":   8,
		"BAC":   8,
		"CCX":   6,
		"XEQ":   4,
		"RUPX":  8,
		"PHL":   8,
		"FCT":   8,
		"DGB":   8,
		"XBR":   8,
		"ETH":   18,
		"BTC":   8,
		"SCC":   8,
		"pUSD":  8,
		"REDN":  8,
		"SNOW":  6,
		"HTR":   2,
		"CPR":   8,
		"CRUZ":  8,
		"VEO":   8,
		"KLP":   12,
		"VLS":   8,
		"NANO":  30,
		"DEFT":  8,
		"LTC":   8,
		"RCO":   8,
		"TAO1":  8,
		"LUCK":  18,
		"WFCT":  8,
		"MMO":   8,
		"NYZO":  6,
		"BWS20": 8,
		"BWS10": 8,
		"THC":   8,
		"XTO":   18,
		"GRIN":  9,
		"PASC":  4,
		"ARMS":  8,
		"RTM":   8,
		"ZANO":  12,
		"ANU":   8,
		"XCP":   8,
		"ARO":   8,
		"BIS":   8,
		"IDNA":  18,
		"HLS":   18,
		"pFCT":  8,
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
