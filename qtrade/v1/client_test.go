package qtrade

import (
	"net/http"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
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
