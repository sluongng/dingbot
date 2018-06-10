package model

// MarkdownMessage is used to construct Markdown Message body
type MarkdownMessage struct {
	MsgType  string    `json:"msgtype"`
	Markdown *Markdown `json:"markdown"`
	At       *AtTag    `json:"at,omnitempty"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func NewMarkdownMessage(title string, text string, at *AtTag) *MarkdownMessage {
	return &MarkdownMessage{
		MsgType: "markdown",
		Markdown: &Markdown{
			Title: title,
			Text:  text,
		},
		At: at,
	}
}

func SimpleMarkdownMessage(text string) *MarkdownMessage {
	return NewMarkdownMessage("markdown", text, SimpleAtTag())
}
