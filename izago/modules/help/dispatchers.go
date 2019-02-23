package managebot

import (
	"github.com/NyanKiyoshi/izago/izago/globals"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// helpCommand contains the command name with its prefix,
// it will get set once 'dispatchReceivedMessage' is invoked.
var helpCommand string

func dispatchReceivedMessage(message *discordgo.Message) string {
	// Adds the configured prefix to the 'help' command
	// once this function is first executed
	globals.AsyncInit.Do(func() {
		helpCommand = globals.Config.Prefix + "help"
	})

	messageCommands := strings.SplitN(message.Content, " ", 2)
	receivedCommand := strings.TrimSpace(strings.ToLower(messageCommands[0]))

	if len(messageCommands) > 1 && receivedCommand == helpCommand {
		if commandName := strings.TrimSpace(messageCommands[1]); commandName != "" {
			return getCommandHelpText(commandName, message.GuildID == "")
		}
	}

	return ""
}
