package qtrade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthFromKeypair(t *testing.T) {
	testCases := []struct {
		name    string
		keypair string
		want    *Auth
		wantErr bool
	}{
		{
			name:    "empty string results in error",
			keypair: "",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "no colon in input results in error",
			keypair: "4206969",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "multiple colons in input results in error",
			keypair: "420:69:69",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "valid keypair returns auth",
			keypair: "420:6969",
			want: &Auth{
				KeyID: "420",
				Key:   "6969",
			},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := AuthFromKeypair(tc.keypair)

			assert.Equal(t, tc.want, got)

			switch tc.wantErr {
			case true:
				assert.Error(t, err)
			case false:
				assert.NoError(t, err)
			}
		})
	}
}
