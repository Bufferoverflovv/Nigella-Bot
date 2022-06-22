package main

import (
	//"fmt"

	"log"
	"os"
	"os/signal"

	//"github.com/Bufferoverflovv/Nigella-Bot/commands"
	"github.com/Bufferoverflovv/Nigella-Bot/commands/general"
	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session

func init() {
	var err error
	s, err = discordgo.New("Bot " + DiscordToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
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
		if h, ok := general.rouletteCommand[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	s.ApplicationCommandCreate(s.State.User.ID, GuildID, &general.rouletteCMD)
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
