package main

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func initDiscord(c *yamlConf) *discordgo.Session {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + c.Discord.Token)
	if err != nil {
		log.Fatal("[Error] creating Discord session", err)
	}

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatal("[Error] opening Discord session: ", err)
	}

	log.Println("Connection discord... [OK]")
	return dg
}

// Messages from discord chat
func (b *mtdBot) fromDiscord(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	//log.Println(m.State)
	if m.Author.ID == s.State.User.ID {
		return
	}

	msg := message{
		Type:        "chat",
		MessageText: m.Content,
		PlayerName:  m.Author.Username,
	}
	// Getting message from needed channel only
	//if m.ChannelID == b.Config.Discord.ChannelID {
	if strings.HasPrefix(m.Content, "!help") {
		s.ChannelMessageSend(m.ChannelID, b.Config.Discord.Help)
		return
	}
	for _, command := range b.Config.Discord.Commands {
		if strings.HasPrefix(msg.MessageText, "!"+command) {
			msg.Type = "cmd"
			msg.Command = command
			msg.Args = strings.Fields(strings.TrimSpace(strings.TrimPrefix(msg.MessageText, "!"+command)))
			msg.MessageText = ""
		}
	}
	//}
	// sending message to the Minetest
	b.toMT(msg, m.ChannelID)
}

func (b *mtdBot) toDiscord(m message, chID string) {
	switch m.Type {
	case "chat":
		b.Discord.ChannelMessageSend(chID, "<"+m.PlayerName+"> "+m.MessageText)
	case "system":
		b.Discord.ChannelMessageSendEmbed(chID, &discordgo.MessageEmbed{
			Title:       m.TitleText,
			Description: m.MessageText,
			Color:       0x00e200,
		})
	default:
	}
}
