package handlers

import (
	"github.com/bwmarrin/discordgo"
)

var BasicCmd = &discordgo.ApplicationCommand{
	Name: "basic-command",
	// All commands and options must have a description
	// Commands/options without description will fail the registration
	// of the command.
	Description: "Basic command",
}

var msg = `
Hey there! Congratulations, you just executed your first slash command.
> Here's an image: 
> https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fimages7.alphacoders.com%2F106%2F1065297.jpg&f=1&nofb=1&ipt=9cf7a579dac5dcbb888ddbf7e0800b2fd1dd9000f24979ed6c2be7efc7a52a3a&ipo=images
`

func BasicCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
}
