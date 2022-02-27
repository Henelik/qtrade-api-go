package qtrade

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

func (client *Client) GetUserInfo(ctx context.Context) (*UserInfo, error) {
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

func (client *Client) GetBalances(ctx context.Context, params map[string]string) ([]Balance, error) {
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

func (client *Client) GetUserMarket(ctx context.Context, market Market, params map[string]string) (*UserMarketData, error) {
	result := new(GetUserMarketResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		fmt.Sprintf("%s/v1/user/market/%s", client.Config.Endpoint, market.String()),
		nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user market view for "+market.String())
	}

	err = client.doRequest(req, result, params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user market view for "+market.String())
	}

	return &result.Data, nil
}

func (client *Client) GetOrders(ctx context.Context, params map[string]string) ([]Order, error) {
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

func (client *Client) GetOrder(ctx context.Context, id int) (*Order, error) {
	result := new(GetOrderResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		fmt.Sprintf("%s/v1/user/order/%v", client.Config.Endpoint, id),
		nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get order %v", id))
	}

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get order %v", id))
	}

	return &result.Data.Order, nil
}

func (client *Client) GetTrades(ctx context.Context, params map[string]string) ([]PrivateTrade, error) {
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

func (client *Client) CancelOrder(ctx context.Context, id int) error {
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
		return errors.Wrap(err, fmt.Sprintf("failed to cancel order %v", id))
	}

	auth, timestamp, err := client.generateHMAC(req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to cancel order %v", id))
	}

	req.Header.Set("Authorization", auth)
	req.Header.Set("HMAC-Timestamp", timestamp)

	resp, err := client.Client.Do(req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to cancel order %v", id))
	}

	return checkForError(resp)
}

func (client *Client) Withdraw(ctx context.Context, address string, amount float64, currency Currency) (*WithdrawData, error) {
	body := map[string]interface{}{
		"address":  address,
		"amount":   strconv.FormatFloat(amount, 'f', CurrencyDecimalPlaces[currency], 64),
		"currency": currency,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to withdraw "+string(currency))
	}

	result := new(WithdrawResult)

	req, err := http.NewRequestWithContext(ctx,
		"POST",
		client.Config.Endpoint+"/v1/user/withdraw",
		bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "failed to withdraw "+string(currency))
	}

	return &result.Data, client.doRequest(req, result, nil)
}

func (client *Client) GetWithdrawDetails(ctx context.Context, id int) (*WithdrawDetails, error) {
	result := new(GetWithdrawDetailsResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		client.Config.Endpoint+"/v1/user/withdraw/"+strconv.Itoa(id),
		nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get details for withdrawal %v", id))
	}

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get details for withdrawal %v", id))
	}

	return &result.Data.Withdraw, nil
}

func (client *Client) GetWithdrawHistory(ctx context.Context, params map[string]string) ([]WithdrawDetails, error) {
	result := new(GetWithdrawHistoryResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		client.Config.Endpoint+"/v1/user/withdraws",
		nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get withdraw history")
	}

	err = client.doRequest(req, result, params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get withdraw history")
	}

	return result.Data.Withdraws, nil
}

func (client *Client) GetDeposit(ctx context.Context, id string) ([]DepositDetails, error) {
	result := new(GetDepositResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		client.Config.Endpoint+"/v1/user/deposit/"+id,
		nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get deposit %v", id))
	}

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get deposit %v", id))
	}

	return result.Data.Deposit, nil
}

func (client *Client) GetDepositHistory(ctx context.Context, params map[string]string) ([]DepositDetails, error) {
	result := new(GetDepositHistoryResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		client.Config.Endpoint+"/v1/user/deposits",
		nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get deposit history")
	}

	err = client.doRequest(req, result, params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get deposit history")
	}

	return result.Data.Deposits, nil
}

func (client *Client) GetDepositAddress(ctx context.Context, currency Currency) (*DepositAddressData, error) {
	result := new(GetDepositAddressResult)

	req, err := http.NewRequestWithContext(ctx, "POST",
		client.Config.Endpoint+"/v1/user/deposit_address/"+string(currency),
		nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get deposit address for "+string(currency))
	}

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get deposit address for "+string(currency))
	}

	return &result.Data, nil
}

func (client *Client) GetTransfers(ctx context.Context, params map[string]string) ([]Transfer, error) {
	result := new(GetTransfersResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		client.Config.Endpoint+"/v1/user/transfers",
		nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transfers")
	}

	err = client.doRequest(req, result, params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transfers")
	}

	return result.Data.Transfers, nil
}

func (client *Client) CreateSellLimit(ctx context.Context, amount float64, market Market, price float64) (*Order, error) {
	result := new(CreateOrderResult)

	body := map[string]interface{}{
		"amount":    strconv.FormatFloat(amount, 'f', CurrencyDecimalPlaces[market.MarketCurrency()], 64),
		"market_id": market,
		"price":     strconv.FormatFloat(price, 'f', CurrencyDecimalPlaces[market.BaseCurrency()], 64),
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create sell order for "+market.String())
	}

	req, err := http.NewRequestWithContext(ctx,
		"POST",
		client.Config.Endpoint+"/v1/user/sell_limit",
		bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create sell order")
	}

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create sell order for "+market.String())
	}

	return &result.Data.Order, nil
}

func (client *Client) CreateBuyLimit(ctx context.Context, amount float64, market Market, price float64) (*Order, error) {
	result := new(CreateOrderResult)

	body := map[string]interface{}{
		"amount":    strconv.FormatFloat(amount, 'f', CurrencyDecimalPlaces[market.MarketCurrency()], 64),
		"market_id": market,
		"price":     strconv.FormatFloat(price, 'f', CurrencyDecimalPlaces[market.BaseCurrency()], 64),
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create buy order for "+market.String())
	}

	req, err := http.NewRequestWithContext(ctx,
		"POST",
		client.Config.Endpoint+"/v1/user/buy_limit",
		bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create buy order")
	}

	err = client.doRequest(req, result, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create buy order for "+market.String())
	}

	return &result.Data.Order, nil
}
