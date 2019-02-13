//+build !test

package managebot

import (
	"github.com/bwmarrin/discordgo"
)

func commandGetHelp(session *discordgo.Session, create *discordgo.MessageCreate) {
	_, _ = session.ChannelMessageSend(create.ChannelID, getHelp(session))
}
