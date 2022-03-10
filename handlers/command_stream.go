package handlers

import (
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"bot/processor"
	"bot/streamer"
)

var (
	commandStreamHandler = handlers.NewCommand("stream", commandStream)
)

func commandStream(b *gotgbot.Bot, ctx *ext.Context) error {
	if processor.Processing() {
		_, err := ctx.Message.Reply(b, "Can’t stream right now.", nil)
		return err
	}
	input := ""
	args := strings.Fields(ctx.Message.Text)
	if len(args) > 1 {
		input = args[1]
	} else {
		if ctx.Message.ReplyToMessage != nil {
			input = ctx.Message.ReplyToMessage.Text
		}
	}
	if input == "" {
		_, err := ctx.Message.Reply(b, "No input provided.", nil)
		return err
	}
	status, err := ctx.Message.Reply(b, "Processing...", nil)
	if err != nil {
		return err
	}
	if err = streamer.Stream(input, ctx.Message.From); err != nil {
		return err
	}
	_, _, err = status.EditText(b, "Streaming...", nil)
	return err
}
