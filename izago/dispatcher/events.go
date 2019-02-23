package dispatcher

import (
	"github.com/NyanKiyoshi/izago/izago/globals"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// onMessageReceived dispatches a received message
func onMessageReceived(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore the bot' messages and bot messages
	if message.Author.ID == session.State.User.ID || message.Author.Bot {
		return
	}

	// Check if the message is a command
	if !strings.HasPrefix(
		message.Content, globals.Config.Prefix) {
		return
	}
	receivedCommand := strings.ToLower(strings.TrimPrefix(
		message.Content, globals.Config.Prefix))

	// Handle the command
	dispatchCommand(receivedCommand, session, message)
}
