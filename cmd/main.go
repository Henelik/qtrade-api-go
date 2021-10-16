//go:build !test
// +build !test

package main

import (
	"context"
	"fmt"

	"github.com/spf13/viper"

	v1 "github.com/Henelik/qtrade-api-go/qtrade/v1"
)

func main() {
	config := new(v1.Configuration)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("config: %#v\n", config)

	client, err := v1.NewQtradeClient(*config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("auth: %#v\n", client.Auth)

	balances, err := client.GetBalances(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("balances: %#v\n", balances)
}
