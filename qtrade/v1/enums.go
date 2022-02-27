//go:build !test
// +build !test

package qtrade

import (
	"strings"
	"time"
)

type Currency string

const (
	USDT  Currency = "USDT"
	BAN   Currency = "BAN"
	BTM   Currency = "BTM"
	DOGE  Currency = "DOGE"
	MCM   Currency = "MCM"
	WEBD  Currency = "WEBD"
	QUAN  Currency = "QUAN"
	PEG   Currency = "PEG"
	RVN   Currency = "RVN"
	BAC   Currency = "BAC"
	CCX   Currency = "CCX"
	XEQ   Currency = "XEQ"
	RUPX  Currency = "RUPX"
	PHL   Currency = "PHL"
	FCT   Currency = "FCT"
	DGB   Currency = "DGB"
	XBR   Currency = "XBR"
	ETH   Currency = "ETH"
	BTC   Currency = "BTC"
	SCC   Currency = "SCC"
	PUSD  Currency = "pUSD"
	REDN  Currency = "REDN"
	SNOW  Currency = "SNOW"
	HTR   Currency = "HTR"
	CPR   Currency = "CPR"
	CRUZ  Currency = "CRUZ"
	VEO   Currency = "VEO"
	KLP   Currency = "KLP"
	VLS   Currency = "VLS"
	NANO  Currency = "NANO"
	DEFT  Currency = "DEFT"
	LTC   Currency = "LTC"
	RCO   Currency = "RCO"
	TAO1  Currency = "TAO1"
	LUCK  Currency = "LUCK"
	WFCT  Currency = "WFCT"
	MMO   Currency = "MMO"
	NYZO  Currency = "NYZO"
	BWS20 Currency = "BWS20"
	BWS10 Currency = "BWS10"
	THC   Currency = "THC"
	XTO   Currency = "XTO"
	GRIN  Currency = "GRIN"
	PASC  Currency = "PASC"
	ARMS  Currency = "ARMS"
	RTM   Currency = "RTM"
	ZANO  Currency = "ZANO"
	ANU   Currency = "ANU"
	XCP   Currency = "XCP"
	ARO   Currency = "ARO"
	BIS   Currency = "BIS"
	IDNA  Currency = "IDNA"
	HLS   Currency = "HLS"
	PFCT  Currency = "pFCT"
)

type CurrencyStatus string

const (
	CurrencyStatusOK       CurrencyStatus = "ok"
	CurrencyStatusDegraded CurrencyStatus = "degraded"
	CurrencyStatusDisabled CurrencyStatus = "disabled"
	CurrencyStatusOffline  CurrencyStatus = "offline"
	CurrencyStatusDelisted CurrencyStatus = "delisted"
)

type Market int

// nolint: golint
const (
	LTC_BTC     Market = 1
	RCO_BTC     Market = 2
	REDN_BTC    Market = 3
	CPR_BTC     Market = 4
	BAC_BTC     Market = 5
	QUAN_BTC    Market = 6
	RVN_BTC     Market = 7
	MMO_BTC     Market = 8
	BTM_BTC     Market = 9
	ANU_BTC     Market = 10
	BWS20_BTC   Market = 11
	BWS20_BWS10 Market = 12
	DEFT_BTC    Market = 13
	RUPX_BTC    Market = 14
	VEO_BTC     Market = 15
	THC_BTC     Market = 16
	SCC_BTC     Market = 17
	XBR_BTC     Market = 18
	SNOW_BTC    Market = 19
	BIS_BTC     Market = 20
	PHL_BTC     Market = 21
	GRIN_BTC    Market = 23
	NYZO_BTC    Market = 24
	TAO1_BTC    Market = 25
	XEQ_BTC     Market = 26
	VLS_BTC     Market = 27
	ZANO_BTC    Market = 28
	PASC_BTC    Market = 30
	NANO_BTC    Market = 31
	CRUZ_BTC    Market = 32
	BAN_BTC     Market = 33
	MCM_BTC     Market = 34
	ARO_BTC     Market = 35
	DOGE_BTC    Market = 36
	HLS_BTC     Market = 37
	WEBD_BTC    Market = 38
	ARMS_BTC    Market = 39
	CCX_BTC     Market = 40
	ETH_BTC     Market = 41
	PEG_BTC     Market = 42
	BTC_pUSD    Market = 43
	ETH_pUSD    Market = 44
	PEG_pUSD    Market = 45
	PFCT_pUSD   Market = 46
	FCT_pUSD    Market = 47
	FCT_BTC     Market = 48
	IDNA_BTC    Market = 49
	DGB_BTC     Market = 50
	KLP_BTC     Market = 51
	XTO_BTC     Market = 52
	LUCK_BTC    Market = 53
	HTR_BTC     Market = 54
	RTM_BTC     Market = 55
	BTC_USDT    Market = 56
	ETH_USDT    Market = 57
	NYZO_USDT   Market = 58
	KLP_USDT    Market = 59
	HTR_USDT    Market = 60
	WFCT_FCT    Market = 61
)

func (m Market) String() string {
	switch m {
	case LTC_BTC:
		return "LTC_BTC"
	case RCO_BTC:
		return "RCO_BTC"
	case REDN_BTC:
		return "REDN_BTC"
	case CPR_BTC:
		return "CPR_BTC"
	case BAC_BTC:
		return "BAC_BTC"
	case QUAN_BTC:
		return "QUAN_BTC"
	case RVN_BTC:
		return "RVN_BTC"
	case MMO_BTC:
		return "MMO_BTC"
	case BTM_BTC:
		return "BTM_BTC"
	case ANU_BTC:
		return "ANU_BTC"
	case BWS20_BTC:
		return "BWS20_BTC"
	case BWS20_BWS10:
		return "BWS20_BWS10"
	case DEFT_BTC:
		return "DEFT_BTC"
	case RUPX_BTC:
		return "RUPX_BTC"
	case VEO_BTC:
		return "VEO_BTC"
	case THC_BTC:
		return "THC_BTC"
	case SCC_BTC:
		return "SCC_BTC"
	case XBR_BTC:
		return "XBR_BTC"
	case SNOW_BTC:
		return "SNOW_BTC"
	case BIS_BTC:
		return "BIS_BTC"
	case PHL_BTC:
		return "PHL_BTC"
	case GRIN_BTC:
		return "GRIN_BTC"
	case NYZO_BTC:
		return "NYZO_BTC"
	case TAO1_BTC:
		return "TAO1_BTC"
	case XEQ_BTC:
		return "XEQ_BTC"
	case VLS_BTC:
		return "VLS_BTC"
	case ZANO_BTC:
		return "ZANO_BTC"
	case PASC_BTC:
		return "PASC_BTC"
	case NANO_BTC:
		return "NANO_BTC"
	case CRUZ_BTC:
		return "CRUZ_BTC"
	case BAN_BTC:
		return "BAN_BTC"
	case MCM_BTC:
		return "MCM_BTC"
	case ARO_BTC:
		return "ARO_BTC"
	case DOGE_BTC:
		return "DOGE_BTC"
	case HLS_BTC:
		return "HLS_BTC"
	case WEBD_BTC:
		return "WEBD_BTC"
	case ARMS_BTC:
		return "ARMS_BTC"
	case CCX_BTC:
		return "CCX_BTC"
	case ETH_BTC:
		return "ETH_BTC"
	case PEG_BTC:
		return "PEG_BTC"
	case BTC_pUSD:
		return "BTC_pUSD"
	case ETH_pUSD:
		return "ETH_pUSD"
	case PEG_pUSD:
		return "PEG_pUSD"
	case PFCT_pUSD:
		return "pFCT_pUSD"
	case FCT_pUSD:
		return "FCT_pUSD"
	case FCT_BTC:
		return "FCT_BTC"
	case IDNA_BTC:
		return "IDNA_BTC"
	case DGB_BTC:
		return "DGB_BTC"
	case KLP_BTC:
		return "KLP_BTC"
	case XTO_BTC:
		return "XTO_BTC"
	case LUCK_BTC:
		return "LUCK_BTC"
	case HTR_BTC:
		return "HTR_BTC"
	case RTM_BTC:
		return "RTM_BTC"
	case BTC_USDT:
		return "BTC_USDT"
	case ETH_USDT:
		return "ETH_USDT"
	case NYZO_USDT:
		return "NYZO_USDT"
	case KLP_USDT:
		return "KLP_USDT"
	case HTR_USDT:
		return "HTR_USDT"
	case WFCT_FCT:
		return "WFCT_FCT"
	}
	return "unknown"
}

func (m Market) MarketCurrency() Currency {
	return Currency(strings.Split(m.String(), "_")[0])
}

func (m Market) BaseCurrency() Currency {
	return Currency(strings.Split(m.String(), "_")[1])
}

type OrderType string

const (
	SellLimit OrderType = "sell_limit"
	BuyLimit  OrderType = "buy_limit"
)

type Interval string

const (
	FiveMin    = "fivemin"
	FifteenMin = "fifteenmin"
	ThirtyMin  = "thirtymin"
	OneHour    = "onehour"
	TwoHour    = "twohour"
	FourHour   = "fourhour"
	OneDay     = "oneday"
)

func (interval Interval) Duration() time.Duration {
	switch interval {
	case FiveMin:
		return time.Minute * 5
	case FifteenMin:
		return time.Minute * 15
	case ThirtyMin:
		return time.Minute * 30
	case OneHour:
		return time.Hour
	case TwoHour:
		return time.Hour * 2
	case FourHour:
		return time.Hour * 4
	case OneDay:
		return time.Hour * 24
	}

	return 0
}
