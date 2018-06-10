package model

// LinkMessage is used to construct Link Message body
type LinkMessage struct {
	MsgType string `json:"msgtype"`
	Link    *Link  `json:"link"`
}

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicURL     string `json:"picUrl"`
	MessageURL string `json:"messageUrl"`
}

func NewLinkMessage(text string, title string, picURL string, msgURL string) *LinkMessage {
	return &LinkMessage{
		MsgType: "link",
		Link: &Link{
			Text:       text,
			Title:      title,
			PicURL:     picURL,
			MessageURL: msgURL,
		},
	}
}
