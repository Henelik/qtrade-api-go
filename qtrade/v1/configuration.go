//go:build !test
// +build !test

package qtrade

import (
	"time"
)

type Configuration struct {
	HMACKeypair string        `mapstructure:"hmac_keypair"`
	Endpoint    string        `mapstructure:"endpoint"`
	Timeout     time.Duration `mapstructure:"timeout"`
	MaxRetries  int           `mapstructure:"max_retries"`
	Backoff     time.Duration `mapstructure:"backoff"`
}
