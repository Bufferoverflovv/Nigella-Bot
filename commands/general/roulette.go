package general

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	RouletteRegister = discordgo.ApplicationCommand{
		Name:        "roulette",
		Description: "Russian Roulette",
	}

	RouletteHandler = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"roulette": RouletteCommand,
	}
)

func RouletteCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: randomvalue(i.Interaction, s),
		},
	})
}

func randomvalue(i *discordgo.Interaction, s *discordgo.Session) string {
	min := 1
	max := 6
	rand.Seed(time.Now().UnixNano())
	value := (rand.Intn(max-min) + min)

	timeout := time.Now().Add(time.Minute)

	if value == 6 {
		s.GuildMemberTimeout(i.GuildID, i.Member.User.ID, &timeout)
		return "BANG! " + i.Member.Nick + "'s head explodes :exploding_head:"
	} else {
		return "*click* " + i.Member.Nick + " unfortunately survives to see another day. :gun::smiling_face_with_tear:"
	}
}
