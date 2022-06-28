package games

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Bufferoverflovv/Nigella-Bot/extras"
	"github.com/bwmarrin/discordgo"
)

var (
	EightballRegister = discordgo.ApplicationCommand{
		Name:        "eightball",
		Description: "Ask the magic eightball a question",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "question",
				Description: "what is your question?",
				Required:    true,
			},
		},
	}

	EightballHandler = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"eightball": EightBallCommand,
	}
)

func EightBallCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	value := options[0].Value

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Description: fmt.Sprintf("**%s** **Asks**:question:\n %s \n\n:8ball: **8ball Says** \n%s", i.Member.Nick, value, ShakeEightBall()),
					Color:       extras.DISCORD_GREEN,
				},
			},
		},
	})
}

var magicAnswers = [...]string{
	// Positive outcomes
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes definitely",
	"You may rely on it",
	"As I see it, yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",

	// Neutral outcomes
	"Reply hazy try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",

	// Negative outcomes
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
	"Very doubtful",
}

func ShakeEightBall() string {
	rand.Seed(time.Now().UnixNano())
	return magicAnswers[rand.Intn(len(magicAnswers))]
}
