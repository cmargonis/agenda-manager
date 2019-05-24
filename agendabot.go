package main

import (
	"fmt"
	"strings"

	"cmargonis.dev/agenda-manager/configuration"
	"github.com/nlopes/slack"
)

type AuthConfiguration interface {
	GetToken() string
}

func main() {
	var conf AuthConfiguration = configuration.NewFileConfiguration()
	slackApi := slack.New(conf.GetToken())
	rtm := slackApi.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch messageData := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", messageData.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", messageData)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				if messageData.User != info.User.ID && strings.HasPrefix(messageData.Text, prefix) {
					rtm.SendMessage(rtm.NewOutgoingMessage("Hello folks, agenda automation comming soon:tm:!", messageData.Channel))
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", messageData.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}
}
