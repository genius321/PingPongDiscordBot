package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/genius321/PingPongDiscordBot/config"
)

func Start() error {
	// Create a new Discord session using the provided bot token.
	session, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		return err
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	session.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = session.Open()

	if err != nil {
		return err
	}

	return nil
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "!ping" reply with "Pong!"
	if m.Content == config.BotPrefix+"ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "!pong" reply with "Ping!"
	if m.Content == config.BotPrefix+"pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
