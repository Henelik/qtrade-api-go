//go:build !test
// +build !test

package qtrade

import "time"

// base types

type UserInfo struct {
	CanLogin       bool           `json:"can_login"`
	CanTrade       bool           `json:"can_trade"`
	CanWithdraw    bool           `json:"can_withdraw"`
	Email          string         `json:"email"`
	EmailAddresses []EmailAddress `json:"email_addresses"`
	FirstName      string         `json:"fname"`
	LastName       string         `json:"lname"`
	ID             int            `json:"id"`
	ReferralCode   string         `json:"referral_code"`
	TFAEnabled     bool           `json:"tfa_enabled"`
	Verification   string         `json:"verification"`
	VerifiedEmail  bool           `json:"verified_email"`
	WithdrawLimit  int            `json:"withdraw_limit"`
}

type EmailAddress struct {
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	IsPrimary bool      `json:"is_primary"`
	Verified  bool      `json:"verified"`
}

type Balance struct {
	Currency Currency `json:"currency"`
	Balance  string   `json:"balance"`
}

type Order struct {
	BaseAmount            float64        `json:"base_amount,string"`
	CreatedAt             time.Time      `json:"created_at"`
	ID                    int            `json:"id"`
	MarketAmount          float64        `json:"market_amount,string"`
	MarketAmountRemaining float64        `json:"market_amount_remaining,string"`
	Market                Market         `json:"market_id"`
	Open                  bool           `json:"open"`
	OrderType             OrderType      `json:"order_type"`
	Price                 float64        `json:"price,string"`
	Trades                []PrivateTrade `json:"trades"`
	CloseReason           string         `json:"close_reason,omitempty"`
}

// PublicTrade does not contain detailed info about a trade, and is returned by public endpoints.
type PublicTrade struct {
	Amount      float64   `json:"amount,string"`
	CreatedAt   time.Time `json:"created_at"`
	ID          int       `json:"id"`
	Price       float64   `json:"price,string"`
	SellerTaker *bool     `json:"seller_taker,omitempty"`
}

// PrivateTrade contains detailed information about a trade, and is usually only available to one of the users involved.
type PrivateTrade struct {
	BaseAmount   float64   `json:"base_amount,string"`
	BaseFee      float64   `json:"base_fee,string"`
	CreatedAt    time.Time `json:"created_at"`
	ID           int       `json:"id"`
	OrderID      int       `json:"order_id,omitempty"`
	Market       Market    `json:"market_id,omitempty"`
	MarketAmount float64   `json:"market_amount,string"`
	Price        float64   `json:"price,string"`
	Taker        bool      `json:"taker"`
	Side         string    `json:"side,omitempty"`
}

type QtradeError struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Transfer struct {
	Amount         float64                `json:"amount,string"`
	CreatedAt      time.Time              `json:"created_at"`
	Currency       Currency               `json:"currency"`
	ID             int                    `json:"id"`
	ReasonCode     string                 `json:"reason_code"`
	ReasonMetadata map[string]interface{} `json:"reason_metadata"`
	SenderEmail    string                 `json:"sender_email"`
	SenderID       int                    `json:"sender_id"`
}

type Ticker struct {
	Ask             float64 `json:"ask,string"`
	Bid             float64 `json:"bid,string"`
	DayAvgPrice     float64 `json:"day_avg_price,string"`
	DayChange       float64 `json:"day_change,string"`
	DayHigh         float64 `json:"day_high,string"`
	DayLow          float64 `json:"day_low,string"`
	DayOpen         float64 `json:"day_open,string"`
	DayVolumeBase   float64 `json:"day_volume_base,string"`
	DayVolumeMarket float64 `json:"day_volume_market,string"`
	Market          Market  `json:"id"`
	IdHr            string  `json:"id_hr"`
	Last            float64 `json:"last,string"`
}

type CurrencyData struct {
	CanWithdraw bool             `json:"can_withdraw"`
	Code        Currency         `json:"code"`
	Config      CurrencyConfig   `json:"config"`
	LongName    string           `json:"long_name"`
	Metadata    CurrencyMetadata `json:"metadata"`
	Precision   int              `json:"precision"`
	Status      CurrencyStatus   `json:"status"`
	Type        string           `json:"type"`
}

type CurrencyConfig struct {
	AddressVersion                int     `json:"address_version,omitempty"`
	DefaultSigner                 int     `json:"default_signer"`
	Price                         float64 `json:"price"`
	RequiredConfirmations         int     `json:"required_confirmations"`
	RequiredGenerateConfirmations int     `json:"required_generate_confirmations,omitempty"`
	SatoshiPerByte                int     `json:"satoshi_per_byte,omitempty"`
	WifVersion                    int     `json:"wif_version,omitempty"`
	WithdrawFee                   float64 `json:"withdraw_fee,string"`
	ExplorerAddressURL            string  `json:"explorerAddressURL,omitempty"`
	ExplorerTransactionURL        string  `json:"explorerTransactionURL,omitempty"`
	P2ShAddressVersion            int     `json:"p2sh_address_version,omitempty"`
	DataMax                       int     `json:"data_max,omitempty"`
	EnableAddressData             bool    `json:"enable_address_data,omitempty"`
}

type CurrencyMetadata struct {
	DelistingDate   string        `json:"delisting_date,omitempty"`
	WithdrawNotices []interface{} `json:"withdraw_notices,omitempty"`
	DepositNotices  []interface{} `json:"deposit_notices,omitempty"`
	Hidden          bool          `json:"hidden,omitempty"`
}

type MarketData struct {
	BaseCurrency   Currency       `json:"base_currency"`
	CanCancel      bool           `json:"can_cancel"`
	CanTrade       bool           `json:"can_trade"`
	CanView        bool           `json:"can_view"`
	ID             Market         `json:"id"`
	MakerFee       float64        `json:"maker_fee,string"`
	MarketCurrency Currency       `json:"market_currency"`
	Metadata       MarketMetadata `json:"metadata"`
	TakerFee       float64        `json:"taker_fee,string"`
}

type MarketMetadata struct {
	DelistingDate string         `json:"delisting_date,omitempty"`
	MarketNotices []MarketNotice `json:"market_notices,omitempty"`
	Labels        []interface{}  `json:"labels,omitempty"`
}

type MarketNotice struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type Orderbook struct {
	Buy        map[float64]float64 `json:"buy"`
	LastChange int                 `json:"last_change"`
	Sell       map[float64]float64 `json:"sell"`
}

// Public endpoint results

type GetCommonResult struct {
	Data CommonData `json:"data"`
}

type CommonData struct {
	Currencies []CurrencyData `json:"currencies"`
	Markets    []MarketData   `json:"markets"`
	Tickers    []Ticker       `json:"tickers"`
}

type GetTickerResult struct {
	Data Ticker `json:"data"`
}

type GetTickersResult struct {
	Data struct {
		Tickers []Ticker `json:"markets"`
	} `json:"data"`
}

type GetCurrencyResult struct {
	Data struct {
		Currency CurrencyData `json:"currency"`
	} `json:"data"`
}

type GetCurrenciesResult struct {
	Data struct {
		Currencies []CurrencyData `json:"currencies"`
	} `json:"data"`
}

type GetMarketResult struct {
	Data GetMarketData `json:"data"`
}

type GetMarketData struct {
	Market       MarketData    `json:"market"`
	RecentTrades []PublicTrade `json:"recent_trades"`
}

type GetMarketsResult struct {
	Data struct {
		Markets []MarketData `json:"markets"`
	} `json:"data"`
}

type GetMarketTradesResult struct {
	Data struct {
		Trades []PublicTrade `json:"trades"`
	} `json:"data"`
}

type GetOrderbookResult struct {
	Data GetOrderbookData `json:"data"`
}

type GetOrderbookData struct {
	Buy        map[string]string `json:"buy"`
	LastChange int               `json:"last_change"`
	Sell       map[string]string `json:"sell"`
}

type GetOHLCVResult struct {
	Data struct {
		Slices []OHLCVSlice `json:"slices"`
	} `json:"data"`
}

type OHLCVSlice struct {
	Close  float64   `json:"close,string"`
	High   float64   `json:"high,string"`
	Low    float64   `json:"low,string"`
	Open   float64   `json:"open,string"`
	Time   time.Time `json:"time"`
	Volume float64   `json:"volume,string"`
}

// Private endpoint results

type ErrorResult struct {
	Errors []QtradeError `json:"errors"`
}

type GetUserInfoResult struct {
	Data struct {
		User UserInfo `json:"user"`
	} `json:"data"`
}

type GetBalancesResult struct {
	Data struct {
		Balances []Balance `json:"balances"`
	} `json:"data"`
}

type GetUserMarketResult struct {
	Data UserMarketData `json:"data"`
}

type UserMarketData struct {
	BaseBalance   float64 `json:"base_balance,string"`
	ClosedOrders  []Order `json:"closed_orders"`
	MarketBalance float64 `json:"market_balance,string"`
	OpenOrders    []Order `json:"open_orders"`
}

type GetOrdersResult struct {
	Data struct {
		Orders []Order `json:"orders"`
	} `json:"data"`
}

type GetOrderResult struct {
	Data struct {
		Order Order `json:"order"`
	} `json:"data"`
}

type GetTradesResult struct {
	Data struct {
		Trades []PrivateTrade `json:"trades"`
	} `json:"data"`
}

type WithdrawResult struct {
	Data WithdrawData `json:"data"`
}

type WithdrawData struct {
	Code   string `json:"code"`
	ID     int    `json:"id"`
	Result string `json:"result"`
}

type GetWithdrawDetailsResult struct {
	Data struct {
		Withdraw WithdrawDetails `json:"withdraw"`
	} `json:"data"`
}

type WithdrawDetails struct {
	Address         string                 `json:"address"`
	Amount          string                 `json:"amount"`
	CancelRequested bool                   `json:"cancel_requested"`
	CreatedAt       time.Time              `json:"created_at"`
	Currency        Currency               `json:"currency"`
	ID              int                    `json:"id"`
	NetworkData     map[string]interface{} `json:"network_data,omitempty"`
	RelayStatus     string                 `json:"relay_status"`
	Status          string                 `json:"status"`
	UserID          int                    `json:"user_id"`
}

type GetWithdrawHistoryResult struct {
	Data struct {
		Withdraws []WithdrawDetails `json:"withdraws"`
	} `json:"data"`
}

type GetDepositResult struct {
	Data struct {
		Deposit []DepositDetails `json:"deposit"`
	} `json:"data"`
}

type GetDepositHistoryResult struct {
	Data struct {
		Deposits []DepositDetails `json:"deposits"`
	} `json:"data"`
}

type DepositDetails struct {
	Address     string                 `json:"address"`
	Amount      string                 `json:"amount"`
	CreatedAt   time.Time              `json:"created_at"`
	Currency    Currency               `json:"currency"`
	ID          string                 `json:"id"`
	NetworkData map[string]interface{} `json:"network_data,omitempty"`
	RelayStatus string                 `json:"relay_status"`
	Status      string                 `json:"status"`
}

type GetDepositAddressResult struct {
	Data DepositAddressData `json:"data"`
}

type DepositAddressData struct {
	Address        string         `json:"address"`
	CurrencyStatus CurrencyStatus `json:"currency_status"`
}

type GetTransfersResult struct {
	Data struct {
		Transfers []Transfer `json:"transfers"`
	} `json:"data"`
}

type CreateOrderResult struct {
	Data struct {
		Order Order `json:"order"`
	} `json:"data"`
}
