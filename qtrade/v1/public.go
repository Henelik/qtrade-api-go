package qtrade

import (
	"context"
	"net/http"

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
		return nil, errors.Wrap(err, "failed to get market")
	}

	return result.Data.Markets, nil
}
