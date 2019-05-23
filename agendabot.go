package main

import (
	"cmargonis.dev/agenda-manager/configuration"
	"fmt"
	"strings"
	"github.com/nlopes/slack"
)

type AuthConfiguration interface {
	GetToken() string
}

func main() {
	var conf AuthConfiguration = configuration.NewFileConfiguration()
	api := slack.New(conf.GetToken())
	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
					rtm.SendMessage(rtm.NewOutgoingMessage("Hello folks, agenda automation comming soon:tm:!", ev.Channel))
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}
}
