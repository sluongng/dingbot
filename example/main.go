package main

import (
	"fmt"

	"github.com/sluongng/dingbot"
)

const accessToken = "5465acf723849403cfddd8c3af9cb0bc4f3fc7f7624a468c810c05b4c5ff1e82"

func main() {
	ChatBotService := dingbot.NewClient(accessToken).RobotService
	msg := new(dingbot.TextMessage)
	msg.MsgType = "text"
	msg.Text.Content = "Sent From golang Dingbot"

	err := ChatBotService.SendText(msg)
	if err != nil {
		fmt.Println(err)
	}
}
