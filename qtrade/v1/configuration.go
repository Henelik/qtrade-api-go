package v1

import (
	"strings"
	"time"
)

type Configuration struct {
	Auth     Auth
	Endpoint string
	Timeout  time.Duration
}

type Auth struct {
	KeyID string
	Key   string
}

func AuthFromKeypair(keypair string) Auth {
	keys := strings.Split(keypair, ":")

	if len(keys) < 2 {
		panic("AuthFromKeypair: could not parse keypair")
	}

	return Auth{
		KeyID: keys[0],
		Key:   keys[1],
	}
}
