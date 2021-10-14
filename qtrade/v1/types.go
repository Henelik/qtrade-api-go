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
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
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
	Address         string      `json:"address"`
	Amount          string      `json:"amount"`
	CancelRequested bool        `json:"cancel_requested"`
	CreatedAt       time.Time   `json:"created_at"`
	Currency        string      `json:"currency"`
	ID              int         `json:"id"`
	NetworkData     interface{} `json:"network_data"`
	RelayStatus     string      `json:"relay_status"`
	Status          string      `json:"status"`
	UserID          int         `json:"user_id"`
}

type GetWithdrawHistoryResult struct {
	Data struct {
		Withdraws []WithdrawDetails `json:"withdraws"`
	} `json:"data"`
}
