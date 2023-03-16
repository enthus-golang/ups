package ups

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) CreateShipment(ctx context.Context, shipmentRequest ShipmentRequest) (*ShipmentResponse, error) {
	jsonBody, err := json.MarshalIndent(struct {
		ShipmentRequest ShipmentRequest
	}{
		ShipmentRequest: shipmentRequest,
	}, "", "  ")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, string(c.environment), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	c.addAuthorization(req)

	err = c.logHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = c.logHTTPResponse(res)
	if err != nil {
		return nil, err
	}

	var response struct {
		ShipmentResponse ShipmentResponse
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response.ShipmentResponse, nil
}
