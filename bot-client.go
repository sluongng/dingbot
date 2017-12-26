package dingbot

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Client for DingBot integration
type Client struct {
	BaseURL     *url.URL
	AccessToken string
	UserAgent   string
	httpClient  *http.Client
}

// TextMessage create a message with Text type
func (c *Client) TextMessage() error {

	//Todo implement Client.TextMessage()

	return nil
}

// LinkMessage create a message with Link type
func (c *Client) LinkMessage() error {

	//Todo implement Client.LinkMessage()

	return nil
}

// MarkdownMessage create a message with Markdown type
func (c *Client) MarkdownMessage() error {

	//Todo implement Client.MarkdownMessage()

	return nil
}

// ActionCardMessage create a message with ActionCard type
func (c *Client) ActionCardMessage() error {

	//Todo implement Client.ActionCardMessage()

	return nil
}

// FeedCardMessage create a message with FeedCard type
func (c *Client) FeedCardMessage() error {

	//Todo implement Client.FeedCardMessage()

	return nil
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	// Set URL
	rel := &url.URL{
		Path:     path,
		RawQuery: "access_token=" + c.AccessToken,
	}
	u := c.BaseURL.ResolveReference(rel)

	// Set Body
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)

		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// Create request
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// Set Headers
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(v)
	return res, err
}
