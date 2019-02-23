package managebot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"runtime"
)

const mibiByte = float64(1 << 20)

func commandGetBotStatus(session *discordgo.Session, message *discordgo.MessageCreate) {
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)

	_, _ = session.ChannelMessageSend(message.ChannelID, fmt.Sprintf(
		"Bot Uptime: %s\nMemory Usage: %.2f MiB",
		getUptime(),
		float64(memstats.Sys)/mibiByte,
	))
}
