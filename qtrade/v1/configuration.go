package v1

import (
	"time"
)

type Configuration struct {
	HMACKeypair string
	Endpoint    string
	Timeout     time.Duration
}
