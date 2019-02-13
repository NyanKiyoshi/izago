package playingstatus

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func onConnect(session *discordgo.Session, _ *discordgo.Connect) {
	if err := session.UpdateStatus(0, module.Name); err != nil {
		log.Print("failed to update playing status: ", err)
	}
}
