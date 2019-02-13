package dispatcher

import (
	"github.com/bwmarrin/discordgo"
)

// onMessageReceived dispatches a received message
func onMessageReceived(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore the bot' messages and bot messages
	if message.Author.ID == session.State.User.ID || message.Author.Bot {
		return
	}

	// TODO: add proper handling through command prefix
	//  (=> check whether the message is a command)
	receivedCommand := message.Content

	// Handle the command
	dispatchCommand(receivedCommand, session, message)
}
