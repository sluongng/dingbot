package dingbot

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type respErr struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
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
		HideAvatar     string `json:"hideAvatar,omnitempty"`
		BtnOrientation string `json:"btnOrientation,omnitempty"`
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

// SendText create a message with Text type
func (c *Client) SendText(textMsg TextMessage) error { return c.send(textMsg) }

// SendLink create a message with Link type
func (c *Client) SendLink(linkMessage LinkMessage) error { return c.send(linkMessage) }

// SendMarkdown create a message with Markdown type
func (c *Client) SendMarkdown(mdMessage MarkdownMessage) error { return c.send(mdMessage) }

// SendSingleAction create a message with SingleAction type
func (c *Client) SendSingleAction(actionCard SingleActionCardMessage) error { return c.send(actionCard) }

// SendMutliAction create a message with MultiAction type
func (c *Client) SendMutliAction(actionCard MultiActionCardMessage) error { return c.send(actionCard) }

// SendFeed create a message with Feed type
func (c *Client) SendFeed(feedCard FeedCardMessage) error { return c.send(feedCard) }

func (c *Client) send(msg interface{}) error {
	req, err := c.newRobot(msg)
	if err != nil {
		return err
	}

	var res respErr
	_, err = c.do(req, &res)
	if err != nil {
		return err
	}
	if res.ErrCode != 0 {
		return fmt.Errorf(res.ErrMsg)
	}

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
