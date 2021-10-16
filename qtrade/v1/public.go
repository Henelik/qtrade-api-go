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
		return nil, errors.Wrap(err, "failed to create buy order")
	}

	return &result.Data, nil
}
