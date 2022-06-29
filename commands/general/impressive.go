package general

import (
	"github.com/bwmarrin/discordgo"
)

var (
	ImpressiveRegister = discordgo.ApplicationCommand{
		Name:        "impressive",
		Description: "Impressive",
	}

	ImpressiveHandler = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"impressive": ImpressiveCommand,
	}
)

func ImpressiveCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "<:impressive:964738004386201650>",
		},
	})
}
