package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	DirtyTalkRegister = discordgo.ApplicationCommand{
		Name:        "nigella",
		Description: "Dirty Talk",
	}

	DirtyTalkHandler = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"nigella": DirtyTalkCommand,
	}
)

func DirtyTalkCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "https://www.youtube.com/watch?v=RtS2Ikk7A9I",
		},
	})
}
