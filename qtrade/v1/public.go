package qtrade

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
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
