package v1

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type QtradeClient struct {
	Client *http.Client
	Config Configuration
	Auth   Auth
}

func NewQtradeClient(config Configuration) (*QtradeClient, error) {
	client := &http.Client{
		Timeout: config.Timeout,
	}

	auth, err := AuthFromKeypair(config.HMACKeypair)
	if err != nil {
		return nil, err
	}

	return &QtradeClient{
		Client: client,
		Config: config,
		Auth:   *auth,
	}, nil
}

func (client *QtradeClient) generateHMAC(req *http.Request) (string, error) {
	timestamp := fmt.Sprintf("%v", time.Now().Unix())

	reqDetails := bytes.NewBufferString(req.Method)
	reqDetails.WriteString("\n")
	reqDetails.WriteString(req.URL.RequestURI())
	reqDetails.WriteString("\n")
	reqDetails.WriteString(timestamp)
	reqDetails.WriteString("\n")

	if req.Body != nil {
		bodyBytes := []byte{}

		_, err := req.Body.Read(bodyBytes)
		if err != nil {
			return "", err
		}

		reqDetails.Write(bodyBytes)
	}

	reqDetails.WriteString("\n")
	reqDetails.WriteString(client.Auth.Key)

	hash := sha256.Sum256(reqDetails.Bytes())

	hmac := "HMAC-SHA256 " +
		client.Auth.KeyID + ":" +
		base64.StdEncoding.EncodeToString(hash[:])

	return hmac, nil
}

func (client *QtradeClient) doRequest(req *http.Request, result interface{}, queryParams map[string]string) error {
	err := req.ParseForm()
	if err != nil {
		return err
	}

	for k, v := range queryParams {
		req.Form.Set(k, v)
	}

	auth, err := client.generateHMAC(req)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", auth)

	resp, err := client.Client.Do(req)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, result)
}

func (client *QtradeClient) GetUserInfo(ctx context.Context) (*GetUserInfoResult, error) {
	result := new(GetUserInfoResult)

	req, err := http.NewRequestWithContext(ctx, "GET", client.Config.Endpoint+"/v1/user/me", nil)
	if err != nil {
		return nil, err
	}

	return result, client.doRequest(req, result, nil)
}

func (client *QtradeClient) GetBalances(ctx context.Context, params map[string]string) (*GetBalancesResult, error) {
	result := new(GetBalancesResult)

	req, err := http.NewRequestWithContext(ctx, "GET", client.Config.Endpoint+"/v1/user/balances", nil)
	if err != nil {
		return nil, err
	}

	return result, client.doRequest(req, result, params)
}

func (client *QtradeClient) GetUserMarket(ctx context.Context, market string, params map[string]string) (*GetUserMarketResult, error) {
	result := new(GetUserMarketResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		fmt.Sprintf("%s/v1/user/market/%s", client.Config.Endpoint, market),
		nil)
	if err != nil {
		return nil, err
	}

	return result, client.doRequest(req, result, params)
}

func (client *QtradeClient) GetOrders(ctx context.Context, params map[string]string) (*GetOrdersResult, error) {
	result := new(GetOrdersResult)

	req, err := http.NewRequestWithContext(ctx, "GET", client.Config.Endpoint+"/v1/user/orders", nil)
	if err != nil {
		return nil, err
	}

	return result, client.doRequest(req, result, params)
}

func (client *QtradeClient) GetOrder(ctx context.Context, id int) (*GetOrderResult, error) {
	result := new(GetOrderResult)

	req, err := http.NewRequestWithContext(ctx, "GET",
		fmt.Sprintf("%s/v1/user/order/%v", client.Config.Endpoint, id),
		nil)
	if err != nil {
		return nil, err
	}

	return result, client.doRequest(req, result, nil)
}

func (client *QtradeClient) GetTrades(ctx context.Context, params map[string]string) (*GetTradesResult, error) {
	result := new(GetTradesResult)

	req, err := http.NewRequestWithContext(ctx, "GET", client.Config.Endpoint+"/v1/user/trades", nil)
	if err != nil {
		return nil, err
	}

	return result, client.doRequest(req, result, params)
}

func (client *QtradeClient) CancelOrder(ctx context.Context, id int) (*CancelOrderResult, error) {
	result := new(CancelOrderResult)

	body := map[string]interface{}{
		"id": id,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bodyBytes))

	req, err := http.NewRequestWithContext(ctx,
		"POST",
		client.Config.Endpoint+"/v1/user/cancel_order",
		bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	return result, client.doRequest(req, result, nil)
}
