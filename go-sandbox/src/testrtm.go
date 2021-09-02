package main

import (
	"github.com/nlopes/slack"
)

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()
}

func main() {
	api := slack.New("xoxb-37305511778-395592091399-6EeD7FaDXehbxM2in4GEBSmw")

	rtm := api.NewRTM()

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			rtm.SendMessage(rtm.NewOutgoingMessage("やぁ", ev.Channel))
		}
	}
}
