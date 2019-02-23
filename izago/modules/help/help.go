package managebot

import (
	"fmt"
	"github.com/NyanKiyoshi/izago/izago/dispatcher"
	"github.com/NyanKiyoshi/izago/izago/globals"
	"github.com/bwmarrin/discordgo"
)

var helpText string

func generateHelpText(session *discordgo.Session) {
	session.Lock()
	defer session.Unlock()

	helpText = "```Usage:"

	for _, module := range dispatcher.ActivatedModules {
		for name, command := range module.ServerCommands {
			helpText += fmt.Sprintf(
				"\n%s%-15s%s", globals.Config.Prefix, name, command.ShortHelp)
		}
	}

	helpText += "```"
}

// getHelp returns the cached help text or generated the help text.
func getHelp(session *discordgo.Session) string {
	if helpText == "" {
		generateHelpText(session)
	}
	return helpText
}

func getCommandHelpText(commandName string, isDm bool) string {
	commandDefs := dispatcher.FindCommand(commandName, isDm)

	if commandDefs == nil {
		return fmt.Sprintf("help: no such command (%s).", commandName)
	}

	return fmt.Sprintf(
		"```%s%s:\n%s```", globals.Config.Prefix, commandName, commandDefs.LongHelp)
}
