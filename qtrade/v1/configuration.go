package v1

import (
	"time"
)

type Configuration struct {
	HMACKeypair string        `mapstructure:"hmac_keypair"`
	Endpoint    string        `mapstructure:"endpoint"`
	Timeout     time.Duration `mapstructure:"timeout"`
}
