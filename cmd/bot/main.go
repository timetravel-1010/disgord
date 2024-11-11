package main

import (
	"github.com/timetravel-1010/disgord/internal/bot"
)

func main() {
	b := bot.NewBot()
	b.RegisterCommands()
	b.Start()
	b.RemoveCommands()
	b.Stop()
}
