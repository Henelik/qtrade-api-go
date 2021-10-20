package qtrade

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

func (client *Client) GetCommon(ctx context.Context) (*CommonData, error) {
	result := new(GetCommonResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		client.Config.Endpoint+"/v1/common",
		nil)

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get common data")
	}

	return &result.Data, nil
}

func (client *Client) GetTicker(ctx context.Context, market Market) (*Ticker, error) {
	result := new(GetTickerResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		client.Config.Endpoint+"/v1/ticker/"+market.String(),
		nil)

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ticker")
	}

	return &result.Data, nil
}

func (client *Client) GetTickers(ctx context.Context) ([]Ticker, error) {
	result := new(GetTickersResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		client.Config.Endpoint+"/v1/tickers",
		nil)

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tickers")
	}

	return result.Data.Tickers, nil
}

func (client *Client) GetCurrency(ctx context.Context, currency Currency) (*CurrencyData, error) {
	result := new(GetCurrencyResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		client.Config.Endpoint+"/v1/currency/"+string(currency),
		nil)

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get currency")
	}

	return &result.Data.Currency, nil
}

func (client *Client) GetCurrencies(ctx context.Context) ([]CurrencyData, error) {
	result := new(GetCurrenciesResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		client.Config.Endpoint+"/v1/currencies",
		nil)

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get currencies")
	}

	return result.Data.Currencies, nil
}

func (client *Client) GetMarket(ctx context.Context, market Market) (*GetMarketData, error) {
	result := new(GetMarketResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		client.Config.Endpoint+"/v1/market/"+market.String(),
		nil)

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get market")
	}

	return &result.Data, nil
}

func (client *Client) GetMarkets(ctx context.Context) ([]MarketData, error) {
	result := new(GetMarketsResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		client.Config.Endpoint+"/v1/markets",
		nil)

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get markets")
	}

	return result.Data.Markets, nil
}

func (client *Client) GetMarketTrades(ctx context.Context, market Market) ([]PublicTrade, error) {
	result := new(GetMarketTradesResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		fmt.Sprintf("%s/v1/market/%s/trades", client.Config.Endpoint, market.String()),
		nil)

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get market trades")
	}

	return result.Data.Trades, nil
}

func (client *Client) GetOrderbook(ctx context.Context, market Market) (*Orderbook, error) {
	result := new(GetOrderbookResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		client.Config.Endpoint+"/v1/orderbook/"+market.String(),
		nil)

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get orderbook")
	}

	floatBuy, err := stringMapToFloatMap(result.Data.Buy)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get orderbook")
	}

	floatSell, err := stringMapToFloatMap(result.Data.Sell)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get orderbook")
	}

	orderbook := &Orderbook{
		Buy:        floatBuy,
		LastChange: result.Data.LastChange,
		Sell:       floatSell,
	}

	return orderbook, nil
}

func stringMapToFloatMap(input map[string]string) (map[float64]float64, error) {
	output := make(map[float64]float64, len(input))

	for k, v := range input {
		floatK, err := strconv.ParseFloat(k, 64)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse string map to float map")
		}

		floatV, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse string map to float map")
		}

		output[floatK] = floatV
	}

	return output, nil
}
