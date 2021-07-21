package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundFloat64(t *testing.T) {
	testCases := []struct {
		name   string
		x      float64
		places int
		want   float64
	}{
		{
			name:   "Round 0 to 2 places",
			x:      0,
			places: 2,
			want:   0,
		},
		{
			name:   "Round 101 to 2 places",
			x:      101,
			places: 2,
			want:   101,
		},
		{
			name:   "Round 12.345 to 2 places",
			x:      12.345,
			places: 2,
			want:   12.35,
		},
		{
			name:   "Round 12.3454321 to 3 places",
			x:      12.3454321,
			places: 3,
			want:   12.345,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := RoundFloat64(tc.x, tc.places)

			assert.Equal(t, tc.want, got)
		})
	}
}
