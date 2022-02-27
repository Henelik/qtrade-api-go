package qtrade

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

var (
	ErrTooManyRequests = errors.New("too many requests")
	ErrHTTPRetryable   = errors.New("a retryable HTTP error occurred")
)

type Client struct {
	Client *http.Client
	Config Configuration
	Auth   Auth
}

func NewClient(config Configuration) (*Client, error) {
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
		bodyBytes := make([]byte, 0)

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
	retries := 0

	q := req.URL.Query()

	for k, v := range queryParams {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	for retries <= client.Config.MaxRetries {
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
		// retry, if applicable
		switch {
		case errors.Is(err, ErrTooManyRequests) && retries < client.Config.MaxRetries:
			retries++

			reset, atoiErr := strconv.Atoi(resp.Header.Get("x-ratelimit-reset"))
			if atoiErr != nil {
				return errors.Wrap(atoiErr, "could not parse ratelimit reset header")
			}

			time.Sleep(time.Duration(reset) * time.Second)

			continue

		case errors.Is(err, ErrHTTPRetryable) && retries < client.Config.MaxRetries:
			retries++
			time.Sleep(client.Config.Backoff)
			continue

		case err != nil:
			return err
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

	return nil
}

func checkForError(resp *http.Response) error {
	if resp.StatusCode == http.StatusTooManyRequests {
		return ErrTooManyRequests
	}

	if resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusServiceUnavailable {
		return ErrHTTPRetryable
	}

	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "failed to read HTTP response body with error "+resp.Status)
		}

		apiErrors := new(ErrorResult)

		err = json.Unmarshal(b, apiErrors)
		if err != nil {
			return fmt.Errorf("got API error with bad JSON: %s: %s", resp.Status, b)
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
