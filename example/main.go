package main

import (
	"fmt"

	"github.com/sluongng/dingbot"
)

const accessToken = "***"

func main() {
	ChatBotService := dingbot.NewClient(accessToken).RobotService
	msg := &dingbot.TextMessage{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: "May the force be with you",
		},
	}

	err := ChatBotService.SendText(msg)
	if err != nil {
		fmt.Println(err)
	}
}
