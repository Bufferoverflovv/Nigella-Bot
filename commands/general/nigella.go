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
	}

	DirtyTalkCommand = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
        {
		"nigella": func(s *discordgo.Session, i *discordgo.InteractionCreate)
            {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse
                {
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData
                    {
					Content: "https://www.youtube.com/watch?v=RtS2Ikk7A9I",
                    },
                })
            },
        }
    DirtyXmasCMD = discordgo.ApplicationCommand{
		Name:        "nigellla xmas",
		Description: "XXXmas Talk",
	}

	DirtyXmasCommand = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
        {
		"nigellla xmas": func(s *discordgo.Session, i *discordgo.InteractionCreate)
            {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse
                {
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData
                    {
					Content: "https://www.youtube.com/watch?v=2wncEeJZqzM",
                    },
                })
            },
        }
    )
