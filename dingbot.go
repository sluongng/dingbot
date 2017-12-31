package dingbot

import (
	"fmt"

	"github.com/dghubble/sling"
)

const (
	dingTalkAPI = "https://oapi.dingtalk.com"
)

type defaultParams struct {
	AccessToken string `url:"access_token"`
}

type responseError struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type coreMsg struct {
	MsgType string `json:"msgtype"`
}

type atTag struct {
	AtMobiles []string `json:"atMobiles,omnitempty"`
	IsAtAll   bool     `json:"isAtAll,omnitempty"`
}

// TextMessage is used to construct Text Message body
type TextMessage struct {
	*coreMsg
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	At atTag `json:"at,omnitempty"`
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
	At atTag `json:"at,omnitempty"`
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

// RobotService provides methods for sending messages
type RobotService struct {
	sling *sling.Sling
}

// NewRobotService returns a new RobotService
func NewRobotService(sling *sling.Sling) *RobotService {
	return &RobotService{
		sling: sling.Path("/robot"),
	}
}

func (rs *RobotService) send(v interface{}) error {
	defaultError := new(responseError)
	_, err := rs.sling.New().Post("/send").BodyJSON(v).ReceiveSuccess(defaultError)
	if err != nil {
		return err
	}
	if defaultError.ErrCode != 0 {
		return fmt.Errorf(defaultError.ErrMsg)
	}
	return nil
}

// SendText create a message with Text type
func (rs *RobotService) SendText(textMsg *TextMessage) error { return rs.send(textMsg) }

// SendLink create a message with Link type
func (rs *RobotService) SendLink(linkMessage *LinkMessage) error { return rs.send(linkMessage) }

// SendMarkdown create a message with Markdown type
func (rs *RobotService) SendMarkdown(mdMessage *MarkdownMessage) error { return rs.send(mdMessage) }

// SendSingleAction create a message with SingleAction type
func (rs *RobotService) SendSingleAction(actionCard *SingleActionCardMessage) error {
	return rs.send(actionCard)
}

// SendMutliAction create a message with MultiAction type
func (rs *RobotService) SendMutliAction(actionCard *MultiActionCardMessage) error {
	return rs.send(actionCard)
}

// Client is a General DingTalk client which contains many services
type Client struct {
	sling        *sling.Sling
	RobotService *RobotService
}

// NewClient returns a new Client
func NewClient(accessToken string) *Client {
	base := sling.New().
		Client(nil).
		Base(dingTalkAPI).
		QueryStruct(&defaultParams{AccessToken: accessToken})
	return &Client{
		sling:        base,
		RobotService: NewRobotService(base.New()),
	}
}
