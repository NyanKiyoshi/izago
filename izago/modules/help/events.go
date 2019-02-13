//+build !test

package managebot

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func init() {
	module.AddListener(onMessageReceived)
}

func onMessageReceived(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot {
		return
	}

	if messageToSend := dispatchReceivedMessage(message.Message); messageToSend != "" {
		_, err := session.ChannelMessageSend(message.ChannelID, messageToSend)

		if err != nil {
			log.Print("failed to send help: ", err)
		}
	}
}
