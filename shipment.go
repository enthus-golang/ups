package ups

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) CreateShipment(ctx context.Context, shipmentRequest ShipmentRequest) error {
	jsonBody, err := json.MarshalIndent(struct {
		ShipmentRequest ShipmentRequest
	}{
		ShipmentRequest: shipmentRequest,
	}, "", "  ")
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, string(c.environment), bytes.NewReader(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	c.addAuthorization(req)

	err = c.logHTTPRequest(req)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	err = c.logHTTPResponse(res)
	if err != nil {
		return err
	}

	return nil
}
