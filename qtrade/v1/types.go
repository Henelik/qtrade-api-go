package v1

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

// API results

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
