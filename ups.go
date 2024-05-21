package ups

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"
)

type Environment string

const (
	Testing    Environment = "https://wwwcie.ups.com"
	Production Environment = "https://onlinetools.ups.com"

	shipmentURL = "/api/shipments/v2403/ship"
	oauthURL    = "/security/v1/oauth"
)

type Client struct {
	httpClient *http.Client

	environment Environment

	accessLicenseNumber string

	// authorization
	username                string
	password                string
	clientID                string
	clientSecret            string
	accessToken             string
	accessTokenIsValidUntil time.Time

	logWriter io.Writer
}

type OptionFunction func(*Client)

func New(options ...OptionFunction) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		option(c)
	}

	return c
}

// WithEnvironment defines the used Environment for API requests.
func WithEnvironment(environment Environment) OptionFunction {
	return func(c *Client) {
		c.environment = environment
	}
}

// WithUsernameAndPassword uses username and password for authentication.
func WithUsernameAndPassword(username, password string) OptionFunction {
	return func(c *Client) {
		c.username = username
		c.password = password
	}
}

// WithClientIDAndSecret uses an access token for authentication.
func WithClientIDAndSecret(clientID, clientSecret string) OptionFunction {
	return func(c *Client) {
		c.clientID = clientID
		c.clientSecret = clientSecret
	}
}

// WithAccessLicenseNumber define the access key used in API requests.
func WithAccessLicenseNumber(accessLicenseNumber string) OptionFunction {
	return func(c *Client) {
		c.accessLicenseNumber = accessLicenseNumber
	}
}

// WithHTTPClient uses custom http.Client
func WithHTTPClient(client *http.Client) OptionFunction {
	return func(c *Client) {
		c.httpClient = client
	}
}

// WithLogWriter will write http.Request and http.Response into provided writer
func WithLogWriter(writer io.Writer) OptionFunction {
	return func(c *Client) {
		c.logWriter = writer
	}
}

func (c *Client) addAuthorization(ctx context.Context, req *http.Request) error {
	if c.accessLicenseNumber != "" {
		req.Header.Set("AccessLicenseNumber", c.accessLicenseNumber)
	}

	if c.clientID != "" && c.clientSecret != "" && c.accessTokenIsValidUntil.Before(time.Now()) {
		err := c.getOAuthAccessToken(ctx)
		if err != nil {
			return err
		}
	}

	if c.username != "" && c.password != "" {
		req.Header.Set("Username", c.username)
		req.Header.Set("Password", c.password)
	}

	if c.accessToken != "" {
		req.Header.Set("Authorization", c.accessToken)
	}

	return nil
}

func (c *Client) logHTTPRequest(req *http.Request) error {
	if c.logWriter == nil {
		return nil
	}

	b, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(c.logWriter, string(b))
	return err
}

func (c *Client) logHTTPResponse(res *http.Response) error {
	if c.logWriter == nil {
		return nil
	}

	b, err := httputil.DumpResponse(res, true)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(c.logWriter, string(b))
	return err
}
