package commands

import
    (
	"github.com/bwmarrin/discordgo"
    )

var
    (
	DirtyTalkCMD = discordgo.ApplicationCommand{
		Name:        "nigella",
		Description: "Dirty Talk",
        Options: []*discordgo.ApplicationCommandOption{
            {
            Type:        discordgo.ApplicationCommandOptionString,
            Name:        "xmas",
            Description: "XXXmas",
            Required:    false,
	}

	DirtyTalkCommand = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
        {
		"nigella": func(s *discordgo.Session, i *discordgo.InteractionCreate)
            {
			options := i.ApplicationCommandData().Options
                {
				if option, ok := optionMap["xmas"]; ok {
				margs = append(margs, option.StringValue())
				msgformat += "https://www.youtube.com/watch?v=2wncEeJZqzM"{
                    }
                }
                else s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
                    Type: discordgo.InteractionResponseChannelMessageWithSource,
                    Data: &discordgo.InteractionResponseData{
					Content: "https://www.youtube.com/watch?v=RtS2Ikk7A9I",
                    }
                    })
                }
            }
        }
    )
