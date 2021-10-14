package qtrade

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Client struct {
	Client *http.Client
	Config Configuration
	Auth   Auth
}

func NewQtradeClient(config Configuration) (*Client, error) {
	client := &http.Client{
		Timeout: config.Timeout,
	}

	auth, err := AuthFromKeypair(config.HMACKeypair)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: client,
		Config: config,
		Auth:   *auth,
	}, nil
}

func (client *Client) generateHMAC(req *http.Request) (string, string, error) {
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
			return "", "", err
		}

		reqDetails.Write(bodyBytes)
	}

	reqDetails.WriteString("\n")
	reqDetails.WriteString(client.Auth.Key)

	hash := sha256.Sum256(reqDetails.Bytes())

	hmac := "HMAC-SHA256 " +
		client.Auth.KeyID + ":" +
		base64.StdEncoding.EncodeToString(hash[:])

	return hmac, timestamp, nil
}

func (client *Client) doRequest(req *http.Request, result interface{}, queryParams map[string]string) error {
	q := req.URL.Query()

	for k, v := range queryParams {
		q.Add(k, v)
	}

	auth, timestamp, err := client.generateHMAC(req)
	if err != nil {
		return errors.Wrap(err, "could not generate HMAC")
	}

	req.Header.Set("Authorization", auth)
	req.Header.Set("HMAC-Timestamp", timestamp)

	resp, err := client.Client.Do(req)
	if err != nil {
		return errors.Wrap(err, "could not complete HTTP request")
	}

	err = checkForError(resp)
	if err != nil {
		return errors.Wrap(err, "HTTP error")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "could not read response body")
	}

	err = json.Unmarshal(b, result)
	if err != nil {
		return errors.Wrap(err, "could not unmarshal request result")
	}

	return nil
}

func checkForError(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "API response: "+resp.Status)
		}

		apiErrors := new(ErrorResult)

		err = json.Unmarshal(b, apiErrors)
		if err != nil {
			return errors.Wrap(err, "API response: "+resp.Status)
		}

		resultErr := errors.New("API response: " + resp.Status)
		for _, thisErr := range apiErrors.Errors {
			resultErr = errors.Wrap(resultErr,
				fmt.Sprintf("%s: %s", thisErr.Code, thisErr.Title))
		}

		return resultErr
	}

	return nil
}
