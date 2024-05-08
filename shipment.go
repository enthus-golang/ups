package ups

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s%s", c.environment, shipmentURL), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	err = c.addAuthorization(ctx, req)
	if err != nil {
		return nil, err
	}

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
		ShipmentResponse *ShipmentResponse
		ErrorResponse    *ErrorResponse `json:"response"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResponse != nil {
		return nil, response.ErrorResponse
	}

	return response.ShipmentResponse, nil
}

func (c *Client) VoidShipment(ctx context.Context, shipmentIdentificationNumber string) (*VoidShipmentResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, fmt.Sprintf("%s%s/cancel/%s", c.environment, shipmentURL, shipmentIdentificationNumber), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	err = c.addAuthorization(ctx, req)
	if err != nil {
		return nil, err
	}

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
		VoidShipmentResponse *VoidShipmentResponse
		ErrorResponse        *ErrorResponse `json:"response"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResponse != nil {
		return nil, response.ErrorResponse
	}

	return response.VoidShipmentResponse, nil

}
