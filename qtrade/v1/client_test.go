package v1

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

type testBody struct {
	data   []byte
	closed bool
}

func (t testBody) Read(p []byte) (n int, err error) {
	p = t.data
	return len(t.data), nil
}

func (t testBody) Close() error {
	if t.closed {
		return io.ErrClosedPipe
	}
	t.closed = true
	return nil
}

const balancesTestData = `{"data": {"balances": [{"balance": "100000000","currency": "BCH"},{"balance": "99992435.78253015","currency": "LTC"},{"balance": "99927153.76074182","currency": "BTC"}]}}`

func TestQtradeClient_generateHMAC(t *testing.T) {
	testCases := []struct {
		name string
		auth Auth
		url  string
		want string
	}{
		{
			name: "no query string",
			auth: Auth{
				KeyID: "256",
				Key:   "vwj043jtrw4o5igw4oi5jwoi45g",
			},
			url:  "http://google.com/",
			want: "HMAC-SHA256 256:iyfC4n+bE+3hLgMJns1Z67FKA7O5qm5PgDvZHGraMTQ=",
		},
		{
			name: "with query string",
			auth: Auth{
				KeyID: "1",
				Key:   "1111111111111111111111111111111111111111111111111111111111111111",
			},
			url:  "https://api.qtrade.io/v1/user/orders?open=false",
			want: "HMAC-SHA256 1:4S8CauoSJcBbQsdcqpqvzN/aFyVJgADXU05eppDxiFA=",
		},
	}

	monkey.Patch(time.Now, func() time.Time {
		return time.Unix(12345, 0)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := NewQtradeClient(
				Configuration{
					Auth:     tc.auth,
					Endpoint: "localhost:420",
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

func TestQtradeClient_GetBalances(t *testing.T) {
	monkey.Patch((*http.Client).Do, func(c *http.Client, req *http.Request) (*http.Response, error) {
		return &http.Response{
			Status:        "OK",
			StatusCode:    200,
			Body:          testBody{data: []byte(balancesTestData)},
			ContentLength: int64(len(balancesTestData)),
		}, nil
	})

	client := NewQtradeClient(
		Configuration{
			Auth: Auth{
				KeyID: "1",
				Key:   "1111111111111111111111111111111111111111111111111111111111111111",
			},
			Endpoint: "http://localhost:420",
		})

	want := GetBalancesResult{
		Data: Balances{
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

	got, _ := client.GetBalances(context.Background())
	assert.Equal(t, want, got)
}
