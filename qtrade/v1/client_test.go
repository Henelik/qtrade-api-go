package v1

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
	userTestData     = `{"data": {"user": {"can_login": true,"can_trade": true,"can_withdraw": true,"email": "hugh@test.com","email_addresses": [{"address": "hugh@test.com","created_at": "2019-10-14T14:41:43.506827Z","id": 10000,"is_primary": true,"verified": true},{"address": "jass@test.com","created_at": "2019-11-14T18:51:23.816532Z","id": 10001,"is_primary": false,"verified": true}],"fname": "Hugh","id": 1000000,"lname": "Jass","referral_code": "6W56QFFVIIJ2","tfa_enabled": true,"verification": "none","verified_email": true,"withdraw_limit": 0}}}`
	balancesTestData = `{"data": {"balances": [{"balance": "100000000","currency": "BCH"},{"balance": "99992435.78253015","currency": "LTC"},{"balance": "99927153.76074182","currency": "BTC"}]}}`
)

var testClient, _ = NewQtradeClient(
	Configuration{
		HMACKeypair: "1:1111111111111111111111111111111111111111111111111111111111111111",
		Endpoint:    "http://localhost",
		Timeout:     time.Second * 10,
	})

func TestQtradeClient_generateHMAC(t *testing.T) {
	testCases := []struct {
		name string
		hmac string
		url  string
		want string
	}{
		{
			name: "no query string",
			hmac: "256:vwj043jtrw4o5igw4oi5jwoi45g",
			url:  "http://google.com/",
			want: "HMAC-SHA256 256:iyfC4n+bE+3hLgMJns1Z67FKA7O5qm5PgDvZHGraMTQ=",
		},
		{
			name: "with query string",
			hmac: "1:1111111111111111111111111111111111111111111111111111111111111111",
			url:  "https://api.qtrade.io/v1/user/orders?open=false",
			want: "HMAC-SHA256 1:4S8CauoSJcBbQsdcqpqvzN/aFyVJgADXU05eppDxiFA=",
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
				got, gotErr := client.generateHMAC(req)
				if assert.NoError(t, gotErr) {
					assert.Equal(t, tc.want, got)
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

	want := &GetUserInfoResult{
		Data: struct {
			User UserInfo `json:"user"`
		}{
			User: UserInfo{
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
			},
		},
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

	want := &GetBalancesResult{
		Data: struct {
			Balances []Balance `json:"balances"`
		}{
			Balances: []Balance{
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
			},
		},
	}

	got, err := testClient.GetBalances(context.Background())
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/user/balances"])
}
