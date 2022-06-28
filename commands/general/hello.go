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

	var username string

	if i.Member.Nick == "" {
		username = i.Member.User.Username
	} else {
		username = i.Member.Nick
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hello " + username,
		},
	})
}
