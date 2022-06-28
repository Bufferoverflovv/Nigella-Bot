package commands

import (
	"log"

	"github.com/Bufferoverflovv/Nigella-Bot/commands/games"
	"github.com/Bufferoverflovv/Nigella-Bot/commands/general"
	"github.com/bwmarrin/discordgo"
)

// Add commands here
var slashcommands = []discordgo.ApplicationCommand{
	games.EightballRegister,
	general.HelloRegister,
}

var registeredCommands = make([]*discordgo.ApplicationCommand, len(slashcommands))

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	general.HelloRegister.Name:     general.HelloCommand,
	general.RouletteRegister.Name:  general.RouletteCommand,
	general.DirtyTalkRegister.Name: general.DirtyTalkCommand,
	games.EightballRegister.Name:   games.EightBallCommand,
}

// Register commands
func RegisterCommands(s *discordgo.Session, GuildID string) {
	log.Println("Loading commands...")
	for i, v := range slashcommands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, &v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

//deregister commands
func DeregisterCommands(s *discordgo.Session, GuildID string) {
	log.Println("Unloading commands...")
	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, GuildID, v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}

func cachecommands() {
	// TODO
}

func loadcachecommand() {
	// TODO
}
