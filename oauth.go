package ups

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type OAuthToken struct {
	TokenType   string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
	ClientID    string `json:"client_id"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Status      string `json:"status"`
}

func (c *Client) getOAuthAccessToken(ctx context.Context) error {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s%s/token", c.environment, oauthURL), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.clientID, c.clientSecret)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unknown status code: (%d) %s", res.StatusCode, res.Status)
	}

	var token OAuthToken

	decoder := json.NewDecoder(res.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&token)
	if err != nil {
		return err
	}

	expiresIn, err := strconv.Atoi(token.ExpiresIn)
	if err != nil {
		return err
	}

	c.accessToken = token.TokenType + " " + token.AccessToken
	c.accessTokenIsValidUntil = time.Now().Add(time.Duration(expiresIn) * time.Second)

	return nil
}
