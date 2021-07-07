package v1

import (
	"errors"
	"strings"
)

type Auth struct {
	KeyID string
	Key   string
}

func AuthFromKeypair(keypair string) (*Auth, error) {
	keys := strings.Split(keypair, ":")

	if len(keys) < 2 {
		return nil, errors.New("AuthFromKeypair: could not parse keypair")
	}

	return &Auth{
		KeyID: keys[0],
		Key:   keys[1],
	}, nil
}
