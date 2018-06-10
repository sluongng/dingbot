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

func EmptyAtTag() *AtTag {
	return NewAtTag([]string{}, false)
}

func SingleAtTag(mobileNumber string) *AtTag {
	return NewAtTag([]string{mobileNumber}, true)
}
