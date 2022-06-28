package general

import (
	"github.com/bwmarrin/discordgo"
)

var (
	HelloRegister = discordgo.ApplicationCommand{
		Name:        "hello",
		Description: "Hi",
	}

	HelloWorldCommand = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"hello": HelloCommand,
	}
)

func HelloCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hello " + i.Member.Nick,
		},
	})
}
