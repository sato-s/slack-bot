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
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("<!satos> pongaaaa24")
			client := botCtx.Client()
			ev := botCtx.Event()
			if ev.Channel != "" {
				client.PostMessage(ev.Channel, slack.MsgOptionText("<!satos> saaas", false))
				client.PostMessage(ev.Channel, slack.MsgOptionText("<!UD9T7FXCJ> saaas", false))
				client.PostMessage(ev.Channel, slack.MsgOptionText("/who", false))
			}
		},
	}

	bot.Command("ping", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
