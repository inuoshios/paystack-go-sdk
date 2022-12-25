package paystack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultHTTPTimeout = 60 * time.Second
	baseUrl            = "https://api.paystack.co"
)

type Response map[string]any

type Config struct {
	ApiKey  string
	Client  *http.Client
	baseUrl *url.URL
}

// NewClient instantiates a new paystack client
//
//	client, err := paystack.NewClient(apiKey string)
func NewClient(apiKey string) (*Config, error) {
	parseURL, _ := url.Parse(baseUrl)
	c := &Config{
		ApiKey:  apiKey,
		Client:  httpClient(),
		baseUrl: parseURL,
	}

	return c, nil
}

// makeRequest function makes a request and send a response to the user
func (c *Config) makeRequest(method, path string, body any) ([]byte, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
	}
	parseUrl, err := c.baseUrl.Parse(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	req, err := http.NewRequest(method, fmt.Sprintf(parseUrl.String()+"%s", path), buf)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting a response: %w", err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			return
		}
	}()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response from body: %w", err)
	}

	return response, nil
}

func httpClient() *http.Client {
	var transport http.RoundTripper = &http.Transport{
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
		DisableCompression:  false,
		DisableKeepAlives:   false,
		Proxy:               http.ProxyFromEnvironment,
		TLSHandshakeTimeout: 5 * time.Second,
		ForceAttemptHTTP2:   true,
	}

	client := &http.Client{
		Timeout:   defaultHTTPTimeout,
		Transport: transport,
	}

	return client
}
