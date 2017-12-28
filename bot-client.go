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

type RespErr struct {
	ErrCode int
	ErrMsg  string
}

type coreMsg struct {
	MsgType string `json:"msgtype"`
}

// TextMessage is used to construct Text Message body
type TextMessage struct {
	*coreMsg
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	At struct {
		AtMobile []string `json:"atMobiles,omnitempty"`
		IsAtAll  bool     `json:"isAtAll,omnitempty"`
	} `json:"at,omnitempty"`
}

// LinkMessage is used to construct Link Message body
type LinkMessage struct {
	*coreMsg
	Link struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicURL     string `json:"picUrl"`
		MessageURL string `json:"messageUrl"`
	} `json:"link"`
}

// MarkdownMessage is used to construct Markdown Message body
type MarkdownMessage struct {
	*coreMsg
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		AtMobiles []string `json:"atMobiles,omnitempty"`
		IsAtAll   bool     `json:"isAtAll,omnitempty"`
	} `json:"at,omnitempty"`
}

// SingleActionCardMessage is used to construct ActionCard Message body
type SingleActionCardMessage struct {
	*coreMsg
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		HideAvatar     string `json:"hideAvatar,omnitempty"`
		BtnOrientation string `json:"btnOrientation,omnitempty"`
		SingleTitle    string `json:"singleTitle"`
		SingleURL      string `json:"singleURL"`
	} `json:"actionCard"`
}

// MultiActionCardMessage is used to construct ActionCard Message body
type MultiActionCardMessage struct {
	*coreMsg
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		HideAvatar     string `json:"hideAvatar"`
		BtnOrientation string `json:"btnOrientation"`
		Btns           []struct {
			Title     string `json:"title"`
			ActionURL string `json:"actionURL"`
		} `json:"btns"`
	} `json:"actionCard"`
}

// FeedCardMessage is used to construct FeedCard Message body
type FeedCardMessage struct {
	*coreMsg
	FeedCard struct {
		Links []struct {
			Title      string `json:"title"`
			MessageURL string `json:"messageURL"`
			PicURL     string `json:"picURL"`
		} `json:"links"`
	} `json:"feedCard"`
}

// TextMessage create a message with Text type
func (c *Client) TextMessage() error {

	// Todo implement Client.TextMessage()

	return nil
}

// LinkMessage create a message with Link type
func (c *Client) LinkMessage() error {

	// Todo implement Client.LinkMessage()

	return nil
}

// MarkdownMessage create a message with Markdown type
func (c *Client) MarkdownMessage() error {

	// Todo implement Client.MarkdownMessage()

	return nil
}

// ActionCardMessage create a message with ActionCard type
func (c *Client) ActionCardMessage() error {

	// Todo implement Client.ActionCardMessage()

	return nil
}

// FeedCardMessage create a message with FeedCard type
func (c *Client) FeedCardMessage() error {

	// Todo implement Client.FeedCardMessage()

	return nil
}

func (c *Client) newRobot(body interface{}) (*http.Request, error) {
	req, err := c.newRequest("POST", "/robot/send", body)
	return req, err
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
