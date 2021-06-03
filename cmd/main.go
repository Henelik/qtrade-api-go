package main

import qtrade "github.com/Henelik/qtrade-api-go/qtrade/v1"

func main () {
	config := qtrade.Configuration{
		Auth: "1:11111111111111111111",
	}

	client := qtrade.NewQtradeClient(config)

	req := client.Client.R().
}
