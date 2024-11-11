package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BasicCmdWithFile = &discordgo.ApplicationCommand{
	Name:        "basic-command-with-files",
	Description: "Basic command with files",
}

func BasicCommandWithFileHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey there! Congratulations, you just executed your first slash command with a file in the response",
			Files: []*discordgo.File{
				{
					ContentType: "text/plain",
					Name:        "test.txt",
					Reader:      strings.NewReader("Hello Discord!!"),
				},
			},
		},
	})
}
