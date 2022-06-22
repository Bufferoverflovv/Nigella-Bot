package general

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	RouletteCMD = discordgo.ApplicationCommand{
		Name:        "roulette",
		Description: "Russian Roulette",
	}

	RouletteCommand = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"roulette": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: randomvalue(i.Interaction),
				},
			})
		},
	}
)

func randomvalue(i *discordgo.Interaction, s *discordgo.Session) string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 6
	value := (rand.Intn(max-min+1) + min)

	if value == 6 {
		return "BANG! " + i.Member.Nick + "'s head explodes :exploding_head:"
		GuildMemberTimeout(i.Member, i.Member, until *time.30)
	} else {
		return "*click* " + i.Member.Nick + " unfortunately survives to see another day. :gun::smiling_face_with_tear:"
	}
}
