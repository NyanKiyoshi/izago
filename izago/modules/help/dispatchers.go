package managebot

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func dispatchReceivedMessage(message *discordgo.Message) string {
	messageCommands := strings.SplitN(message.Content, " ", 2)
	receivedCommand := strings.TrimSpace(strings.ToLower(messageCommands[0]))

	// TODO: add prefix
	if len(messageCommands) > 1 && receivedCommand == "help" {
		if commandName := strings.TrimSpace(messageCommands[1]); commandName != "" {
			return getCommandHelpText(commandName, message.GuildID == "")
		}
	}

	return ""
}
