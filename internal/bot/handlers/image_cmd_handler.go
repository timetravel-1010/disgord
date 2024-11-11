package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var ImageCmd = &discordgo.ApplicationCommand{
	Name:        "image-command",
	Description: "Basic command to display a message with an image",
}

func ImgCmdHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	imageURL := "https://www.nautiljon.com/images/manga_persos/00/73/mini/337.jpg"
	messageText := "Here's an image:" // The text you want to show above the image

	embed := &discordgo.MessageEmbed{
		Description: messageText,
		Image: &discordgo.MessageEmbedImage{
			URL: imageURL,
		},
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
	if err != nil {
		log.Printf("error sending interaction response: %v", err)
	}
}
