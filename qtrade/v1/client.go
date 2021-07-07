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

func (client *QtradeClient) doRequest(ctx context.Context, method string, uri string, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, client.Config.Endpoint+uri, nil)
	if err != nil {
		return err
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

	fmt.Println(json.Unmarshal(b, &map[string]interface{}{}))

	return json.Unmarshal(b, result)
}

func (client *QtradeClient) GetUserInfo(ctx context.Context) (*GetUserInfoResult, error) {
	result := new(GetUserInfoResult)

	err := client.doRequest(ctx, "GET", "/v1/user/me", result)

	return result, err
}

func (client *QtradeClient) GetBalances(ctx context.Context) (*GetBalancesResult, error) {
	result := new(GetBalancesResult)

	err := client.doRequest(ctx, "GET", "/v1/user/balances", result)

	return result, err
}
