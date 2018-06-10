package model

type AtTag struct {
	AtMobiles []string `json:"atMobiles,omnitempty"`
	IsAtAll   bool     `json:"isAtAll,omnitempty"`
}

func NewAtTag(mobiles []string, isAtAll bool) *AtTag {
	return &AtTag{
		AtMobiles: mobiles,
		IsAtAll:   isAtAll,
	}
}

func SimpleAtTag() *AtTag {
	return NewAtTag([]string{}, false)
}
