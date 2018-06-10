package main

import (
	"github.com/sluongng/dingbot"
)

const accessToken = "ca45c67f7fc647f843612c77a364cde9770cd07a9224fffc8f79fb7a02bc6faf"

func main() {
	_ = dingbot.SimpleTextMessage("Xin Chao 123").Send(accessToken)
	_ = dingbot.SimpleMarkdownMessage("\n\n- Xin Chao\n- 123123").Send(accessToken)
	_ = dingbot.NewLinkMessage(
		"Emma Watson",
		"A Picture",
		"https://i.redd.it/cb4wugxao5311.jpg",
		"https://i.redd.it/cb4wugxao5311.jpg",
	).Send(accessToken)
}
