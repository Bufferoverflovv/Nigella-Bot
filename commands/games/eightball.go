package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	EightballRegister = discordgo.ApplicationCommand{
		Name:        "hello",
		Description: "Hi",
	}

	EightballHandler = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"hello": EightBallCommand,
	}
)

func EightBallCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hello " + i.Member.Nick,
		},
	})
}
