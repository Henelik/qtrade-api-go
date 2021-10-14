package qtrade

import (
	"context"
	"net/http"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	userTestData            = `{"data": {"user": {"can_login": true,"can_trade": true,"can_withdraw": true,"email": "hugh@test.com","email_addresses": [{"address": "hugh@test.com","created_at": "2019-10-14T14:41:43.506827Z","id": 10000,"is_primary": true,"verified": true},{"address": "jass@test.com","created_at": "2019-11-14T18:51:23.816532Z","id": 10001,"is_primary": false,"verified": true}],"fname": "Hugh","id": 1000000,"lname": "Jass","referral_code": "6W56QFFVIIJ2","tfa_enabled": true,"verification": "none","verified_email": true,"withdraw_limit": 0}}}`
	balancesTestData        = `{"data": {"balances": [{"balance": "100000000","currency": "BCH"},{"balance": "99992435.78253015","currency": "LTC"},{"balance": "99927153.76074182","currency": "BTC"}]}}`
	userMarketTestData      = `{"data": {"base_balance": "99927153.76074182","closed_orders": [{"base_amount": "0.09102782","created_at": "2018-04-06T17:59:36.366493Z","id": 13252,"market_amount": "4.99896025","market_amount_remaining": "0","market_id": 1,"open": false,"order_type": "buy_limit","price": "9.90682437","trades": [{"base_amount": "49.37394186","base_fee": "0.12343485","created_at": "2018-04-06T17:59:36.366493Z","id": 10289,"market_amount": "4.99298105","price": "9.88866999","taker": true},{"base_amount": "0.05907856","base_fee": "0.00014769","created_at": "2018-04-06T17:59:36.366493Z","id": 10288,"market_amount": "0.0059792","price": "9.88068047","taker": true}]}],"market_balance": "99992435.78253015","open_orders": [{"base_amount": "49.45063516","created_at": "2018-04-06T17:59:35.867526Z","id": 13249,"market_amount": "5.0007505","market_amount_remaining": "5.0007505","market_id": 1,"open": true,"order_type": "buy_limit","price": "9.86398279","trades": null},{"created_at": "2018-04-06T17:59:27.347006Z","id": 13192,"market_amount": "5.00245975","market_amount_remaining": "0.0173805","market_id": 1,"open": true,"order_type": "sell_limit","price": "9.90428849","trades": [{"base_amount": "49.37366303","base_fee": "0.12343415","created_at": "2018-04-06T17:59:27.531716Z","id": 10241,"market_amount": "4.98507925","price": "9.90428849","taker": false}]}]}}`
	ordersTestData          = `{"data": {"orders": [{"base_amount": "0.09102782","created_at": "2018-04-06T17:59:36.366493Z","id": 13252,"market_amount": "4.99896025","market_amount_remaining": "0","market_id": 1,"open": false,"order_type": "buy_limit","price": "9.90682437","trades": [{"base_amount": "49.37394186","base_fee": "0.12343485","created_at": "2018-04-06T17:59:36.366493Z","id": 10289,"market_amount": "4.99298105","price": "9.88866999","taker": true},{"base_amount": "0.05907856","base_fee": "0.00014769","created_at": "2018-04-06T17:59:36.366493Z","id": 10288,"market_amount": "0.0059792","price": "9.88068047","taker": true}]},{"base_amount": "49.33046306","created_at": "2018-04-06T17:59:12.941034Z","id": 13099,"market_amount": "4.9950993","market_amount_remaining": "4.9950993","market_id": 1,"open": true,"order_type": "buy_limit","price": "9.85114439","trades": null}]}}`
	orderTestData           = `{"data": {"order": {"base_amount": "0","close_reason": "canceled","created_at": "2018-11-08T00:15:57.258122Z","id": 8806681,"market_amount": "500","market_amount_remaining": "0","market_id": 36,"open": false,"order_type": "sell_limit","price": "0.00000033","trades": null}}}`
	tradesTestData          = `{"data": {"trades": [{"base_amount": "0.00022751","base_fee": "0","created_at": "2019-10-14T17:42:42.874812Z","id": 63286,"market_amount": "733.93113296","market_id": 36,"order_id": 8141515,"price": "0.00000031","side": "sell","taker": false},{"base_amount": "0.000434","base_fee": "0.00000217","created_at": "2019-10-14T17:42:42.874812Z","id": 63287,"market_amount": "1400","market_id": 36,"order_id": 8141515,"price": "0.00000031","side": "sell","taker": true},{"base_amount": "0.000135","base_fee": "0","created_at": "2019-10-19T11:10:19.387393Z","id": 64129,"market_amount": "500","market_id": 36,"order_id": 8209249,"price": "0.00000027","side": "buy","taker": false}]}}`
	withdrawTestData        = `{"data": {"code": "initiated","id": 3,"result": "Withdraw initiated. Please allow 3-5 minutes for our system to process."}}`
	withdrawDetailsTestData = `{"data": {"withdraw": {"address": "mw67t7AE88SBSRWYw1is3JaFbtXVygwpmB","amount": "1","cancel_requested": false,"created_at": "2019-02-01T06:06:16.218062Z","currency": "LTC","id": 2,"network_data": {},"relay_status": "","status": "needs_create","user_id": 0}}}`
)

var testClient, _ = NewQtradeClient(
	Configuration{
		HMACKeypair: "1:1111111111111111111111111111111111111111111111111111111111111111",
		Endpoint:    "http://localhost",
		Timeout:     time.Second * 10,
	})

func TestQtradeClient_generateHMAC(t *testing.T) {
	testCases := []struct {
		name     string
		hmac     string
		url      string
		wantHMAC string
	}{
		{
			name:     "no query string",
			hmac:     "256:vwj043jtrw4o5igw4oi5jwoi45g",
			url:      "http://google.com/",
			wantHMAC: "HMAC-SHA256 256:iyfC4n+bE+3hLgMJns1Z67FKA7O5qm5PgDvZHGraMTQ=",
		},
		{
			name:     "with query string",
			hmac:     "1:1111111111111111111111111111111111111111111111111111111111111111",
			url:      "https://api.qtrade.io/v1/user/orders?open=false",
			wantHMAC: "HMAC-SHA256 1:4S8CauoSJcBbQsdcqpqvzN/aFyVJgADXU05eppDxiFA=",
		},
	}

	monkey.Patch(time.Now, func() time.Time {
		return time.Unix(12345, 0)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, _ := NewQtradeClient(
				Configuration{
					HMACKeypair: tc.hmac,
					Endpoint:    "localhost:420",
					Timeout:     time.Second * 10,
				})

			req, err := http.NewRequest("GET", tc.url, nil)
			if assert.NoError(t, err) {
				gotHMAC, _, gotErr := client.generateHMAC(req)
				if assert.NoError(t, gotErr) {
					assert.Equal(t, tc.wantHMAC, gotHMAC)
				}
			}
		})
	}
}

func TestNewQtradeClient_GetUserInfo(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/v1/user/me",
		httpmock.NewStringResponder(200, userTestData))

	t1, _ := time.Parse(time.RFC3339, "2019-10-14T14:41:43.506827Z")
	t2, _ := time.Parse(time.RFC3339, "2019-11-14T18:51:23.816532Z")

	want := &UserInfo{
		CanLogin:    true,
		CanTrade:    true,
		CanWithdraw: true,
		Email:       "hugh@test.com",
		EmailAddresses: []EmailAddress{
			{
				Address:   "hugh@test.com",
				CreatedAt: t1,
				ID:        10000,
				IsPrimary: true,
				Verified:  true,
			},
			{
				Address:   "jass@test.com",
				CreatedAt: t2,
				ID:        10001,
				IsPrimary: false,
				Verified:  true,
			},
		},
		FirstName:     "Hugh",
		LastName:      "Jass",
		ID:            1000000,
		ReferralCode:  "6W56QFFVIIJ2",
		TFAEnabled:    true,
		Verification:  "none",
		VerifiedEmail: true,
		WithdrawLimit: 0,
	}

	got, err := testClient.GetUserInfo(context.Background())
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/user/me"])
}

func TestQtradeClient_GetBalances(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/v1/user/balances",
		httpmock.NewStringResponder(200, balancesTestData))

	want := []Balance{
		{
			Currency: "BCH",
			Balance:  "100000000",
		},
		{
			Currency: "LTC",
			Balance:  "99992435.78253015",
		},
		{
			Currency: "BTC",
			Balance:  "99927153.76074182",
		},
	}

	got, err := testClient.GetBalances(context.Background(), nil)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/user/balances"])
}

func TestQtradeClient_GetUserMarket(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/v1/user/market/LTC_BTC",
		httpmock.NewStringResponder(200, userMarketTestData))

	t1, _ := time.Parse(time.RFC3339, "2018-04-06T17:59:36.366493Z")
	t2, _ := time.Parse(time.RFC3339, "2018-04-06T17:59:35.867526Z")
	t3, _ := time.Parse(time.RFC3339, "2018-04-06T17:59:27.347006Z")
	t4, _ := time.Parse(time.RFC3339, "2018-04-06T17:59:27.531716Z")

	want := &UserMarketData{
		BaseBalance: 99927153.76074182,
		ClosedOrders: []Order{
			{
				BaseAmount:            0.09102782,
				CreatedAt:             t1,
				ID:                    13252,
				MarketAmount:          4.99896025,
				MarketAmountRemaining: 0,
				MarketID:              1,
				Open:                  false,
				OrderType:             "buy_limit",
				Price:                 9.90682437,
				Trades: []Trade{
					{
						BaseAmount:   49.37394186,
						BaseFee:      0.12343485,
						CreatedAt:    t1,
						ID:           10289,
						MarketAmount: 4.99298105,
						Price:        9.88866999,
						Taker:        true,
					},
					{
						BaseAmount:   0.05907856,
						BaseFee:      0.00014769,
						CreatedAt:    t1,
						ID:           10288,
						MarketAmount: 0.0059792,
						Price:        9.88068047,
						Taker:        true,
					},
				},
			},
		},
		MarketBalance: 99992435.78253015,
		OpenOrders: []Order{
			{
				BaseAmount:            49.45063516,
				CreatedAt:             t2,
				ID:                    13249,
				MarketAmount:          5.0007505,
				MarketAmountRemaining: 5.0007505,
				MarketID:              1,
				Open:                  true,
				OrderType:             "buy_limit",
				Price:                 9.86398279,
				Trades:                nil,
			},
			{
				BaseAmount:            0,
				CreatedAt:             t3,
				ID:                    13192,
				MarketAmount:          5.00245975,
				MarketAmountRemaining: 0.0173805,
				MarketID:              1,
				Open:                  true,
				OrderType:             "sell_limit",
				Price:                 9.90428849,
				Trades: []Trade{
					{
						BaseAmount:   49.37366303,
						BaseFee:      0.12343415,
						CreatedAt:    t4,
						ID:           10241,
						MarketAmount: 4.98507925,
						Price:        9.90428849,
						Taker:        false,
					},
				},
			},
		},
	}

	got, err := testClient.GetUserMarket(context.Background(), "LTC_BTC", nil)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/user/market/LTC_BTC"])
}

func TestQtradeClient_GetOrders(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/v1/user/orders",
		httpmock.NewStringResponder(200, ordersTestData))

	t1, _ := time.Parse(time.RFC3339, "2018-04-06T17:59:36.366493Z")
	t2, _ := time.Parse(time.RFC3339, "2018-04-06T17:59:12.941034Z")

	want := []Order{
		{
			BaseAmount:            0.09102782,
			CreatedAt:             t1,
			ID:                    13252,
			MarketAmount:          4.99896025,
			MarketAmountRemaining: 0,
			MarketID:              1,
			Open:                  false,
			OrderType:             "buy_limit",
			Price:                 9.90682437,
			Trades: []Trade{
				{
					BaseAmount:   49.37394186,
					BaseFee:      0.12343485,
					CreatedAt:    t1,
					ID:           10289,
					MarketAmount: 4.99298105,
					Price:        9.88866999,
					Taker:        true,
				},
				{
					BaseAmount:   0.05907856,
					BaseFee:      0.00014769,
					CreatedAt:    t1,
					ID:           10288,
					MarketAmount: 0.0059792,
					Price:        9.88068047,
					Taker:        true,
				},
			},
		},
		{
			BaseAmount:            49.33046306,
			CreatedAt:             t2,
			ID:                    13099,
			MarketAmount:          4.9950993,
			MarketAmountRemaining: 4.9950993,
			MarketID:              1,
			Open:                  true,
			OrderType:             "buy_limit",
			Price:                 9.85114439,
			Trades:                []Trade(nil),
		},
	}

	got, err := testClient.GetOrders(context.Background(), nil)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/user/orders"])
}

func TestQtradeClient_GetOrder(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/v1/user/order/8806681",
		httpmock.NewStringResponder(200, orderTestData))

	t1, _ := time.Parse(time.RFC3339, "2018-11-08T00:15:57.258122Z")

	want := &Order{
		BaseAmount:            0,
		CreatedAt:             t1,
		ID:                    8806681,
		MarketAmount:          500,
		MarketAmountRemaining: 0,
		MarketID:              36,
		Open:                  false,
		OrderType:             "sell_limit",
		Price:                 0.00000033,
		Trades:                nil,
		CloseReason:           "canceled",
	}

	got, err := testClient.GetOrder(context.Background(), 8806681)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/user/order/8806681"])
}

func TestQtradeClient_GetTrades(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/v1/user/trades",
		httpmock.NewStringResponder(200, tradesTestData))

	t1, _ := time.Parse(time.RFC3339, "2019-10-14T17:42:42.874812Z")
	t2, _ := time.Parse(time.RFC3339, "2019-10-19T11:10:19.387393Z")

	want := []Trade{
		{
			BaseAmount:   0.00022751,
			BaseFee:      0,
			CreatedAt:    t1,
			ID:           63286,
			OrderID:      8141515,
			MarketID:     36,
			MarketAmount: 733.93113296,
			Price:        0.00000031,
			Taker:        false,
			Side:         "sell",
		},
		{
			BaseAmount:   0.000434,
			BaseFee:      0.00000217,
			CreatedAt:    t1,
			ID:           63287,
			OrderID:      8141515,
			MarketID:     36,
			MarketAmount: 1400,
			Price:        0.00000031,
			Taker:        true,
			Side:         "sell",
		},
		{
			BaseAmount:   0.000135,
			BaseFee:      0,
			CreatedAt:    t2,
			ID:           64129,
			OrderID:      8209249,
			MarketID:     36,
			MarketAmount: 500,
			Price:        0.00000027,
			Taker:        false,
			Side:         "buy",
		},
	}

	got, err := testClient.GetTrades(context.Background(), nil)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/user/trades"])
}

func TestQtradeClient_CancelOrder(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("POST", "http://localhost/v1/user/cancel_order",
		httpmock.NewStringResponder(200, ""))

	err := testClient.CancelOrder(context.Background(), 109)
	assert.NoError(t, err)

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["POST http://localhost/v1/user/cancel_order"])
}

func TestQtradeClient_Withdraw(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("POST", "http://localhost/v1/user/withdraw",
		httpmock.NewStringResponder(200, withdrawTestData))

	want := &WithdrawData{
		Code:   "initiated",
		ID:     3,
		Result: "Withdraw initiated. Please allow 3-5 minutes for our system to process.",
	}

	got, err := testClient.Withdraw(context.Background(), "abcd", 20, "BTC")
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["POST http://localhost/v1/user/withdraw"])
}

func TestQtradeClient_GetWithdrawDetails(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/v1/user/withdraw/2",
		httpmock.NewStringResponder(200, withdrawDetailsTestData))

	wantTime, _ := time.Parse(time.RFC3339Nano, "2019-02-01T06:06:16.218062Z")

	want := &WithdrawDetails{
		Address:         "mw67t7AE88SBSRWYw1is3JaFbtXVygwpmB",
		Amount:          "1",
		CancelRequested: false,
		CreatedAt:       wantTime,
		Currency:        "LTC",
		ID:              2,
		NetworkData:     map[string]interface{}{},
		RelayStatus:     "",
		Status:          "needs_create",
		UserID:          0,
	}

	got, err := testClient.GetWithdrawDetails(context.Background(), 2)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/user/withdraw/2"])
}
