package commands

import (
	"log"

	"github.com/Bufferoverflovv/Nigella-Bot/commands/general"
	"github.com/bwmarrin/discordgo"
)

// Add commands here
var commands = []discordgo.ApplicationCommand{
	general.HelloWorldCMD,
	general.DirtyTalkCMD,
}

var registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))

// Register commands
func registercommands(s *discordgo.Session, GuildID string) {
	log.Println("Loading commands...")
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, &v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}

//deregister commands
func deregistercommands(s *discordgo.Session, GuildID string) {
	log.Println("Unloading commands...")
	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, GuildID, v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}
