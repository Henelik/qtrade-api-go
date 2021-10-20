package qtrade

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

var testClient, _ = NewClient(
	Configuration{
		HMACKeypair: "1:1111111111111111111111111111111111111111111111111111111111111111",
		Endpoint:    "http://localhost",
		Timeout:     time.Second * 10,
	})

func TestClient_generateHMAC(t *testing.T) {
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
			client, _ := NewClient(
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

func TestClient_checkForError(t *testing.T) {
	testCases := []struct {
		name    string
		resp    *http.Response
		wantErr bool
		errMsg  string
	}{
		{
			name: "418 with bad JSON",
			resp: &http.Response{
				Status:     "418 I'm a teapot",
				StatusCode: 418,
				Body:       io.NopCloser(strings.NewReader("short and stout")),
			},
			wantErr: true,
			errMsg:  "got API error with bad JSON: 418 I'm a teapot: short and stout",
		},
		{
			name: "error with valid JSON",
			resp: &http.Response{
				Status:     "403 Forbidden",
				StatusCode: 403,
				Body:       io.NopCloser(strings.NewReader(`{"errors": [{"code": "invalid_auth","title": "Invalid HMAC signature"}]}`)),
			},
			wantErr: true,
			errMsg:  "invalid_auth: Invalid HMAC signature: API response: 403 Forbidden",
		},
		{
			name: "non-error response",
			resp: &http.Response{
				Status:     "403 Forbidden",
				StatusCode: 403,
				Body:       io.NopCloser(strings.NewReader(`{"errors": [{"code": "invalid_auth","title": "Invalid HMAC signature"}]}`)),
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := checkForError(tc.resp)

			if tc.wantErr {
				assert.Equal(t, tc.errMsg, err.Error())
			}
		})
	}
}
