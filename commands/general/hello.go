package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	HelloWorldCMD = discordgo.ApplicationCommand{
		Name:        "hello",
		Description: "Hi",
	}

	HelloWorldCommand = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"hello": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hi There",
				},
			})
		},
	}
)
