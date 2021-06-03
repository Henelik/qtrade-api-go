package v1

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type QtradeClient struct {
	Client *resty.Client
	Config Configuration
}

func NewQtradeClient(config Configuration) *QtradeClient {
	client := resty.New()

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

	fmt.Printf("reqDetails:\n%s\n", reqDetails)

	return hmac, nil
}
