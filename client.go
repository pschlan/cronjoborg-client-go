package cronjoborg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const Endpoint string = "https://api.cron-job.org"

type Client struct {
	Endpoint   string
	HTTPClient *http.Client
	APIKey     string
}

func NewClient(endpoint, apiKey *string) *Client {
	c := Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		Endpoint:   Endpoint,
		APIKey:     *apiKey,
	}

	if endpoint != nil {
		c.Endpoint = *endpoint
	}

	return &c
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
