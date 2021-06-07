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
}

func NewQtradeClient(config Configuration) *QtradeClient {
	client := &http.Client{
		Timeout: 10,
	}

	return &QtradeClient{
		Client: client,
		Config: config,
	}
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

	reqDetails.WriteString(client.Config.Auth.Key)

	hash := sha256.Sum256(reqDetails.Bytes())

	hmac := "HMAC-SHA256 " +
		client.Config.Auth.KeyID + ":" +
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

func (client *QtradeClient) GetBalances(ctx context.Context) (*GetBalancesResult, error) {
	result := new(GetBalancesResult)

	err := client.doRequest(ctx, "GET", "/user/balances", result)

	return result, err
}
