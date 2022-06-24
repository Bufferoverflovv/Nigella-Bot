package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	"github.com/Bufferoverflovv/Nigella-Bot/commands"
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

	commands.RegisterCommands(s, GuildID)

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	commands.DeregisterCommands(s, GuildID)

	log.Println("Gracefully shutting down.")
}
