package main

import (
	"context"
	"log"
	"os"

	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	definition := &slacker.CommandDefinition{
		Description: "デプロイのリマインド",
		Example:     "remind @sato-s",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			// user := request.StringParam("user", "")
			response.Reply("<@satos> test.....")
			client := botCtx.Client()
			ev := botCtx.Event()
			// https://pkg.go.dev/github.com/slack-go/slack#Client.AddChannelReminder
			if ev.Channel != "" {
				client.PostMessage(ev.Channel, slack.MsgOptionText("<@satos> saaas", false))
				client.PostMessage(ev.Channel, slack.MsgOptionText("<@UD9T7FXCJ> saaas", false))
				client.PostMessage(ev.Channel, slack.MsgOptionText("/who", false))
				_, err := client.AddChannelReminder(ev.Channel, "今日のデスク当番は aa", "15:00")
				if err != nil {
					log.Printf("Failed to set error %+v", err)
				}
				client.AddChannelReminder(ev.Channel, "デプロイの時間だね aa", "18:00")
			}
		},
	}

	bot.Command("remind <user>", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
