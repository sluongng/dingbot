package model

// TextMessage is used to construct Text Message body
type TextMessage struct {
	MsgType string `json:"msgtype"`
	Text    *Text  `json:"text"`
	At      *AtTag `json:"at,omitempty"`
}

type Text struct {
	Content string `json:"content"`
}

func NewTextMessage(content string, at *AtTag) *TextMessage {
	return &TextMessage{
		MsgType: "text",
		Text:    &Text{Content: content},
		At:      at,
	}
}

func SimpleTextMessage(content string) *TextMessage {
	return NewTextMessage(content, SimpleAtTag())
}
