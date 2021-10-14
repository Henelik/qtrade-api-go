package qtrade

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func (client *QtradeClient) GetUserInfo(ctx context.Context) (*UserInfo, error) {
	result := new(GetUserInfoResult)

	req, err := http.NewRequestWithContext(ctx, "GET", client.Config.Endpoint+"/v1/user/me", nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user info")
	}

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user info")
	}

	return &result.Data.User, nil
}

func (client *QtradeClient) GetBalances(ctx context.Context, params map[string]string) ([]Balance, error) {
	result := new(GetBalancesResult)

	req, err := http.NewRequestWithContext(ctx, "GET", client.Config.Endpoint+"/v1/user/balances", nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get balances")
	}

	err = client.doRequest(req, result, params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get balances")
	}

	return result.Data.Balances, nil
}

func (client *QtradeClient) GetUserMarket(ctx context.Context, market string, params map[string]string) (*UserMarketData, error) {
	result := new(GetUserMarketResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		fmt.Sprintf("%s/v1/user/market/%s", client.Config.Endpoint, market),
		nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user market view")
	}

	err = client.doRequest(req, result, params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user market view")
	}

	return &result.Data, nil
}

func (client *QtradeClient) GetOrders(ctx context.Context, params map[string]string) ([]Order, error) {
	result := new(GetOrdersResult)

	req, err := http.NewRequestWithContext(ctx, "GET", client.Config.Endpoint+"/v1/user/orders", nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get orders")
	}

	err = client.doRequest(req, result, params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get orders")
	}

	return result.Data.Orders, nil
}

func (client *QtradeClient) GetOrder(ctx context.Context, id int) (*Order, error) {
	result := new(GetOrderResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		fmt.Sprintf("%s/v1/user/order/%v", client.Config.Endpoint, id),
		nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get order")
	}

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get order")
	}

	return &result.Data.Order, nil
}

func (client *QtradeClient) GetTrades(ctx context.Context, params map[string]string) ([]Trade, error) {
	result := new(GetTradesResult)

	req, err := http.NewRequestWithContext(ctx, "GET", client.Config.Endpoint+"/v1/user/trades", nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get trades")
	}

	err = client.doRequest(req, result, params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get trades")
	}

	return result.Data.Trades, nil
}

func (client *QtradeClient) CancelOrder(ctx context.Context, id int) error {
	body := map[string]interface{}{
		"id": id,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx,
		"POST",
		client.Config.Endpoint+"/v1/user/cancel_order",
		bytes.NewReader(bodyBytes))
	if err != nil {
		return errors.Wrap(err, "failed to cancel order")
	}

	auth, timestamp, err := client.generateHMAC(req)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", auth)
	req.Header.Set("HMAC-Timestamp", timestamp)

	resp, err := client.Client.Do(req)
	if err != nil {
		return err
	}

	return checkForError(resp)
}

func (client *QtradeClient) Withdraw(ctx context.Context, address string, amount float64, currency string) (*WithdrawData, error) {
	places, err := GetPlaces(currency)
	if err != nil {
		return nil, errors.Wrap(err, "failed to withdraw")
	}

	body := map[string]interface{}{
		"address":  address,
		"amount":   strconv.FormatFloat(amount, 'f', places, 64),
		"currency": currency,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	result := new(WithdrawResult)

	req, err := http.NewRequestWithContext(ctx,
		"POST",
		client.Config.Endpoint+"/v1/user/withdraw",
		bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "failed to withdraw")
	}

	return &result.Data, client.doRequest(req, result, nil)
}
