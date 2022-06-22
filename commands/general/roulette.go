package general

use rand::Rng;

fn main() {
		let mut rng = rand::thread_rng();

		let num: u8 = rng.gen_range(0..6);
}

import (
	"github.com/bwmarrin/discordgo"
)

var (
	rouletteCMD = discordgo.ApplicationCommand{
		Name:        "roulette",
		Description: "Russian Roulette",
	}

	rouletteCommand = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		if num = 6;
			"roulette": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "BANG! <@!"+o.ID+">'s head explodes :exploding_head: ",
					
				},
			})
		},
		else; 
			"roulette": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "*click* <@!"+o.ID+"> unfortunately survives to see another day. :gun::nigellaglow:",
				
				},
			})
		},
	}	 
)