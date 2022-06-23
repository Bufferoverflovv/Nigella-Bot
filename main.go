package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	"github.com/Bufferoverflovv/Nigella-Bot/commands/general"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v3"
)

var s *discordgo.Session

var GuildID string

type Config struct {
	Discord struct {
		Token   string `yaml:"token"`
		Guildid string `yaml:"guildid"`
	} `yaml:"discord"`
}

func init() {

	f := &Config{}
	source, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Invalid Configuration: %v", err)
	}

	err = yaml.Unmarshal([]byte(source), &f)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	GuildID = f.Discord.Guildid

	var err2 error
	s, err2 = discordgo.New("Bot " + f.Discord.Token)
	if err2 != nil {
		log.Fatalf("Invalid bot parameters: %v", err2)
	}
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := general.HelloWorldCommand[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := general.DirtyTalkCommand[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := general.RouletteCommand[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	s.ApplicationCommandCreate(s.State.User.ID, GuildID, &general.RouletteCMD)
	s.ApplicationCommandCreate(s.State.User.ID, GuildID, &general.HelloWorldCMD)
	s.ApplicationCommandCreate(s.State.User.ID, GuildID, &general.DirtyTalkCMD)

	// log.Println("Adding commands...")
	// registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.SlashCommands))
	// //registeredCommands[0] = &commands.HelloWorldCMD
	// for i, v := range commands.SlashCommands {
	// 	cmd, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, v)
	// 	if err != nil {
	// 		log.Panicf("Cannot create '%v' command: %v", v.Name, err)
	// 	}
	// 	registeredCommands[i] = cmd
	// }

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	// for _, v := range registeredCommands {
	// 	err := s.ApplicationCommandDelete(s.State.User.ID, GuildID, v.ID)
	// 	if err != nil {
	// 		log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
	// 	}
	// }

	log.Println("Gracefully shutting down.")
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if m.Content == "hello" {
		s.ChannelMessageSend(m.ChannelID, "Hello, "+m.Author.Username)
	}

	if m.Content == "hello there" {
		s.ChannelMessageSend(m.ChannelID, "General Kenobi!")
	}
}
