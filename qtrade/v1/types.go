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
	BaseAmount            float64   `json:"base_amount,string"`
	CreatedAt             time.Time `json:"created_at"`
	ID                    int       `json:"id"`
	MarketAmount          float64   `json:"market_amount,string"`
	MarketAmountRemaining float64   `json:"market_amount_remaining,string"`
	MarketID              int       `json:"market_id"`
	Open                  bool      `json:"open"`
	OrderType             string    `json:"order_type"`
	Price                 float64   `json:"price,string"`
	Trades                []Trade   `json:"trades"`
	CloseReason           string    `json:"close_reason,omitempty"`
}

type Trade struct {
	BaseAmount   float64   `json:"base_amount,string"`
	BaseFee      float64   `json:"base_fee,string"`
	CreatedAt    time.Time `json:"created_at"`
	ID           int       `json:"id"`
	OrderID      int       `json:"order_id,omitempty"`
	MarketID     int       `json:"market_id,omitempty"`
	MarketAmount float64   `json:"market_amount,string"`
	Price        float64   `json:"price,string"`
	Taker        bool      `json:"taker"`
	Side         string    `json:"side,omitempty"`
}

type QtradeError struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

// API results

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
		Trades []Trade `json:"trades"`
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

type Transfer struct {
	Amount         string                 `json:"amount"`
	CreatedAt      time.Time              `json:"created_at"`
	Currency       Currency               `json:"currency"`
	ID             int                    `json:"id"`
	ReasonCode     string                 `json:"reason_code"`
	ReasonMetadata map[string]interface{} `json:"reason_metadata"`
	SenderEmail    string                 `json:"sender_email"`
	SenderID       int                    `json:"sender_id"`
}
