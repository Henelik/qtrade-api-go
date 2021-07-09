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

func (client *QtradeClient) doRequest(ctx context.Context, method string, uri string, result interface{}, queryParams map[string]string) error {
	req, err := http.NewRequestWithContext(ctx, method, client.Config.Endpoint+uri, nil)
	if err != nil {
		return err
	}

	auth, err := client.generateHMAC(req)
	if err != nil {
		return err
	}

	err = req.ParseForm()
	if err != nil {
		return err
	}

	for k, v := range queryParams {
		req.Form.Set(k, v)
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

	err := client.doRequest(ctx, "GET", "/v1/user/me", result, nil)

	return result, err
}

func (client *QtradeClient) GetBalances(ctx context.Context, params map[string]string) (*GetBalancesResult, error) {
	result := new(GetBalancesResult)

	err := client.doRequest(ctx, "GET", "/v1/user/balances", result, params)

	return result, err
}

func (client *QtradeClient) GetUserMarket(ctx context.Context, market string, params map[string]string) (*GetUserMarketResult, error) {
	result := new(GetUserMarketResult)

	err := client.doRequest(ctx, "GET", "/v1/user/market/"+market, result, params)

	return result, err
}

func (client *QtradeClient) GetOrders(ctx context.Context, params map[string]string) (*GetOrdersResult, error) {
	result := new(GetOrdersResult)

	err := client.doRequest(ctx, "GET", "/v1/user/orders", result, params)

	return result, err
}

func (client *QtradeClient) GetOrder(ctx context.Context, id int) (*GetOrderResult, error) {
	result := new(GetOrderResult)

	err := client.doRequest(ctx, "GET", fmt.Sprintf("/v1/user/order/%v", id), result, nil)

	return result, err
}
