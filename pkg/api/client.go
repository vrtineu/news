package api

import (
	"fmt"
	"net/http"
)

type Client struct {
	BaseURL    string
	APIKey     string
	Httpclient *http.Client
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		Httpclient: &http.Client{},
	}
}

func (c *Client) Get(path string) (*http.Response, error) {
	fmt.Println(c.BaseURL + path + "&apiKey=" + c.APIKey)
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+path+"&apiKey="+c.APIKey, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.Httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
