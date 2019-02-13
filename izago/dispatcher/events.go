package dispatcher

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

// onMessageReceived dispatches a received message
func onMessageReceived(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore the bot' messages and bot messages
	if message.Author.ID == session.State.User.ID || message.Author.Bot {
		return
	}

	// TODO: detect if the message is DM or not
	log.Printf("Received message from %s", message.GuildID)

	// TODO: add proper handling through command prefix
	//  (=> check whether the message is a command)
	receivedCommand := message.Content

	// Handle the command
	dispatchServerCommand(receivedCommand, session, message)
}
