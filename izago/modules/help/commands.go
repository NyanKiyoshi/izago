//+build !test

package managebot

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	module.AddCommand(
		"help", commandGetHelp,
		`Gets the list of commands.`)
}

func commandGetHelp(session *discordgo.Session, create *discordgo.MessageCreate) {
	_, _ = session.ChannelMessageSend(create.ChannelID, getHelp(session))
}
