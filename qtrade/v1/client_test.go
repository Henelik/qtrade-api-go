package v1

import (
	"net/http"
	"testing"
	"time"

	"bou.ke/monkey"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHMAC(t *testing.T) {
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
				Configuration{tc.auth})

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
